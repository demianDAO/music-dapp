package services

import (
	"context"
	"errors"
	"log"
	"sync"
	"web3-music-platform/internal/app/user/models"
	"web3-music-platform/internal/app/user/repositories"

	"web3-music-platform/pkg/grpc/pb"
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
	var post = new(models.Post)
	post.UserAddr = request.GetUserAddr()
	post.Content = request.GetContent()
	err := repositories.NewPostDao(ctx).CreatePost(post)
	if err != nil {
		return err
	}
	response.Post = utils.ToRPCPost(post)
	return nil
}

func (us *UserService) GetPosts(ctx context.Context, request *pb.GetPostsRequest, response *pb.GetPostsResponse) error {
	var postModels []*pb.PostModel
	posts, err := repositories.NewPostDao(ctx).GetPosts(request.GetUserAddr())
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
	user, err := repositories.NewUserDao(ctx).FindUserByAddress(request.GetAddress())
	log.Print("find user:", user)

	if err != nil {
		return err
	} else if user.Name == "" {
		return errors.New("user not exist")
	}

	response.User = utils.ToRPCUser(user)
	return nil
}

func (us *UserService) UserRegister(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	var userModel = new(pb.UserModel)
	userModel.Name = request.GetName()
	userModel.Address = request.GetAddress()
	user := utils.ToDBUser(userModel)
	err := repositories.NewUserDao(ctx).CreateUser(user)
	if err != nil {
		return err
	}
	response.User = utils.ToRPCUser(user)
	return nil
}
