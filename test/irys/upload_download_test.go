package irys

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"web3-music-platform/app/song/repository/irys_cli"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()

	err := irys_cli.NewIrysClient("641943fa6c3fab18fed274c2b3194f0d71383ecfb9a58b2d70188e693c245510", "https://sepolia.infura.io/v3/6b7f3960da564093ade725a5b8e6d3b4")
	if err != nil {
		log.Print(err)
	}

	file, err := os.Open("./small_music.mp3")
	if err != nil {
		t.Fatalf("Failed to open audio file: %v", err)
	}

	bytes, err := ioutil.ReadAll(file)

	fileInfo, err := file.Stat()

	fileSize := int(fileInfo.Size())

	price, err := irys_cli.IrysClientInstance.GetPrice(ctx, fileSize)

	balance, err := irys_cli.IrysClientInstance.GetBalance(ctx)

	t.Logf("price:%v", price)
	t.Logf("balance:%v", balance)
	t.Logf("fileSize:%v", fileSize)

	txId, err := irys_cli.Upload(irys_cli.IrysClientInstance, ctx, bytes)

	if err != nil {
		t.Fatalf("Failed to upload audio file: %v", err)
	} else {
		t.Logf("Uploaded audio file with transaction ID: %v", txId)
	}

	data, err := irys_cli.Download(irys_cli.IrysClientInstance, ctx, txId)

	if err != nil {
		t.Fatalf("Failed to download audio file: %v", err)
	} else {
		t.Logf("Downloaded audio file: %v", len(data))
	}
}
