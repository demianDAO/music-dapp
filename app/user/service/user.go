package service

import (
	"context"
	"log"
	"sync"
	"web3-music-platform/app/user/repository/db/dao"
	"web3-music-platform/app/user/repository/db/model"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/utils"
)

var UserServiceInstance *UserService
var UserSrvOnce sync.Once

type UserService struct {
}

func NewUserService() *UserService {
	UserSrvOnce.Do(func() {
		UserServiceInstance = &UserService{}
	})
	return UserServiceInstance
}

func (us *UserService) CreatePost(ctx context.Context, request *pb.PostCreateRequest, response *pb.PostCreateResponse) error {
	var post = new(model.Post)
	post.AuthorAddr = request.GetAuthorAddr()
	post.Content = request.GetContent()
	err := dao.NewPostDao(ctx).CreatePost(post)
	if err != nil {
		return err
	}
	response.Post = utils.ToRPCPost(post)
	return nil
}

func (us *UserService) GetPostsByAuthor(ctx context.Context, request *pb.GetPostsByAuthorRequest, response *pb.GetPostsByAuthorResponse) error {
	var postModels []*pb.PostModel
	posts, err := dao.NewPostDao(ctx).GetPostsByAuthor(request.GetAuthorAddr())
	if err != nil {
		return err
	}
	for _, post := range posts {
		postModels = append(postModels, utils.ToRPCPost(post))
	}
	response.Post = postModels
	return nil
}

func (us *UserService) UserLogin(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	log.Print("enter user login")
	user, err := dao.NewUserDao(ctx).FindUserByAddress(request.GetAddress())
	log.Print("find user:", user)

	if err != nil {
		return err
	}

	response.User = utils.ToRPCUser(user)
	return nil
}

func (us *UserService) UserRegister(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	var userModel = new(pb.UserModel)
	userModel.Name = request.GetName()
	userModel.Address = request.GetAddress()
	user := utils.ToDBUser(userModel)
	err := dao.NewUserDao(ctx).CreateUser(user)
	if err != nil {
		return err
	}
	response.User = utils.ToRPCUser(user)
	return nil
}
