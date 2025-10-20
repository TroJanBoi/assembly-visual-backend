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

	operationRepository := repository.NewOperationRepository(s.db)
	operationUsecase := usecases.NewOperationUsecase(operationRepository)
	operationController := controller.NewOperationController(operationUsecase)

	authRepository := repository.NewAuthRepository(s.db)
	authUseCase := usecases.NewAuthUseCase(authRepository)
	authController := controller.NewAuthController(authUseCase)

	profileReposrtory := repository.NewProfileRepository(s.db)
	profileUseCase := usecases.NewProfileUseCase(profileReposrtory)
	profileController := controller.NewProfileController(profileUseCase)

	classRepository := repository.NewClassRepository(s.db)
	classUseCase := usecases.NewClassUseCase(classRepository)
	classController := controller.NewClassController(classUseCase)

	classNotLoginRepository := repository.NewClassRepository(s.db)
	classNotLoginUseCase := usecases.NewClassUseCase(classNotLoginRepository)
	classNotLoginController := controller.NewClassController(classNotLoginUseCase)

	assignmentRepositoryInClass := repository.NewAssignmentRepository(s.db)
	assignmentUseCaseInClass := usecases.NewAssignmentUseCase(assignmentRepositoryInClass)
	assignmentControllerInClass := controller.NewAssignmentController(assignmentUseCaseInClass)

	assignmentNotLogin := repository.NewAssignmentRepository(s.db)
	assignmentUseCaseNotLogin := usecases.NewAssignmentUseCase(assignmentNotLogin)
	assignmentControllerNotLogin := controller.NewAssignmentController(assignmentUseCaseNotLogin)

	invitationRepository := repository.NewInvitationRepository(s.db)
	invitationUseCase := usecases.NewInvitationUseCase(invitationRepository)
	invitationController := controller.NewInvitationController(invitationUseCase)

	invitationMeRepository := repository.NewInvitationRepository(s.db)
	invitationMeUseCase := usecases.NewInvitationUseCase(invitationMeRepository)
	invitationMeController := controller.NewInvitationController(invitationMeUseCase)

	// Initialize TestSuiteController
	testSuiteRepository := repository.NewTestSuiteRepository(s.db)
	testSuiteUseCase := usecases.NewTestSuitesUseCases(testSuiteRepository)
	testSuiteController := controller.NewTestSuiteController(testSuiteUseCase)

	testCaseRepository := repository.NewTestCaseRepository(s.db)
	testCaseUseCase := usecases.NewTestCaseUseCases(testCaseRepository)
	testCaseController := controller.NewTestCaseController(testCaseUseCase)

	systemsRepository := repository.NewSystemsRepository(s.db)
	systemsUseCase := usecases.NewSystemsUseCase(systemsRepository)
	systemsController := controller.NewSystemsController(systemsUseCase)

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
		operationGroup := api.Group("/operations").Use(security.Middleware())
		{
			operationController.OperationRegisterRoutes(operationGroup)
		}
		authGroup := api.Group("/auth")
		{
			authController.AuthRoutes(authGroup)
		}
		profileGroup := api.Group("/profile").Use(security.Middleware())
		{
			profileController.ProfileRoutes(profileGroup)
		}
		classGroup := api.Group("/classes").Use(security.Middleware()).(*gin.RouterGroup)
		{
			classController.ClassRoutes(classGroup)
			assignmentGroup := classGroup.Group("/:class_id/assignments").Use(security.Middleware()).(*gin.RouterGroup)
			{
				assignmentControllerInClass.AssignmentRoutes(assignmentGroup)
				// Register TestSuite routes within the assignment group
				testSuiteGroup := assignmentGroup.Group("/:assignment_id/test-suites").Use(security.Middleware()).(*gin.RouterGroup)
				{
					testSuiteController.TestSuiteRoutes(testSuiteGroup)
					testCaseGroup := testSuiteGroup.Group("/:test_suite_id/test-cases").Use(security.Middleware()).(*gin.RouterGroup)
					{
						testCaseController.TestCaseRoutes(testCaseGroup)
					}
				}
			}
			invitationGroup := classGroup.Group("/:class_id/invitations").Use(security.Middleware()).(*gin.RouterGroup)
			{
				invitationController.InvitaionRoutes(invitationGroup)
			}
		}
		invitationMeGroup := api.Group("/invitations").Use(security.Middleware()).(*gin.RouterGroup)
		{
			invitationMeController.InvitationMeRoutes(invitationMeGroup)
		}
		classNotLoginGroup := api.Group("/classes")
		{
			classNotLoginController.ClassNotLoginRoutes(classNotLoginGroup)
		}
		assignmentNotLoginGroup := api.Group("/classes/:class_id/assignments")
		{
			assignmentControllerNotLogin.AssignmentNotLoginRoutes(assignmentNotLoginGroup)
		}

		systemsGroup := api.Group("/systems")
		{
			systemsController.SystemsRoutes(systemsGroup)
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
