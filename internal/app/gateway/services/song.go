package services

import (
	"context"
	"web3-music-platform/pkg/grpc/pb"
)

func DiscoverySongs(ctx context.Context, limit uint64) ([]*pb.DiscoveryRes_SongShortInfo, error) {
	res, err := SongService.Discovery(ctx, &pb.DiscoveryReq{
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}
	return res.GetSongShortInfos(), nil
}

func UploadSong(ctx context.Context, title, artistAddr, overview string, content []byte) error {
	_, err := SongService.UploadSong(ctx, &pb.CreateSongReq{
		Title:      title,
		ArtistAddr: artistAddr,
		Overview:   overview,
		Content:    content,
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

//func PurchaseSong(ctx context.Context, userAddr, singerAddr string, tokenId uint64) error {
//	_, err := SongService.PurchaseSong(ctx, &pb.PurchaseSongReq{
//		TokenId:    tokenId,
//		UserAddr:   userAddr,
//		SingerAddr: singerAddr,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
