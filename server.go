package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jeongukjae/faiss-server/faiss"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gw "github.com/jeongukjae/faiss-server/protos/faiss/service"
)

var (
	index *faiss.FaissIndex
)

type faissServer struct {
	gw.UnimplementedFaissServer
}

func (s *faissServer) GetMetadata(ctx context.Context, in *gw.EmptyMessage) (*gw.GetMetadataResponse, error) {
	return &gw.GetMetadataResponse{
		IndexName:  index.Path,
		Dimension:  index.Dimension,
		MetricType: gw.GetMetadataResponse_MetricType(index.MetricType),
		Ntotal:     index.GetNtotal(),
	}, nil
}

func (s *faissServer) Search(ctx context.Context, in *gw.SearchRequest) (*gw.SearchResponse, error) {
	numElements := int32(len(in.Vectors))
	if numElements != in.NumVectors*index.Dimension {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"num elements of vector(%d) != num vectors(%d) * dimension(%d)",
			numElements, in.NumVectors, index.Dimension,
		)
	}

	results := index.Search(in.NumVectors, in.TopK, in.Vectors)
	return &gw.SearchResponse{
		Ids:       results.Ids,
		Distances: results.Distances,
	}, nil
}

func runGrpcServer(endpoint string) error {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return errors.WithStack(err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	gw.RegisterFaissServer(s, &faissServer{})
	grpc_prometheus.Register(s)
	go func() { log.Fatalln(s.Serve(lis)) }()
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

func RunServer(faissPath string, grpcEndpoint string, httpEndpoint string) {
	if faissPath == "" {
		glog.Fatal("You should pass faiss index path (-faiss_index option)")
	}

	var err error
	glog.Info("Loading faiss index from ", faissPath)
	index, err = faiss.LoadIndex(faissPath)
	if err != nil {
		glog.Fatal(err)
	}
	defer index.Free()

	if err = runGrpcServer(grpcEndpoint); err != nil {
		glog.Fatal(err)
	}

	if err = runGrpcGateway(grpcEndpoint, httpEndpoint); err != nil {
		glog.Fatal(err)
	}
}
