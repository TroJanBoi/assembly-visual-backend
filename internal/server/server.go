package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/conf"
	"github.com/TroJanBoi/assembly-visual-backend/internal/database"
	"gorm.io/gorm"
)

type Server struct {
	port int

	db *gorm.DB
}

const (
	ServerReadHeaderTimeout = 5 * time.Second
	ServerReadTimeout       = 5 * time.Second
	ServerWriteTimeout      = 10 * time.Second
	handlerTimeout          = ServerWriteTimeout - (time.Millisecond * 100)
)

func NewServer() *http.Server {
	config := conf.NewConfig()
	NewServer := &Server{
		port: config.PORT,

		db: database.New().GetClient(),
	}

	r, stop := NewServer.Router()
	defer stop()
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", NewServer.port),
		Handler:           r,
		ReadHeaderTimeout: ServerReadHeaderTimeout,
		ReadTimeout:       ServerReadTimeout,
		WriteTimeout:      ServerWriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	return server
}
