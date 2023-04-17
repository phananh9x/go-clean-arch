package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-clean-arch/pkg/ginwrapper"
	"go-clean-arch/pkg/middleware"
	customerHdl "go-clean-arch/service/domain/customers/delivery/http"
	"go-clean-arch/service/domain/customers/usecase"
	"go-clean-arch/service/repository"
	"net/http"
	"os"
)

type domains struct {
	customer usecase.ICustomerUseCase
}

func (s *Server) initCORS() {
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowAllOrigins = true
	// corsConfig.AllowHeaders = []string{
	// 	"*",
	// 	"Origin",
	// 	"Content-Length",
	// 	"Content-Type",
	// 	"Authorization",
	// 	"X-Access-Token",
	// 	"X-Google-Access-Token",
	// }
	// s.router.Use(cors.New(corsConfig))
}

// @title [API Document] - Golang API
// @version 1.0.0

// @securityDefinitions.apikey AccessToken
// @in header
// @name x-access-token

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:3070
// @BasePath  /v1
// @termsOfService http://swagger.io/terms/
// InitSwagger ...
func (s *Server) initSwagger() {
	if s.cfg.Environment != "prod" {
		//docs.SwaggerInfo.Host = s.cfg.Swagger.Host
		//docs.SwaggerInfo.BasePath = s.cfg.Swagger.BasePath
		s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (s *Server) initHealthCheck() {
	version := os.Getenv("SERVICE_VERSION")

	pingHandler := ginwrapper.WithContext(func(ctx *ginwrapper.Context) {
		ctx.JSONData(http.StatusOK, map[string]interface{}{
			"version": version,
		})
	})

	s.router.GET("/health-check", pingHandler)
	s.router.HEAD("/health-check", pingHandler)
}

func (s *Server) initDomains(ctx context.Context, repo repository.IRepo, redisClient *redis.Client) *domains {
	customerUseCase := usecase.NewCustomerUseCase(s.cfg, s.signer, repo)
	return &domains{
		customer: customerUseCase,
	}
}

func (s *Server) initRouters(domains *domains) {
	// Handler
	customerHandler := customerHdl.NewCustomerHandler(domains.customer)
	// Grouping routes
	router := s.router.Group("v1")
	customerHandler.CustomerAPIRoutePublic(router)

	router.Use(middleware.APIAuthentication(s.verifier))
	customerHandler.CustomerAPIRoute(router)
}
