package main

import (
	filev1 "connect-go-practice/gen/proto"
	"connect-go-practice/gen/proto/filev1connect"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type server struct{}

func (s *server) ListFiles(ctx context.Context, req *connect.Request[filev1.ListFilesRequest]) (*connect.Response[filev1.ListFilesResponse], error) {

	fmt.Println("ListFiles was invoked")

	dir := "/Users/fujiwarahideyuki/project/connect-go-practice/strage"

	paths, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string

	for _, path := range paths {

		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := connect.NewResponse(&filev1.ListFilesResponse{
		Filenames: filenames,
	})

	return res, nil

}

func main() {

	a := &server{}

	mux := http.NewServeMux()

	path, handler := filev1connect.NewFileServiceHandler(a)

	mux.Handle(path, handler)

	fmt.Println("サーバー起動中。。。")

	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))

}
