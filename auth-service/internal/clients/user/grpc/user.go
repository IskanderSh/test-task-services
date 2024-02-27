package grpcclient

import (
	"context"
	"fmt"
	"log/slog"

	userv1 "github.com/IskanderSh/test-task-protos/gen/go/user-service"
	"github.com/IskanderSh/test-task-services/auth-service/internal/config"
	"github.com/IskanderSh/test-task-services/auth-service/internal/lib/error/wrapper"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	log        *slog.Logger
	grpcClient userv1.UserClient
}

func New(
	log *slog.Logger,
	config config.UserService,
) (*UserClient, error) {
	const op = "clients.user.grpc.user"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(config.GRPCRetries)),
		grpcretry.WithPerRetryTimeout(config.GRPCTimeout)}

	logOpts := []grpclog.Option{grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent)}

	cc, err := grpc.Dial(fmt.Sprintf(":%d", config.GRPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		))
	if err != nil {
		return nil, wrapper.Wrap(op, err)
	}

	grpcClient := userv1.NewUserClient(cc)

	return &UserClient{
		log:        log,
		grpcClient: grpcClient,
	}, nil
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
