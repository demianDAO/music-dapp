package services

import (
	"context"
	"web3-music-platform/internal/app/song/models"
	"web3-music-platform/pkg/grpc/pb"
)

func UploadSong(ctx context.Context, song models.Song, content []byte, tokenUri string, amount, price uint64) error {
	_, err := SongService.UploadSong(ctx, &pb.CreateSongReq{
		Title:      song.Title,
		ArtistAddr: song.ArtistAddr,
		Overview:   song.Overview,
		Content:    content,
		TokenUri:   tokenUri,
		Amount:     amount,
		Price:      price,
	})
	if err != nil {
		return err
	}
	return nil
}

func FindSongs(ctx context.Context, userAddr string) ([]*pb.FindSongsByAddrRes_SongInfo, error) {
	res, err := SongService.FindSongs(ctx, &pb.FindSongsByAddrReq{
		Addr: userAddr,
	})
	if err != nil {
		return nil, err
	}
	return res.GetSongInfos(), nil
}

func DownloadSong(ctx context.Context, txId string) ([]byte, error) {

	res, err := SongService.DownloadSong(ctx, &pb.DownloadSongReq{
		TxId: txId,
	})
	if err != nil {
		return nil, err
	}
	return res.SongBytes, nil
}
