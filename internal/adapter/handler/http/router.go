package http

import (
	"github.com/gin-gonic/gin"
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
) (*Router, error) {

	if config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	role := api.Group("/role")
	{
		role.POST("/create", roleHandler.CreateRole)
	}

	user := api.Group("/user")
	{
		user.POST("/create", userHandler.CreateUser)
	}

	profile := api.Group("/profile").Use(authMiddleware(token))
	{
		profile.GET("/getme", profileHandler.GetProfileByID)
		profile.PATCH("/update-profile/:id", profileHandler.UpdateProfileByUserID)
	}

	admin := api.Group("/admin")
	admin.Use(authMiddleware(token), adminMiddleware())

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

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
