package api

import (
	db "github.com/Yessentemir256/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}
