package main

import (
	"flag"

	"github.com/golang/glog"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc_server_endpoint", "localhost:8000", "gRPC server endpoint")
	// http server endpoint
	httpServerEndpoint = flag.String("http_server_endpoint", "localhost:8001", "http server endpoint")
	// faiss index path
	faissIndexPath = flag.String("faiss_index", "", "path of prebuilt faiss index")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	RunServer(*faissIndexPath, *grpcServerEndpoint, *httpServerEndpoint)
}
