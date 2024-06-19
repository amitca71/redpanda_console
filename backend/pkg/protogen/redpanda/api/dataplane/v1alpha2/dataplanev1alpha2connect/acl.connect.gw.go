// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha2/acl.proto

package dataplanev1alpha2connect

import (
	context "context"
	fmt "fmt"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	connect_gateway "go.vallahaye.net/connect-gateway"

	v1alpha2 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha2"
)

// ACLServiceGatewayServer implements the gRPC server API for the ACLService service.
type ACLServiceGatewayServer struct {
	v1alpha2.UnimplementedACLServiceServer
	listACLs   connect_gateway.UnaryHandler[v1alpha2.ListACLsRequest, v1alpha2.ListACLsResponse]
	createACL  connect_gateway.UnaryHandler[v1alpha2.CreateACLRequest, v1alpha2.CreateACLResponse]
	deleteACLs connect_gateway.UnaryHandler[v1alpha2.DeleteACLsRequest, v1alpha2.DeleteACLsResponse]
}

// NewACLServiceGatewayServer constructs a Connect-Gateway gRPC server for the ACLService service.
func NewACLServiceGatewayServer(svc ACLServiceHandler, opts ...connect_gateway.HandlerOption) *ACLServiceGatewayServer {
	return &ACLServiceGatewayServer{
		listACLs:   connect_gateway.NewUnaryHandler(ACLServiceListACLsProcedure, svc.ListACLs, opts...),
		createACL:  connect_gateway.NewUnaryHandler(ACLServiceCreateACLProcedure, svc.CreateACL, opts...),
		deleteACLs: connect_gateway.NewUnaryHandler(ACLServiceDeleteACLsProcedure, svc.DeleteACLs, opts...),
	}
}

func (s *ACLServiceGatewayServer) ListACLs(ctx context.Context, req *v1alpha2.ListACLsRequest) (*v1alpha2.ListACLsResponse, error) {
	return s.listACLs(ctx, req)
}

func (s *ACLServiceGatewayServer) CreateACL(ctx context.Context, req *v1alpha2.CreateACLRequest) (*v1alpha2.CreateACLResponse, error) {
	return s.createACL(ctx, req)
}

func (s *ACLServiceGatewayServer) DeleteACLs(ctx context.Context, req *v1alpha2.DeleteACLsRequest) (*v1alpha2.DeleteACLsResponse, error) {
	return s.deleteACLs(ctx, req)
}

// RegisterACLServiceHandlerGatewayServer registers the Connect handlers for the ACLService "svc" to
// "mux".
func RegisterACLServiceHandlerGatewayServer(mux *runtime.ServeMux, svc ACLServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1alpha2.RegisterACLServiceHandlerServer(context.TODO(), mux, NewACLServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}
