package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"nblog/app/user/service/internal/biz"

	pb "nblog/api/user/service/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.UserUseCase

	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return s.uc.CreateUser(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return s.uc.GetUser(ctx, req)
}
