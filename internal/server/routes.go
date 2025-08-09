package server

import (
	"net/http"

	"github.com/TroJanBoi/assembly-visual-backend/cmd/api/docs"
	"github.com/TroJanBoi/assembly-visual-backend/internal/conf"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/controller"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/repository"
	"github.com/TroJanBoi/assembly-visual-backend/internal/services/usecases"
	"github.com/TroJanBoi/assembly-visual-backend/security"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-User-ID")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) Router() (http.Handler, func()) {
	config := conf.NewConfig()
	r := gin.Default()

	r.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v2"

	catRepository := repository.NewCatRepository(s.db)
	catUseCase := usecases.NewCatUseCase(catRepository)
	catController := controller.NewCatController(catUseCase)

	oauthRepository := repository.NewOAuthRepository()
	oauthUseCase := usecases.NewOAuthUseCase(oauthRepository)
	oauthController := controller.NewOAuthController(oauthUseCase)

	userRepository := repository.NewUserRepository(s.db)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	api := r.Group("/api/v2")
	{
		oauthController.OAuthRegisterRoutes(api)
		catGroup := api.Group("/cats").Use(security.Middleware())
		{
			catController.CatRegisterRoutes(catGroup)
		}
		userGroup := api.Group("/users").Use(security.Middleware())
		{
			userController.UserRoutes(userGroup)
		}
	}
	if config.ENV == "dev" || config.ENV == "uat" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return r, func() {
		// sqlDB, err := s.db.DB()
		// if err != nil {
		// 	panic("Failed to get sql.DB from gorm.DB")
		// }
		// sqlDB.Close()
	}
}
