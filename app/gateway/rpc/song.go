package rpc

import (
	"context"
	"web3-music-platform/app/song/repository/db/model"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/utils"
)

func Upload(ctx context.Context, song model.Song, content []byte) error {
	_, err := SongService.UploadSong(ctx, &pb.CreateSongRequest{
		Song:    utils.ToSongModel(&song),
		Content: content,
	})
	if err != nil {
		return err
	}
	return nil
}

func FindSongs(ctx context.Context, userAddr string) ([]*model.Song, error) {

	var songs []*model.Song

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
