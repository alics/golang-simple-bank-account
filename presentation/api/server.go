package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"mondu-challenge-alihamedani/infrastructure/di"
	util "mondu-challenge-alihamedani/infrastructure/utils"
	"mondu-challenge-alihamedani/presentation"
)

type Server struct {
	config util.Config
	//db     *sql.DB
	router *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	var accountService = di.NewAccountServiceResolve(server.config)
	controller := presentation.AccountController{AccountService: accountService}

	router.Use(TokenAuthMiddleware(server.config))
	router.POST("/accounts", controller.Create)
	router.GET("/accounts/:id/balance", controller.GetBalance)
	router.PUT("/accounts/deposit", controller.Deposit)
	router.PUT("/accounts/withdraw", controller.Withdraw)

	server.router = router
}

func TokenAuthMiddleware(config util.Config) gin.HandlerFunc {
	requiredToken := config.TokenKey

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			respondWithError(c, 401, "api token required")
			return
		}
		if token != requiredToken {
			respondWithError(c, 401, "invalid api token")
			return
		}
		c.Next()
	}

}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"authorization error": message})
}
