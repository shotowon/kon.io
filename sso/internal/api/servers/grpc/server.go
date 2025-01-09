package grpc

import (
	"context"

	auth "github.com/shotowon/kon.io/api/gen/go/sso/v1"
	"google.golang.org/grpc"
)

type apiServer struct {
	auth.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	auth.RegisterAuthServer(gRPC, &apiServer{})
}

func (a *apiServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	panic("implement me")
}

func (a *apiServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	panic("implement me")
}

func (a *apiServer) IsAdmin(ctx context.Context, req *auth.IsAdminRequest) (*auth.IsAdminResponse, error) {
	panic("implement me")
}
