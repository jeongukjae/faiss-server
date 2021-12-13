package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/jeongukjae/faiss-server/faiss"
	gw "github.com/jeongukjae/faiss-server/protos/faiss/service"
)

func runGrpcServer(endpoint string, faissPath string, withReloadMethod bool) error {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return errors.WithStack(err)
	}

	glog.Info("Loading faiss index from ", faissPath)
	index, err := faiss.LoadIndex(faissPath)
	if err != nil {
		return errors.WithStack(err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	fs := &faissServer{Index: index, WithReloadMethod: withReloadMethod}
	gw.RegisterFaissServer(s, fs)
	grpc_prometheus.Register(s)
	go func() {
		defer fs.Index.Free()
		log.Fatalln(s.Serve(lis))
	}()
	glog.Info("Serve grpc server at ", endpoint)

	return nil
}

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

func RunServer(faissPath string, grpcEndpoint string, httpEndpoint string, withReloadMethod bool) {
	if faissPath == "" {
		glog.Fatal("You should pass faiss index path (-faiss_index option)")
	}

	if err := runGrpcServer(grpcEndpoint, faissPath, withReloadMethod); err != nil {
		glog.Fatal(err)
	}

	if err := runGrpcGateway(grpcEndpoint, httpEndpoint); err != nil {
		glog.Fatal(err)
	}
}
