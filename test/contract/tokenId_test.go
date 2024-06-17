package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"testing"
	"web3-music-platform/config"
	"web3-music-platform/pkg/contract"
)

func TestGetTokenId(t *testing.T) {
	os.Setenv("APP_ENV", "dev")
	config.Init()
	t.Logf("config.BSC_RPC %v", config.BSC_RPC)
	err := contract.NewGethClient()
	if err != nil {
		//log.Print("NewGethClient error ", err.Error())
		panic(err)
	}
	tokenId, err := contract.SongNFT.CurrentID(&bind.CallOpts{Pending: true})
	if err != nil {
		t.Errorf("CurrentID err = %v", err)
	}
	t.Log("CurrentID = ", tokenId)

}
