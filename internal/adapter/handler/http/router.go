package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/jenish-brainztechs/go-backend/docs"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/config"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.Container,
	token port.TokenService,
	roleHandler RoleHandler,
	userHandler UserHandler,
	profileHandler ProfileHandler,
	authHandler AuthHandler,
	labHandler LabHandler,
	labTestsHandler LabTestsHandler,
	docHandler DoctorHandler,
	ptHandler PatientHandler,
	visitHandler VisitHandler,
	orderHandler OrderHandler,
	resultHandler ResultHandler,
	reportHandler ReportHandler,
) (*Router, error) {

	if config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Middleware order matters - apply in this sequence:

	// 2. For Lambda Function URL deployments, let AWS handle CORS if explicitly enabled.
	if !config.HTTP.UseFunctionURLCORS {
		allowedOrigins := config.HTTP.AllowedOrigins
		router.Use(CORSMiddleware(allowedOrigins))
	}

	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	role := api.Group("/role")
	{
		role.POST("/create", roleHandler.CreateRole)
	}

	spc := api.Group("/special")
	{
		spc.POST("/user/create", userHandler.CreateUser)
	}

	profile := api.Group("/profile").Use(authMiddleware(token))
	{
		profile.GET("/getme", profileHandler.GetProfileByID)
		profile.PATCH("/update-profile/:id", profileHandler.UpdateProfileByUserID)
	}

	admin := api.Group("/admin")
	admin.Use(authMiddleware(token), adminMiddleware())

	user := admin.Group("/user")
	{
		user.POST("/create", userHandler.CreateUser)
	}

	profiles := admin.Group("/profile").Use(authMiddleware(token))
	{
		profiles.GET("/getme", profileHandler.GetProfileByID)
		profiles.GET("/profile-details", profileHandler.GetProfiles)
		profiles.PATCH("/update-profile/:id", profileHandler.UpdateProfileByUserID)
	}

	labs := admin.Group("/lab")
	{
		labs.POST("", labHandler.InsertLab)
		labs.GET("", labHandler.GetAllLabs)
		labs.GET("/:id", labHandler.GetLabByID)
		labs.PATCH("/:id", labHandler.UpdateLab)
	}
	labTests := admin.Group("/lab-test")
	{
		labTests.POST("/create-department", labTestsHandler.CreateDepartment)
		labTests.POST("/create-panel", labTestsHandler.CreatePanel)
		labTests.POST("/create-catalog", labTestsHandler.CreateTestCatalog)
		labTests.POST("/create-panel-component", labTestsHandler.CreatePanelComponent)
		labTests.POST("/create-test-parameter", labTestsHandler.CreateTestParameter)
		labTests.POST("/create-reference", labTestsHandler.CreateReferenceRange)
		labTests.PATCH("/update-department", labTestsHandler.UpdateDepartment)
		labTests.PATCH("/update-panel", labTestsHandler.UpdatePanel)
		labTests.PATCH("/update-catalog", labTestsHandler.UpdateTestCatalog)
		labTests.PATCH("/update-test-parameter", labTestsHandler.UpdateTestParameter)
		labTests.PATCH("/update-reference", labTestsHandler.UpdateReferenceRange)
	}

	lab := api.Group("/lab").Use(authMiddleware(token))
	{
		lab.GET("/get/:id", labHandler.GetLabByID)
	}

	labTest := api.Group("/lab-test").Use(authMiddleware(token))
	{

		labTest.GET("/list-department", labTestsHandler.GetDepartments)
		labTest.GET("/list-panel/:id", labTestsHandler.GetPanelsByDepartmentID)
		labTest.GET("/list-catalog/:id", labTestsHandler.GetTestCatalogByDepartmentID)
		labTest.GET("/list-panel-catalog/:id", labTestsHandler.GetPanelComponentsByPanelID)
		labTest.GET("/list-test-parameter/:id", labTestsHandler.GetTestParametersByTestCatalogID)
		labTest.GET("/list-reference/:id", labTestsHandler.GetReferenceRangesByTestParameterID)

	}

	doc := admin.Group("/doctor")
	{
		doc.POST("", docHandler.CreateDoctor)
		doc.PATCH("/:id", docHandler.UpdateDoctor)
	}

	doctor := api.Group("/doctor").Use(authMiddleware(token))
	{
		doctor.GET("/:id", docHandler.GetDoctorByID)
		doctor.GET("", docHandler.GetDoctors)
	}

	pt := api.Group("/patient").Use(authMiddleware(token))
	{
		pt.POST("", ptHandler.CreatePatient)
		pt.GET("/:id", ptHandler.GetPatientByID)
		pt.GET("", ptHandler.GetPatients)
		pt.PATCH("/:id", ptHandler.UpdatePatient)
	}

	visit := api.Group("/visit").Use(authMiddleware(token))
	{
		visit.POST("", visitHandler.CreateVisit)
		visit.GET("/:id", visitHandler.GetVisitByID)
		visit.PATCH("/:id", visitHandler.UpdateVisitByID)
		visit.GET("/patient/:id", visitHandler.GetVisitByPatientID)
	}

	order := api.Group("/order").Use(authMiddleware(token))
	{
		order.POST("", orderHandler.CreateOrder)
		order.GET("/:id", orderHandler.GetOrderByID)
		order.PATCH("/:id", orderHandler.UpdateOrder)
		order.GET("/visit/:visit_id", orderHandler.GetOrdersByVisitID)
	}

	result := api.Group("/result").Use(authMiddleware(token))
	{
		result.POST("", resultHandler.CreateResult)
		result.GET("/:id", resultHandler.GetResultByID)
		result.PATCH("/:id", resultHandler.UpdateResult)
		result.GET("/order/:order_id", resultHandler.GetResultsByOrderID)
	}

	report := api.Group("/report").Use(authMiddleware(token))
	{
		report.POST("", reportHandler.CreateReport)
		report.GET("/:id", reportHandler.GetReportByID)
		report.PATCH("/:id", reportHandler.UpdateReport)
		report.GET("/visit/:visit_id", reportHandler.GetReportsByVisitID)
		report.POST("/:id/print", reportHandler.CreateReportPrint)
		report.GET("/:id/prints", reportHandler.GetReportPrints)
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
