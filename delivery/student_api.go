package delivery

import (
	"github.com/gin-gonic/gin"
	"go-mongo/db"
	"go-mongo/model"
	"go-mongo/repository"
	"go-mongo/usecase"
	"net/http"
	"strconv"
)

type IStudentApi interface {
	createStudent(c *gin.Context)
	getStudentByName(c *gin.Context)
	getAllStudentWithPagination(c *gin.Context)
}
type StudentApi struct {
	publicRoute *gin.RouterGroup
	useCase     *usecase.StudentUseCase
}

func (s *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
		})
		return
	}
	registeredStudent, err := s.useCase.RegisterStudent(student)
	c.JSON(http.StatusOK, gin.H{
		"data": registeredStudent,
	})
}

func (s *StudentApi) getStudentByName(c *gin.Context) {
	name := c.Param("name")
	student, err := s.useCase.FindStudentByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func (s *StudentApi) getAllStudentWithPagination(c *gin.Context) {
	skip := c.Query("skip")
	limit := c.Query("limit")
	convertSkip, _ := strconv.Atoi(skip)
	convertLimit, _ := strconv.Atoi(limit)
	student, err := s.useCase.FindStudentsWithPagination(convertSkip, convertLimit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func NewStudentApi(publicRoute *gin.RouterGroup, resource *db.Resource) *StudentApi {
	userRoute := publicRoute.Group("/students")
	studentRepo := repository.NewStudentRepository(resource)
	studentApi := StudentApi{
		publicRoute: userRoute,
		useCase:     usecase.NewStudentUseCase(studentRepo),
	}
	studentApi.initRouter(userRoute)
	return &studentApi
}

func (s *StudentApi) initRouter(studentRoute *gin.RouterGroup) {
	studentRoute.GET("", s.getAllStudentWithPagination)
	studentRoute.GET("/:name", s.getStudentByName)
	studentRoute.POST("", s.createStudent)
	studentRoute.GET("/paginate", s.getAllStudentWithPagination)
}
