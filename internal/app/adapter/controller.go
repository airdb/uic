package adapter

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/airdb/sailor/gin/handlers"
	"github.com/airdb/sailor/sliceutil"
	"github.com/airdb/sailor/version"
	"github.com/airdb/uic/internal/app/domain/service"

	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"

	"github.com/gin-gonic/gin"
	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"

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
// @termsOfService https://airdb.github.io/terms.html

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
	r.Use(
		handlers.Jsonifier(),
	)

	project := sliceutil.LastStringWithSplit(version.Repo, "/")

	projectPath := "/" + project
	r.LoadHTMLGlob("internal/app/adapter/view/*")

	r.GET(projectPath, DefaultRoot)

	APIGroup := r.Group(projectPath)
	handlers.RegisterSwagger(APIGroup)

	APIGroup.GET("/", redirect)

	APIGroup.GET("/login", index)
	APIGroup.GET("/login/github", loginGithub)
	APIGroup.GET("/login/gitee", loginGitee)

	APIGroup.Any("/user/query", getUser)
	APIGroup.GET("/user/signup", signup)

	APIGroup.POST("/user/login", signin)
	APIGroup.POST("/user/signin", signin)

	APIGroup.POST("/user/logout", signout)
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

// @Security ApiKeyAuth
// @Description get struct array by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /login/token [get]
func index(c *gin.Context) {
	config := service.GetOauthConfig()

	loginURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s",
		"https://github.com/login/oauth/authorize",
		config.ClientID,
		config.RedirectURL,
		"user:email read:org",
		config.State,
	)

	fmt.Print(config.ID)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"login_url": loginURL,
		"token":     "68b329da9893e34099c7d8ad5cb9c940",
	})
}

func DefaultRoot(c *gin.Context) {
	handlers.SetResp(c, version.GetBuildInfo())
	/*
		c.JSON(http.StatusOK, gin.H{
			"deploy_info": version.GetBuildInfo(),
		})
	*/
}

func redirect(c *gin.Context) {
	fmt.Println(os.Environ())
	fmt.Println("xx", c.Request.RequestURI)
	redirectURL := "/" + os.Getenv("ENV") + c.Request.RequestURI + "/swagger/index.html"
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
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
	// c.JSON(http.StatusOK, user)
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// 	"data":    user,
	// })

	handlers.SetResp(c, &user)
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

// @Security ApiKeyAuth
// @Description Reference: https://docs.github.com/en/developers/apps/building-oauth-apps/authorizing-oauth-apps
// @Tags login
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /login/github [get]
func loginGithub(c *gin.Context) {
	user := service.GetUser(user) // Dependency Injection

	redirectURL := "https://noah.airdb.io/?#/callback?token=" + user.Token

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

// @Security ApiKeyAuth
// @Description Reference: https://gitee.com/api/v5/oauth_doc#/
// @Tags login
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /login/gitee [get]
func loginGitee(c *gin.Context) {
	user := service.GetUser(user) // Dependency Injection
	// c.JSON(http.StatusOK, user)

	redirectURL := "https://noah.airdb.io/index.html?token=" + user.Token
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
