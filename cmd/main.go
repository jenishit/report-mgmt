// @title           LIS Backend API
// @version         1.0
// @description     Laboratory Information System API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".
package main

import (
	"context"
	"fmt"
	"log/slog"

	"os"

	auth "github.com/jenish-brainztechs/go-backend/internal/adapter/auth/jwt"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/config"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres/repository"
	"github.com/jenish-brainztechs/go-backend/internal/core/services"
)

func main() {
	config, err := config.New() //Creating a new configuration for the application
	if err != nil {             //if there is some error then print and log the error and exit from the application
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Database has been initializerd and connected successfully", "db", config.DB.Connection)

	tokenService, err := auth.New(config.Token)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}

	roleRepo := repository.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)
	roleHandler := http.NewRoleHandler(roleService)

	profileRepo := repository.NewProfileRepository(db)
	profileService := services.NewProfileService(profileRepo)
	profileHandler := http.NewProfileHandler(profileService)

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, tokenService)
	authHandler := http.NewAuthHandler(authService)
	userService := services.NewUserService(userRepo, roleService, profileService)
	userHandler := http.NewUsersHandler(userService)

	labRepo := repository.NewLabRepository(db)
	labService := services.NewLabService(labRepo)
	labHandler := http.NewLabHandler(labService)

	labTestsRepo := repository.NewLabTestsRepository(db)
	labTestsService := services.NewLabTestsService(labTestsRepo)
	lasTestsHandler := http.NewLabTestsHandler(labTestsService)

	docRepo := repository.NewDoctorRepository(db)
	docService := services.NewDoctorService(docRepo)
	docHandler := http.NewDoctorHandler(docService)

	ptRepo := repository.NewPatientRepository(db)
	ptService := services.NewPatientService(ptRepo)
	ptHandler := http.NewPatientHandler(ptService)

	visitRepo := repository.NewVisitRepository(db)
	visitService := services.NewVisitService(visitRepo)
	visitHandler := http.NewVisitHandler(visitService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := http.NewOrderHandler(orderService)

	resultRepo := repository.NewResultRepository(db)
	resultService := services.NewResultService(resultRepo)
	resultHandler := http.NewResultHandler(resultService)

	reportRepo := repository.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := http.NewReportHandler(reportService)

	router, err := http.NewRouter(
		config,
		tokenService,
		*roleHandler,
		*userHandler,
		*profileHandler,
		*authHandler,
		*labHandler,
		*lasTestsHandler,
		*docHandler,
		*ptHandler,
		*visitHandler,
		*orderHandler,
		*resultHandler,
		*reportHandler,
	)

	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)

	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
