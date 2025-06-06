package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/orenvadi/auth-grpc/grpc/auth"
	"github.com/orenvadi/auth-grpc/internal/storage/surrdb"

	// "github.com/orenvadi/auth-grpc/internal/services/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	db         *surrdb.Storage
	port       int
}

// type Stoppable interface {
// 	Stop() error
// }

// New creates new gRPCServer app.
func New(log *slog.Logger, authService authgrpc.Auth, confCodeService authgrpc.ConfCodes, db *surrdb.Storage, port int) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer, authService, confCodeService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		db:         db,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op: ", op),
		slog.Int("port: ", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr ", l.Addr().String()))

	reflection.Register(a.gRPCServer)

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Graceful Shutdown needed to stop receiving new connections
// and to finish current ones in order to break nothing

// Stop stops gRPCServer.
func (a *App) Stop() {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op: ", op))
	log.Info("stopping gRPC server", slog.Int("port", a.port))

	defer a.db.Closer.Close()
	log.Info("DB connection closed")

	a.gRPCServer.GracefulStop()
}
