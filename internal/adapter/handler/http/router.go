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
		profile.GET("/profile-details", profileHandler.GetProfiles)
		profile.PATCH("/update-profile/:id", profileHandler.UpdateProfileByUserID)
	}

	lab := api.Group("/lab").Use(authMiddleware(token))
	{
		lab.POST("/create", labHandler.InsertLab)
		lab.GET("/get/:id", labHandler.GetLabByID)
		lab.PATCH("/update/:id", labHandler.UpdateLab)
	}

	labTest := api.Group("/lab-test").Use(authMiddleware(token))
	{
		labTest.POST("/create-department", labTestsHandler.CreateDepartment)
		labTest.POST("/create-panel", labTestsHandler.CreatePanel)
		labTest.POST("/create-catalog", labTestsHandler.CreateTestCatalog)
		labTest.POST("/create-panel-component", labTestsHandler.CreatePanelComponent)
		labTest.POST("/create-test-parameter", labTestsHandler.CreateTestParameter)
		labTest.POST("/create-reference", labTestsHandler.CreateReferenceRange)
		labTest.GET("/list-department", labTestsHandler.GetDepartments)
		labTest.GET("/list-panel/:id", labTestsHandler.GetPanelsByDepartmentID)
		labTest.GET("/list-catalog/:id", labTestsHandler.GetTestCatalogByDepartmentID)
		labTest.GET("/list-panel-catalog/:id", labTestsHandler.GetPanelComponentsByPanelID)
		labTest.GET("/list-test-parameter/:id", labTestsHandler.GetTestParametersByTestCatalogID)
		labTest.GET("/list-reference/:id", labTestsHandler.GetReferenceRangesByTestParameterID)
		labTest.PATCH("/update-department", labTestsHandler.UpdateDepartment)
		labTest.PATCH("/update-panel", labTestsHandler.UpdatePanel)
		labTest.PATCH("/update-catalog", labTestsHandler.UpdateTestCatalog)
		labTest.PATCH("/update-test-parameter", labTestsHandler.UpdateTestParameter)
		labTest.PATCH("/update-reference", labTestsHandler.UpdateReferenceRange)
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
