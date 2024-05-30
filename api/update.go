package api

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	db "github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateUpdateRequest struct {
	Version     string `json:"version" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) CreateUpdate(ctx *gin.Context) {
	var req CreateUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updateDir := "./updates"
	if err := os.MkdirAll(updateDir, os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	filePath := fmt.Sprintf("%s/%s.zip", updateDir, req.Version)

	_, err := server.store.GetUpdate(context.Background(), req.Version)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Update with the same version already exists"})
		return
	} else if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	update, err := server.store.CreateUpdate(context.Background(), db.CreateUpdateParams{
		Version:     req.Version,
		Description: req.Description,
		Path:        filePath,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, update)
}

type CheckUpdateRequest struct {
	DeviceId string `json:"device_id" binding:"required"`
	Version  string `json:"version" binding:"required"`
}

func (server *Server) CheckUpdate(ctx *gin.Context) {
	var req CheckUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	device, err := server.store.GetDevice(context.Background(), req.DeviceId)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	result, err := compareVersions(req.Version, device.DeviceVersion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if result == 1 {
		_, err = server.store.UpdateDevice(ctx, db.UpdateDeviceParams{
			DeviceID:      req.DeviceId,
			DeviceVersion: req.Version,
			LastUpdate:    time.Now(),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	nextUpdate, err := server.store.GetNextVersion(context.Background(), req.Version)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, gin.H{"message": "No new updates available"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	filePath := nextUpdate.Path
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Update file not found"})
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer file.Close()

	fileName := filepath.Base(filePath)
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
	ctx.Header("Content-Checksum", nextUpdate.Checksum)

	n, err := io.Copy(io.Writer(ctx.Writer), file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if n != fileInfo.Size() {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("failed to copy the entire file")))
		return
	}
}

func (server *Server) UploadUpdate(ctx *gin.Context) {
	description := ctx.PostForm("description")

	if description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fileName := filepath.Base(file.Filename)

	valid, err := isValidFileName(fileName)
	if err != nil || !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filename format"})
		return
	}

	versionFromFileName := getVersionFromFileName(fileName)

	_, err = server.store.GetUpdate(context.Background(), versionFromFileName)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Update with the same version already exists"})
		return
	} else if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	filePath := "./updates/" + fileName

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	checksum, err := calculateChecksum(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.CreateUpdate(context.Background(), db.CreateUpdateParams{
		Version:     versionFromFileName,
		Description: description,
		Path:        filePath,
		Checksum:    checksum,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{"error": "Update with this version already exists"})
				return
			case "foreign_key_violation":
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid foreign key reference"})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": fileName})
}

func calculateChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func isValidFileName(fileName string) (bool, error) {
	pattern := `^\d+\.\d+\.\d+\.zip$`
	matched, err := regexp.MatchString(pattern, fileName)
	if err != nil {
		return false, err
	}
	return matched, nil
}

func getVersionFromFileName(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1] + "." + parts[2]
	}
	return ""
}

type AllUpdatesRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

type UpdateInfo struct {
	Version     string    `json:"version"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (req *AllUpdatesRequest) Validate() error {
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return errors.New("invalid start date format")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return errors.New("invalid end date format")
	}

	if startDate.After(endDate) || startDate.Equal(endDate) {
		return errors.New("start date must be earlier than end date")
	}
	return nil
}

func (server *Server) GetUpdatesBetweenDates(ctx *gin.Context) {
	var req AllUpdatesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	updates, err := server.store.ListUpdatesBetweenDates(ctx, db.ListUpdatesBetweenDatesParams{
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updates"})
		return
	}

	var updatesInfo []UpdateInfo
	for _, update := range updates {
		updatesInfo = append(updatesInfo, UpdateInfo{
			Version:     update.Version,
			Description: update.Description,
			Date:        update.Date,
		})
	}

	ctx.JSON(http.StatusOK, updatesInfo)
}

type DeleteUpdateRequest struct {
	Version string `json:"version" binding:"required"`
}

func (server *Server) DeleteUpdate(ctx *gin.Context) {
	var req DeleteUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.DeleteUpdate(ctx, req.Version)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Update deleted"})
}

func compareVersions(version1, version2 string) (int, error) {
	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")

	if len(v1Parts) != len(v2Parts) {
		return 0, fmt.Errorf("versions have different number of parts")
	}

	for i := 0; i < len(v1Parts); i++ {
		v1Part, err := strconv.Atoi(v1Parts[i])
		if err != nil {
			return 0, fmt.Errorf("failed to parse version part: %v", err)
		}

		v2Part, err := strconv.Atoi(v2Parts[i])
		if err != nil {
			return 0, fmt.Errorf("failed to parse version part: %v", err)
		}

		if v1Part < v2Part {
			return -1, nil
		} else if v1Part > v2Part {
			return 1, nil
		}
	}
	return 0, nil
}
