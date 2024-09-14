// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/qualifications/qualifications.proto

package qualifications

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func NewTestQualificationsServiceClient(srv QualificationsServiceServer) (QualificationsServiceClient, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	server := grpc.NewServer()
	RegisterQualificationsServiceServer(server, srv)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Printf("error serving test grpc server: %v", err)
		}
	}()

	conn, err := grpc.NewClient("",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to test grpc server: %v", err)
	}

	go func() {
		<-ctx.Done()
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		server.Stop()
	}()

	client := NewQualificationsServiceClient(conn)
	return client, ctx, cancel
}
