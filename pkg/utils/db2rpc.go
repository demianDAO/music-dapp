package utils

import (
	song_service_model "web3-music-platform/internal/app/song/models"
	user_service_model "web3-music-platform/internal/app/user/models"
	"web3-music-platform/pkg/grpc/pb"
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
		Id:       uint64(post.ID),
		UserAddr: post.UserAddr,
		Content:  post.Content,
	}
}

func ToDBPost(post *pb.PostModel) *user_service_model.Post {
	return &user_service_model.Post{
		UserAddr: post.UserAddr,
		Content:  post.Content,
	}
}

func ToSong(model *pb.SongModel) *song_service_model.Song {
	return &song_service_model.Song{
		Title:      model.Title,
		ArtistAddr: model.ArtistAddr,
		UserAddr:   model.UserAddr,
		Overview:   model.Overview,
		NFTAddr:    model.NftAddr,
		TokenID:    model.TokenId,
		TxId:       model.TxId,
	}
}

func ToSongModel(song *song_service_model.Song) *pb.SongModel {
	return &pb.SongModel{
		Id:         uint64(song.ID),
		Title:      song.Title,
		ArtistAddr: song.ArtistAddr,
		UserAddr:   song.UserAddr,
		Overview:   song.Overview,
		NftAddr:    song.NFTAddr,
		TokenId:    song.TokenID,
		TxId:       song.TxId,
	}
}
