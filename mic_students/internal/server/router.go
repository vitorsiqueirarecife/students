package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vitorsiqueirarecife/students/internal/config"
	handlerH "github.com/vitorsiqueirarecife/students/internal/modules/healthcheck/handler"
	serviceH "github.com/vitorsiqueirarecife/students/internal/modules/healthcheck/service"
	handlerS "github.com/vitorsiqueirarecife/students/internal/modules/student/handler"
	"github.com/vitorsiqueirarecife/students/internal/modules/student/repository"
	repositoryS "github.com/vitorsiqueirarecife/students/internal/modules/student/repository"
	serviceS "github.com/vitorsiqueirarecife/students/internal/modules/student/service"
)

func SetupRouter() *gin.Engine {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error when trying to connect to database: %v", err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	opts := repository.StudentRepositoryOptions{
		DB: db,
	}

	studentRepository := repositoryS.NewStudentRepository(opts)

	studentService := serviceS.NewStudentService(studentRepository)
	studentHandler := handlerS.NewStudentHandler(studentService)

	healthService := serviceH.NewHealthUseCase()
	healthHandler := handlerH.NewHealthHandler(healthService)

	router.GET("/health-check", healthHandler.Check)

	router.GET("/students", func(c *gin.Context) {
		students := studentHandler.GetAll(c.Writer, c.Request)
		c.JSON(200, students)
	})

	return router
}
