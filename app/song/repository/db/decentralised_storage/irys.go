package decentralised_storage

import (
	"context"
	"errors"
	"github.com/Ja7ad/irys"
	"github.com/Ja7ad/irys/currency"
	"github.com/goccy/go-json"
	"io"
	"log"
)

var IrysClientInstance *IrysClient

// IrysClient 封装了与 Irys 网络交互的功能
type IrysClient struct {
	*irys.Client
}

// NewIrysClient 创建一个新的 IrysClient 实例
func NewIrysClient(privateKey string) error {
	log.Print("Initializing Irys client...", privateKey)
	currency, err := currency.NewEthereum(privateKey, "https://sepolia.infura.io/v3/6b7f3960da564093ade725a5b8e6d3b4")
	//currency, err := currency.NewArweave(privateKey)
	if err != nil {
		return err
	}

	irysInterface, err := irys.New(irys.DefaultNode2, currency, true)
	if err != nil {
		return err
	}

	// 将接口类型转换为 irys.Client 类型
	irysClientTmp, ok := irysInterface.(*irys.Client)
	if !ok {
		return errors.New("failed to convert Irys interface to irys.Client type")
	}

	// 将 irysClientTmp 赋值给全局变量 irysClient
	IrysClientInstance = &IrysClient{Client: irysClientTmp}
	return nil
}

// UploadAudioFile 上传音频文件到 Irys 网络，并返回生成的索引
func (ic *IrysClient) UploadAudioFile(ctx context.Context, audioData []byte) (string, error) {
	log.Print("Uploading audio file to Irys network...", len(audioData))
	if len(audioData) == 0 {
		return "", errors.New("audio data is empty")
	}
	if ic == nil {
		return "", errors.New("Irys client is nil")
	}

	price, err := ic.GetPrice(ctx, len(audioData))
	if err != nil {
		return "", err
	}
	log.Print("Uploading audio file to Irys network with price:", price)

	//tx, err := ic.BasicUpload(ctx, audioData)
	tx, err := ic.Upload(ctx, audioData)

	bytes, err := json.MarshalIndent(tx, "", "  ")
	log.Print("Uploaded audio file to Irys network with tx ID:", string(bytes))
	if err != nil {
		return "", err
	}

	return tx.ID, nil
}

// GetAudioFile 根据索引从 Irys 网络中获取音频文件
func (ic *IrysClient) GetAudioFile(ctx context.Context, index string) ([]byte, error) {
	file, err := ic.Download(ctx, index)
	if err != nil {
		return nil, err
	}
	defer file.Data.Close()

	audioData, err := io.ReadAll(file.Data)
	if err != nil {
		return nil, err
	}

	return audioData, nil
}
