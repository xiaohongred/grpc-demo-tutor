package main

import (
	"context"
	"log"
	"net"

	"github.com/xiaohongred/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type invoicerImp struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *invoicerImp) SayHello(ctx context.Context, req *invoicer.SayHelloReq) (*invoicer.SayHelloRsp, error) {

	return &invoicer.SayHelloRsp{
		Name: req.Name + " hi",
	}, nil
}

func (s *invoicerImp) Create(ctx context.Context, req *invoicer.CreateReq) (*invoicer.CreateRsp, error) {
	return &invoicer.CreateRsp{
		Pdf:  []byte("text"),
		Docx: []byte("text"),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	service := &invoicerImp{}
	invoicer.RegisterInvoicerServer(s, service)
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
	return
}
