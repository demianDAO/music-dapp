package listener

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/sirupsen/logrus"
	"web3-music-platform/internal/app/song/repositories"
	"web3-music-platform/pkg/contract/sm"
	"web3-music-platform/pkg/rdb"
)

var logIns = log.WithFields(log.Fields{
	"module": "listener",
})

type Listener struct {
	songRepo             *repositories.SongRepository
	songNFTTradeFilterer *sm.SongNFTTradeFilterer
}

func NewListener(songRepo *repositories.SongRepository, songNFTTradeFilterer *sm.SongNFTTradeFilterer) *Listener {
	return &Listener{
		songRepo:             songRepo,
		songNFTTradeFilterer: songNFTTradeFilterer,
	}
}

func (l *Listener) WatchReleasedSong() {
	logIns.Print("Starting ReleasedSong listener...")
	releasedSongEventChan := make(chan *sm.SongNFTTradeReleasedSong)
	releasedSongEventSub, err := l.songNFTTradeFilterer.WatchReleasedSong(&bind.WatchOpts{Context: context.Background()}, releasedSongEventChan, nil)
	if err != nil {
		logIns.Fatalf("Failed to subscribe to ReleasedSong events: %v", err)
	}
	defer releasedSongEventSub.Unsubscribe()

	for {
		select {
		case err := <-releasedSongEventSub.Err():
			logIns.Printf("Subscription error: %v", err)
			return
		case event := <-releasedSongEventChan:
			logIns.Printf("ReleasedSong event detected in listener.WatchReleasedSong:\n")
			logIns.Printf("  Singer: %s\n", event.Singer.Hex())
			logIns.Printf("  Price: %s\n", event.Price.String())
			logIns.Printf("  TokenURI: %s\n", event.TokenURI)
			logIns.Printf("  TokenId: %s\n", event.TokenId.String())
			logIns.Printf("  Amount: %s\n", event.Amount.String())
			err := handleReleasedSongEvent(event.Singer.String(), event.TokenId.Uint64(), l.songRepo)
			if err != nil {
				logIns.Printf("Failed to handleReleasedSongEvent in listener.WatchReleasedSong: %v", err)
			}
		}
	}
}

func handleReleasedSongEvent(artistAddr string, tokenId uint64, repo *repositories.SongRepository) error {
	logIns.Printf("Handling ReleasedSong event in handleReleasedSongEvent:\n")
	logIns.Printf("  ArtistAddr: %s\n", artistAddr)
	logIns.Printf("  TokenId: %d\n", tokenId)

	song, err := repo.GetLatestSong(artistAddr)
	if err != nil || song.TokenID != 0 {
		err := rdb.RedisInstance.CacheReleasedSongEvent(context.Background(), artistAddr, tokenId)
		if err != nil {
			return err
		}
		return err
	}

	if song.TokenID == 0 {
		logIns.Print("Fill tokenId in handleReleasedSongEvent")
		err = repo.SetTokenId(artistAddr, tokenId)
		if err != nil {
			return err
		}
	}
	return nil
}
