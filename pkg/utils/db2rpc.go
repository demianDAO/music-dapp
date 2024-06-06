package utils

import (
	song_service_model "web3-music-platform/app/song/repository/db/model"
	user_service_model "web3-music-platform/app/user/repository/db/model"
	"web3-music-platform/idl/pb"
)

func ToRPCUser(user *user_service_model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:      uint64(user.ID),
		Name:    user.Name,
		Address: user.Address,
	}
}

func ToDBUser(user *pb.UserModel) *user_service_model.User {
	return &user_service_model.User{
		Name:    user.Name,
		Address: user.Address,
	}
}

func ToRPCPost(post *user_service_model.Post) *pb.PostModel {
	return &pb.PostModel{
		Id:         uint64(post.ID),
		AuthorAddr: post.AuthorAddr,
		Content:    post.Content,
	}
}

func ToDBPost(post *pb.PostModel) *user_service_model.Post {
	return &user_service_model.Post{
		AuthorAddr: post.AuthorAddr,
		Content:    post.Content,
	}
}

// 将 SongModel 转换为 Song 的接收器函数
func ToSong(model *pb.SongModel) *song_service_model.Song {
	return &song_service_model.Song{
		Title:         model.Title,
		ArtistAddress: model.ArtistAddress,
		Overview:      model.Overview,
		NFTAddress:    model.NftAddress,
		TokenID:       model.TokenId,
		IrysTxId:      model.IrysTxId,
	}
}

// 将 Song 转换为 SongModel 的接收器函数
func ToSongModel(song *song_service_model.Song) *pb.SongModel {
	return &pb.SongModel{
		Id:            uint64(song.ID),
		Title:         song.Title,
		ArtistAddress: song.ArtistAddress,
		Overview:      song.Overview,
		NftAddress:    song.NFTAddress,
		TokenId:       song.TokenID,
		IrysTxId:      song.IrysTxId,
	}
}
