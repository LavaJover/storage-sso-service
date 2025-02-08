package main

import (
	"log"
	"log/slog"
	"net"

	config "github.com/LavaJover/storage-sso-service/sso-service/internal/config"
	"github.com/LavaJover/storage-sso-service/sso-service/internal/db"
	repo "github.com/LavaJover/storage-sso-service/sso-service/internal/repository"
	"github.com/LavaJover/storage-sso-service/sso-service/internal/server"
	"github.com/LavaJover/storage-sso-service/sso-service/internal/service"
	ssopb "github.com/LavaJover/storage-sso-service/sso-service/proto/gen"
	"google.golang.org/grpc"
)

func main(){

	// Loading config
	cfg := config.MustLoad()

	// Init db layer
	db := db.InitDB(cfg.Dsn)

	// Init repos layer
	userRepo := repo.UserRepo{
		DB: db,
	}

	// Init service layer
	ssoService := service.SSOService{
		UserRepo: &userRepo,
	}

	// Init server layer
	ssoServer := server.SSOServer{
		SSOService: &ssoService,
	}
	grpcServer := grpc.NewServer()
	ssopb.RegisterAuthServiceServer(grpcServer, &ssoServer)

	// Starting server
	listener, err := net.Listen("tcp", cfg.Host + ":" + cfg.Port)

	if err != nil{
		log.Fatalf("failed to start server: %v\n")
	}

	slog.Info("gRPC server running on " + cfg.Host + ":" + cfg.Port)

	if err := grpcServer.Serve(listener); err != nil{
		log.Fatalf("failed to serve gRPC server: %v\n", err)
	}
}