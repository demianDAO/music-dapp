package contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/ini.v1"
	"testing"
	"web3-music-platform/pkg/contract/sm"
)

func TestListenReleasedSong(t *testing.T) {
	// 读取配置文件
	cfg, err := ini.Load("config.ini")

	if err != nil {
		t.Fatalf("Failed to read config file: %v", err)
	}

	//rpc := cfg.Section("smart_contract").Key("BSC_RPC").String()
	wsRpc := cfg.Section("smart_contract").Key("BSC_RPC_WS").String()
	songNFTTradeAddrStr := cfg.Section("smart_contract").Key("SongNFTTradeAddr").String()

	if err != nil {
		t.Fatalf("Failed to map config file: %v", err)
	}

	// 连接到以太坊客户端
	client, err := ethclient.Dial(wsRpc)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 合约地址
	songNFTTradeAddr := common.HexToAddress(songNFTTradeAddrStr)

	// 创建合约实例
	songNFTTrade, err := sm.NewSongNFTTradeFilterer(songNFTTradeAddr, client)

	if err != nil {
		t.Fatalf("Failed to instantiate a SongNFTTrade contract: %v", err)
	}

	// 监听 ReleasedSong 事件
	eventChan := make(chan *sm.SongNFTTradeReleasedSong)
	sub, err := songNFTTrade.WatchReleasedSong(&bind.WatchOpts{Context: context.Background()}, eventChan, nil)
	if err != nil {
		t.Fatalf("Failed to subscribe to ReleasedSong events: %v", err)
	}

	// 处理事件
	select {
	case err := <-sub.Err():
		t.Fatalf("Subscription error: %v", err)
	case event := <-eventChan:
		fmt.Printf("ReleasedSong event detected:\n")
		fmt.Printf("  Singer: %s\n", event.Singer.Hex())
		fmt.Printf("  Price: %s\n", event.Price.String())
		fmt.Printf("  TokenURI: %s\n", event.TokenURI)
		fmt.Printf("  TokenId: %s\n", event.TokenId.String())
		fmt.Printf("  Amount: %s\n", event.Amount.String())
	}
}
