package services

import (
	"context"
	"web3-music-platform/internal/app/song/models"
	"web3-music-platform/pkg/grpc/pb"
	"web3-music-platform/pkg/utils"
)

func UploadSong(ctx context.Context, song models.Song, content []byte, tokenUri string) error {
	_, err := SongService.UploadSong(ctx, &pb.CreateSongRequest{
		Song:     utils.ToSongModel(&song),
		Content:  content,
		TokenUri: tokenUri,
	})
	if err != nil {
		return err
	}
	return nil
}

func FindSongs(ctx context.Context, userAddr string) ([]*models.Song, error) {

	var songs []*models.Song

	res, err := SongService.FindSongs(ctx, &pb.FindSongsRequest{
		ArtistAddress: userAddr,
	})
	if err != nil {
		return nil, err
	}
	for _, song := range res.Songs {
		songs = append(songs, utils.ToSong(song))
	}

	return songs, nil
}

func DownloadSong(ctx context.Context, txId string) ([]byte, error) {

	res, err := SongService.DownloadSong(ctx, &pb.DownloadSongRequest{
		TxId: txId,
	})
	if err != nil {
		return nil, err
	}
	return res.SongBytes, nil
}
