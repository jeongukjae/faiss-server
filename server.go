package main

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "github.com/jeongukjae/faiss-server/protos/faiss/service"
)

func runGrpcGateway(grpcEndpoint string, httpEndpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	glog.Info("Register grpc gateway server at ", grpcEndpoint)
	err := gw.RegisterFaissHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return err
	}

	handler := promhttp.Handler()
	err = mux.HandlePath("GET", "/metrics", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		handler.ServeHTTP(w, r)
	})
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	glog.Info("Serve http server at ", httpEndpoint)
	return http.ListenAndServe(httpEndpoint, mux)
}

func RunServer(faissPath string, grpcEndpoint string, httpEndpoint string) {
	if faissPath == "" {
		glog.Fatal("You should pass faiss index path (-faiss_index option)")
	}

	if err := runGrpcServer(grpcEndpoint, faissPath); err != nil {
		glog.Fatal(err)
	}

	if err := runGrpcGateway(grpcEndpoint, httpEndpoint); err != nil {
		glog.Fatal(err)
	}
}
