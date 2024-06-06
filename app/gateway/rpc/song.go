package rpc

import (
	"context"
	"web3-music-platform/idl/pb"
)

func Upload(ctx context.Context, req *pb.CreateSongRequest) (*pb.CreateSongResponse, error) {
	res, err := SongService.UploadSong(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}
