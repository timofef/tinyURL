package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery"
	server "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/server"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/repository"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/usecase"
	"github.com/timofef/tinyURL/internal/tinyURL/logger"
	"github.com/timofef/tinyURL/internal/tinyURL/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)

	reqId := rand.Uint64()

	logger.MainLogger = logger.MainLogger.WithFields(logrus.Fields{
		"requestId":      reqId,
		"method":         info.FullMethod,
		"context":        md,
		"request":        req,
		"response":       resp,
		"error":          err,
		"execution_time": time.Since(start),
	})

	reply, err := handler(ctx, req)

	return reply, err
}

func main() {
	// Parse storage option flag
	useInMemoryStorage := flag.Bool("in-memory", false, "use in-memory storage")
	flag.Parse()

	// Init repository
	var tinyUrlRepository usecase.IRepository
	if *useInMemoryStorage {
		tinyUrlRepository = &repository.TinyUrlInMemoryRepository{
			Mux: sync.RWMutex{},
			DB:  make(map[string]string),
		}
		logger.MainLogger.Info("Using in-memory storage")
	} else {
		db, err := repository.InitPostgres(os.Getenv("DB"))
		if err != nil {
			logger.MainLogger.Fatal("Can't connect to database: " + err.Error())
		}
		tinyUrlRepository = &repository.TinyUrlSqlRepository{
			DB: db,
		}
		logger.MainLogger.Info("Using PostgreSQL storage")
	}

	// Init tinyUrlUsecase
	baseUrl := "http://default.base.url.com/"
	if val := os.Getenv("BASE_URL"); val != "" {
		baseUrl = val
	}
	logger.MainLogger.Info("Base URL: " + baseUrl)
	tinyUrlUsecase := usecase.TinyUrlUsecase{
		BaseUrl:         baseUrl,
		Repository:      tinyUrlRepository,
		GenerateTinyUrl: utils.GenerateString,
	}

	// Init handler
	tinyUrlHandler := delivery.TinyUrlHandler{
		Usecase: &tinyUrlUsecase,
	}

	// Init server
	listen, err := net.Listen("tcp", ":5555")
	if err != nil {
		logger.MainLogger.Fatal("Failed to listen: " + err.Error())
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	server.RegisterTinyUrlServiceServer(grpcServer, &tinyUrlHandler)

	// Serve
	logger.MainLogger.Info("Started server on localhost:5555")
	err = grpcServer.Serve(listen)
	if err != nil {
		logger.MainLogger.Fatal("Server fail: " + err.Error())
	}
}
