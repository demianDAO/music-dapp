package decentralised_storage

import (
	"context"
	"log"
	"testing"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()

	err := NewIrysClient("641943fa6c3fab18fed274c2b3194f0d71383ecfb9a58b2d70188e693c245510")
	//if err != nil {
	//	log.Print(err)
	//}
	//
	//txId, err := IrysClientInstance.UploadAudioFile(ctx, []byte("Hello World"))
	file, err := IrysClientInstance.GetAudioFile(ctx, "Z8hPiq_8d9s1AwGsBxKDb0fuFcjE8A8iDfk6ZkwXw4g")
	log.Print("get the audio string", string(file))

	if err != nil {
		log.Print(err)
	}

	//log.Print(txId)
}
