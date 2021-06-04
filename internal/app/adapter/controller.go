package adapter

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/airdb/uic/internal/app/domain/service"
	"github.com/airdb/uic/internal/version"

	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"

	"github.com/gin-gonic/gin"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/airdb/uic/docs"
)

var (
	user = service.Bitbank{}
	// parameterRepository = repository.Parameter{}
	// orderRepository     = repository.Order{}
)

// Handler serverless faas handler.
func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	return GinFaas.ProxyWithContext(ctx, req)
}

var GinFaas *ginAdapter.GinFaas

// @title UIC Swagger API
// @version 1.0
// @description User Information Center
// @termsOfService https://airdb.io/terms/

// @contact.name Discussion
// @contact.url https://github.com/airdb/airdb.github.io/discussions

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes https
// @host scf.baobeihuijia.com
// @BasePath /test/
func NewRouter() {
	fmt.Printf("Gin start")

	r := gin.Default()

	projectPath := "/uic"
	r.LoadHTMLGlob("internal/app/adapter/view/*")

	r.GET(projectPath, DefaultRoot)

	APIGroup := r.Group(projectPath)
	APIGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("./doc.json")))
	APIGroup.GET("/", DefaultRoot)

	APIGroup.GET("/login", index)
	APIGroup.GET("/user/query", getUser)
	APIGroup.GET("/user/signup", signup)
	APIGroup.POST("/user/signin", signin)
	APIGroup.POST("/user/signout", signout)

	if os.Getenv("env") == "dev" {
		defaultAddr := ":8081"
		err := r.Run(defaultAddr)
		if err != nil {
			panic(err)
		}

		return
	}

	GinFaas = ginAdapter.New(r)

	faas.Start(Handler)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"token": "68b329da9893e34099c7d8ad5cb9c940",
	})
}

func DefaultRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"deploy_info": version.GetBuildInfo(),
	})
}

// @Security ApiKeyAuth
// @Description get struct array by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /user/query [get]
func getUser(c *gin.Context) {
	user := service.GetUser(user) // Dependency Injection
	c.JSON(http.StatusOK, user)
}

// @Security ApiKeyAuth
// @Description Sign up
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /user/signup [post]
func signup(c *gin.Context) {
	user := service.GetUser(user) // Dependency Injection
	c.JSON(http.StatusOK, user)
}

// @Security ApiKeyAuth
// @Description Sign in
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /user/signin [post]
func signin(c *gin.Context) {
	user := service.GetUser(user) // Dependency Injection
	c.JSON(http.StatusOK, user)
}

// @Security ApiKeyAuth
// @Description Sign out
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /user/signout [post]
func signout(c *gin.Context) {
	user := service.GetUser(user) // Dependency Injection
	c.JSON(http.StatusOK, user)
}
