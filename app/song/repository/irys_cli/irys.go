package irys_cli

import (
	"context"
	"github.com/Ja7ad/irys"
	"github.com/Ja7ad/irys/currency"
	"github.com/goccy/go-json"
	"io"
	"log"
)

var IrysClientInstance *irys.Client

func NewIrysClient(privateKey, rpc string) error {
	log.Print("Initializing Irys client...", privateKey)

	currency, err := currency.NewEthereum(privateKey, rpc)
	//currency, err := currency.NewArweave(privateKey)

	if err != nil {
		return err
	}

	irysInterface, err := irys.New(irys.DefaultNode1, currency, true)
	if err != nil {
		return err
	}

	IrysClientInstance = irysInterface.(*irys.Client)

	return nil
}

func Upload(ic *irys.Client, ctx context.Context, data []byte) (string, error) {
	//deadline, ok := ctx.Deadline()
	//if ok {
	//	log.Printf("Context deadline is set at %v\n", deadline)
	//} else {
	//	log.Println("No deadline is set for the context")
	//}

	tx, err := ic.Upload(ctx, data)
	bytes, err := json.MarshalIndent(tx, "", "  ")
	log.Print("Uploaded audio file to Irys network with tx ID:", string(bytes))
	if err != nil {
		return "", err
	}

	return tx.ID, nil
}

func Download(ic *irys.Client, ctx context.Context, index string) ([]byte, error) {
	file, err := ic.Download(ctx, index)
	if err != nil {
		return nil, err
	}
	defer file.Data.Close()

	data, err := io.ReadAll(file.Data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
