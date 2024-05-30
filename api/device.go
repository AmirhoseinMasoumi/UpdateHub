package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateDeviceRequest struct {
	DeviceId string `json:"device_id" binding:"required"`
}

func (server *Server) CreateDevice(ctx *gin.Context) {
	var req CreateDeviceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctxBg := context.Background()
	startDate := time.Time{}
	endDate := time.Now()

	updates, err := server.store.ListUpdatesBetweenDates(ctxBg, db.ListUpdatesBetweenDatesParams{
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var existingUpdate *db.Update
	for _, update := range updates {
		if update.Version == "1.0.0" {
			existingUpdate = &update
			break
		}
	}

	if existingUpdate == nil {
		update, err := server.store.CreateUpdate(ctxBg, db.CreateUpdateParams{
			Version:     "1.0.0",
			Description: "Initial version",
			Path:        "-",
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		log.Println("Created initial update:", update.Version)
	}

	device, err := server.store.CreateDevice(ctx, req.DeviceId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, device)
}

func (server *Server) GetAllDevices(ctx *gin.Context) {
	devices, err := server.store.ListAllDevices(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, devices)
}

type DeleteDeviceRequest struct {
	DeviceID string `json:"device_id" binding:"required"`
}

func (server *Server) DeleteDevice(ctx *gin.Context) {
	var req DeleteDeviceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.DeleteDevice(ctx, req.DeviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Device deleted"})
}
