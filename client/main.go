package main

import (
	filev1 "connect-go-practice/gen/proto"
	"connect-go-practice/gen/proto/filev1connect"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
)

func main() {

	client := filev1connect.NewFileServiceClient(http.DefaultClient, "http://localhost:8080")
	fmt.Printf("サーバーとのコネクションスタート！:%v", client)

	callListFiles(client)

}

func callListFiles(client filev1connect.FileServiceClient) {

	res, err := client.ListFiles(context.Background(), connect.NewRequest(&filev1.ListFilesRequest{}))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Msg.GetFilenames())

}
