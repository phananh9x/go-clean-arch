package server

import (
	"context"
	"crypto"
	"fmt"
	"go-clean-arch/pkg/infra"
	"go-clean-arch/pkg/jwt"
	"go-clean-arch/service/repository"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"

	"go-clean-arch/config"
)

//Server ...
type Server struct {
	httpServer   *http.Server
	router       *gin.Engine
	metricRouter *gin.Engine
	cfg          *config.AppConfig
	verifier     jwt.IVerifier
	signer       jwt.ISigner
	rateLimiter  *limiter.Limiter
	prometheus   *ginprometheus.Prometheus
}

// NewServer construct server
func NewServer(cfg *config.AppConfig) (*Server, error) {
	router := gin.New()
	metricRouter := gin.New()

	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))
	router.Use(gin.Recovery())

	s := &Server{
		metricRouter: metricRouter,
		router:       router,
		cfg:          cfg,
	}

	return s, nil
}

// Init server
func (s *Server) Init() {
	ctx := context.Background()
	//init redis
	redisClient, err := infra.InitRedis(s.cfg.Redis)
	if err != nil {
		zap.S().Errorf("Init redis error: %v", err)
		panic(err)
	}
	//init redis
	mgoClient, err := infra.InitMongo(ctx, s.cfg.Mgo)
	if err != nil {
		zap.S().Errorf("Init redis error: %v", err)
		panic(err)
	}
	if err != nil {
		zap.S().Errorf("Init limiter error: %v", err)
		panic(err)
	}
	//init jwt verifier
	verifier, err := jwt.NewVerifier(s.cfg.Authen.APIAuthenticator.JWTPublicKeyBase64, jwt.DefaultSigningMethodAlg)
	if err != nil {
		zap.S().Errorf("Init jwt verifier error: %v", err)
		panic(err)
	}
	s.verifier = verifier
	//init jwt signer
	zap.S().Infof("Init jwt signer with private key: %v", s.cfg.Authen.APIAuthenticator.JWTPrivateKeyBase64)
	signer, err := jwt.NewSigner(s.cfg.Authen.APIAuthenticator.JWTPrivateKeyBase64, jwt.DefaultSigningMethodAlg, crypto.SHA256)
	if err != nil {
		zap.S().Errorf("Init jwt signer error: %v", err)
		panic(err)
	}
	s.signer = signer
	//init repos
	repo := repository.NewRepo(mgoClient, s.cfg, redisClient)
	//init list service funcs
	domains := s.initDomains(ctx, repo, redisClient)
	s.initCORS()
	s.initHealthCheck()
	s.initRouters(domains)
	s.initSwagger()
}

// Close help cleanup um-managed resources
func (s *Server) Close() {
	_ = s.httpServer.Close()
}

//ListenHTTP ...
func (s *Server) ListenHTTP() error {
	address := fmt.Sprintf(":%s", os.Getenv("PORT"))

	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}
	//Start listen metrics
	go func() {
		port := fmt.Sprintf(":%s", os.Getenv("PORT_METRICS"))
		s.metricRouter.Run(port)
	}()

	zap.S().Infof("starting http server at port %v ...", os.Getenv("PORT"))

	return s.httpServer.ListenAndServe()
}
