package main

import (
	"context"

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
	numElements := int32(len(in.Vectors))
	if numElements != in.NumVectors*s.Index.Dimension {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"num elements of vector(%d) != num vectors(%d) * dimension(%d)",
			numElements, in.NumVectors, s.Index.Dimension,
		)
	}

	if numElements == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "num elements of vector = 0")
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
	numElements := int32(len(in.Vectors))
	if numElements != in.NumVectors*s.Index.Dimension {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"num elements of vector(%d) != num vectors(%d) * dimension(%d)",
			numElements, in.NumVectors, s.Index.Dimension,
		)
	}

	if numElements == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "num elements of vector = 0")
	}

	ids, err := s.Index.AddVectors(in.NumVectors, in.Vectors)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gw.AddVectorsResponse{Ids: ids}, nil
}

func (s *faissServer) AddVectorsWithIds(ctx context.Context, in *gw.AddVectorsWithIdsRequest) (*gw.EmptyMessage, error) {
	numElements := int32(len(in.Vectors))
	if numElements != in.NumVectors*s.Index.Dimension {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"num elements of vector(%d) != num vectors(%d) * dimension(%d)",
			numElements, in.NumVectors, s.Index.Dimension,
		)
	}

	if numElements == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "num elements of vector = 0")
	}

	numElements = int32(len(in.Ids))
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
