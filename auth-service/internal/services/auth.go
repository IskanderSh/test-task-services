package services

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"

	userv1 "github.com/IskanderSh/test-task-protos/gen/go/user-service"
	"github.com/IskanderSh/test-task-services/auth-service/internal/config"
	"github.com/IskanderSh/test-task-services/auth-service/internal/lib/error/wrapper"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	log        *slog.Logger
	userClient userv1.UserClient
}

func NewAuthService(log *slog.Logger, config config.UserService) (*AuthService, error) {
	const op = "internal.services.New"

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

	return &AuthService{log: log, userClient: grpcClient}, nil
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")

	nameRegexp  = regexp.MustCompile(`[A-Z][a-z]{5,30}`).MatchString
	emailRegexp = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString
)

func (s *AuthService) SignUp(ctx context.Context, name, surname, password, email string) (string, error) {
	return nil, nil
}

func (s *AuthService) SignIn(ctx context.Context, email, password string) (string, error) {
	const op = "services.auth.SignIn"

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", email))

	if !emailRegexp(email) {
		log.Warn("invalid email")

		return "", errors.Wrap(ErrInvalidCredentials, "invalid email")
	}

	s.userClient.Get(ctx)
}
