package biz

import (
	"context"
	"errors"
	"nblog/pkg/util"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"

	v1 "nblog/api/user/service/v1"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id       string
	Username string
	Password string
	Nickname string
	Gender   v1.Gender
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id string, username string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	nUid := util.NewNUid()
	user := &User{
		Id:       strconv.FormatUint(nUid.Generate(), 10),
		Username: in.Username,
		Password: in.Password,
		Nickname: in.Nickname,
		Gender:   in.Gender,
	}
	_, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		// TODO handle error
		return nil, err
	}
	return &v1.CreateUserReply{
		Id: user.Id,
	}, nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	result, err := uc.repo.GetUser(ctx, in.GetId(), in.GetUsername())
	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		User: &v1.UserDTO{
			Id:       result.Id,
			Username: result.Username,
			Password: result.Password,
			Nickname: result.Nickname,
			Gender:   result.Gender,
		},
	}, nil
}
