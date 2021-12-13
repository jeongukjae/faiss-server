package main

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jeongukjae/faiss-server/faiss"
	gw "github.com/jeongukjae/faiss-server/protos/faiss/service"
)

type faissServer struct {
	gw.UnimplementedFaissServer

	Index *faiss.FaissIndex
}

func (s *faissServer) GetMetadata(ctx context.Context, in *gw.EmptyMessage) (*gw.GetMetadataResponse, error) {
	return &gw.GetMetadataResponse{
		IndexName:  s.Index.Path,
		Dimension:  s.Index.Dimension,
		MetricType: gw.GetMetadataResponse_MetricType(s.Index.MetricType),
		Ntotal:     s.Index.GetNtotal(),
	}, nil
}

func (s *faissServer) Search(ctx context.Context, in *gw.SearchRequest) (*gw.SearchResponse, error) {
	if err := checkVectorDimension(in.Vectors, in.NumVectors, s.Index.Dimension); err != nil {
		return nil, err
	}
	if in.TopK == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "top k argument = 0")
	}

	results := s.Index.Search(in.NumVectors, in.TopK, in.Vectors)
	return &gw.SearchResponse{
		Ids:       results.Ids,
		Distances: results.Distances,
	}, nil
}

func (s *faissServer) AddVectors(ctx context.Context, in *gw.AddVectorsRequest) (*gw.AddVectorsResponse, error) {
	if err := checkVectorDimension(in.Vectors, in.NumVectors, s.Index.Dimension); err != nil {
		return nil, err
	}

	ids, err := s.Index.AddVectors(in.NumVectors, in.Vectors)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gw.AddVectorsResponse{Ids: ids}, nil
}

func (s *faissServer) AddVectorsWithIds(ctx context.Context, in *gw.AddVectorsWithIdsRequest) (*gw.EmptyMessage, error) {
	if err := checkVectorDimension(in.Vectors, in.NumVectors, s.Index.Dimension); err != nil {
		return nil, err
	}

	numElements := int32(len(in.Ids))
	if numElements != in.NumVectors {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"num elements of ids(%d) != num vectors(%d)",
			numElements, in.NumVectors,
		)
	}

	err := s.Index.AddVectorsWithIds(in.NumVectors, in.Vectors, in.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s. This method is not supported by all indexes.", err.Error())
	}
	return &gw.EmptyMessage{}, nil
}

func (s *faissServer) RemoveVectors(ctx context.Context, in *gw.RemoveVectorsRequest) (*gw.RemoveVectorsResponse, error) {
	if len(in.Ids) == 0 {
		return nil, status.Error(codes.InvalidArgument, "num elements of ids = 0")
	}
	numRemoved, err := s.Index.RemoveVectors(in.Ids)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gw.RemoveVectorsResponse{NumRemoved: numRemoved}, nil
}

func (s *faissServer) Reload(ctx context.Context, in *gw.EmptyMessage) (*gw.EmptyMessage, error) {
	indexPath := s.Index.Path
	glog.Info("Reload triggered, remove old index...")
	s.Index.Free()

	glog.Info("Load new index from ", indexPath, "...")
	newIndex, err := faiss.LoadIndex(indexPath)
	if err != nil {
		glog.Fatal("Cannot load index again!!")
		return nil, status.Errorf(codes.Internal, "Cannot reload faiss")
	}
	s.Index = newIndex
	glog.Info("Reloaded index successfully")
	return &gw.EmptyMessage{}, nil
}

func checkVectorDimension(vectors []float32, numVectors int32, dimension int32) error {
	numElements := int32(len(vectors))
	if numElements != numVectors*dimension {
		return status.Errorf(
			codes.InvalidArgument,
			"num elements of vector(%d) != num vectors(%d) * dimension(%d)",
			numElements, numVectors, dimension,
		)
	}

	if numElements == 0 {
		return status.Errorf(codes.InvalidArgument, "num elements of vector = 0")
	}

	return nil
}
