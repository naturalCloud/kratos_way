package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// User 定义返回数据结构体
type User struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday string
	Gender   string
	Role     int
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}
