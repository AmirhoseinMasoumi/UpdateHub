package api

import (
	"net/http"
	"path/filepath"

	"github.com/rs/zerolog/log"

	db "github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/db/sqlc"
	"github.com/AmirhoseinMasoumi/GoProjects/DeviceUpdateManager/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	staticDir, err := filepath.Abs("static")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot get absolute path to static directory")
	}
	router.StaticFS("/static", http.Dir(staticDir))

	router.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(staticDir, "index.html"))
	})

	router.POST("/device", server.CreateDevice)
	router.POST("/update", server.CreateUpdate)

	router.POST("/update_check", server.CheckUpdate)

	router.POST("/updates", server.GetUpdatesBetweenDates)
	router.GET("/devices", server.GetAllDevices)

	router.POST("/upload_update", server.UploadUpdate)

	router.POST("/update_delete", server.DeleteUpdate)
	router.POST("/device_delete", server.DeleteDevice)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
