package api

// import (
// 	"context"
// )

// type api struct {
// }

// func NewApi() proto.APIServer {
// 	return &api{}
// }

// func (a *api) Login(ctx context.Context, req *proto.LoginRequest) (resp *proto.LoginResponse, err error) {
// 	return &proto.LoginResponse{
// 		StatusCode:    "ACCPET",
// 		ReasonCode:    "",
// 		ReasonMessage: "",
// 		Jwt:           "jwt ...",
// 	}, nil
// }

// func (a *api) Liveness(ctx context.Context, req *proto.EmptyRequestResponse) (resp *proto.EmptyRequestResponse, err error) {
// 	return &proto.EmptyRequestResponse{}, nil
// }

// func (a *api) Readness(ctx context.Context, req *proto.EmptyRequestResponse) (resp *proto.EmptyRequestResponse, err error) {
// 	return &proto.EmptyRequestResponse{}, nil
// }

// func (a *api) mustEmbedUnimplementedAPIServer() {
// }
