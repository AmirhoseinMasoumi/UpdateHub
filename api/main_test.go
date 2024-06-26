package api

import (
	"os"
	"testing"

	db "github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/db/sqlc"
	"github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
