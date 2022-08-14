package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"

	pb "user/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(bc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: bc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {

	log.Info(11)
	user, err := s.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := pb.CreateUserReply{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}

	return &userInfoRsp, nil
}

//func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
//	return &pb.UpdateUserReply{}, nil
//}
//func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
//	return &pb.DeleteUserReply{}, nil
//}
//func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
//	return &pb.GetUserReply{}, nil
//}
//func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
//	return &pb.ListUserReply{}, nil
//}
