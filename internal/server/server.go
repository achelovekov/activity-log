package server

import (
	"context"
	"log"

	api "github.com/achelovekov/activity-log/api/v1"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type grpcServer struct {
	api.UnimplementedActivityLogServer
	Activities *Activities
}

func (s *grpcServer) Insert(ctx context.Context, activity *api.Activity) (*api.InsertResponse, error) {
	id, err := s.Activities.Insert(activity)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := api.InsertResponse{Id: int32(id)}
	return &res, nil
}

func NewGRPCServer() *grpc.Server {
	var activities *Activities
	var err error
	if activities, err = NewActivities(); err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	server := grpcServer{
		Activities: activities,
	}
	api.RegisterActivityLogServer(gRPCServer, &server)
	return gRPCServer
}
