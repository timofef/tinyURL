package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery"
	server_proto "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/server"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/repository"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/usecase"
	"github.com/timofef/tinyURL/internal/tinyURL/logger"
	"github.com/timofef/tinyURL/internal/tinyURL/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

type ServerInterceptor struct {
	Logger *logger.Logger
}

func (s *ServerInterceptor) logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	//md, _ := metadata.FromIncomingContext(ctx)

	reqId := rand.Uint64()

	s.Logger.Logger = s.Logger.Logger.WithFields(logrus.Fields{
		"requestId": reqId,
		"method":    info.FullMethod,
		//"context":   md,
		"request":        req,
		"response":       resp,
		"error":          err,
		"execution_time": time.Since(start),
	})

	s.Logger.LogInfo("Entry Point")

	reply, err := handler(ctx, req)

	s.Logger.LogInfo("USER Interceptor")

	return reply, err
}

func main() {
	// Init logger
	logger.MainLogger = &logger.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}
	logger.MainLogger.Logger.Logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	// Parse storage option flag
	useInMemoryStorage := flag.Bool("in-memory", false, "use in-memory storage")
	flag.Parse()

	// Init repository
	var tinyUrlRepository repository.IRepository
	if *useInMemoryStorage {
		tinyUrlRepository = &repository.TinyUrlInMemRepository{
			Mux: sync.RWMutex{},
			DB:  make(map[string]string),
		}
		logger.MainLogger.LogInfo("Using in-memory storage")
	} else {
		db, err := repository.InitPostgres(os.Getenv("DB"))
		if err != nil {
			logger.MainLogger.LogError("Can't connect to database: " + err.Error())
			return
		}
		tinyUrlRepository = &repository.TinyUrlSqlRepository{
			DB: db,
		}
		logger.MainLogger.LogInfo("Using PostgreSQL storage")
	}

	// Init tinyUrlUsecase
	baseUrl := "http://default.base.url.com/"
	if val := os.Getenv("BASE_URL"); val != "" {
		baseUrl = val
	}
	logger.MainLogger.LogInfo("Base URL: " + baseUrl)
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
		grpclog.Fatal("Failed to listen: " + err.Error())
	}
	ServerInterceptor := ServerInterceptor{logger.MainLogger}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(ServerInterceptor.logger))
	server_proto.RegisterTinyUrlServiceServer(grpcServer, &tinyUrlHandler)

	// Serve
	logger.MainLogger.LogInfo("Started server on localhost:5555")
	err = grpcServer.Serve(listen)
	if err != nil {
		grpclog.Fatalf("Server fail: %v", err)
	}
}
