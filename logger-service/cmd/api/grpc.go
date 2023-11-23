package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/raphaelmb/go-ms-logger-service/data"
	"github.com/raphaelmb/go-ms-logger-service/logs"
	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	if err := l.Models.LogEntry.Insert(logEntry); err != nil {
		res := &logs.LogResponse{Result: "Failed"}
		return res, err
	}

	res := &logs.LogResponse{Result: "Logged."}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()
	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})
	log.Printf("gRPC server started on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
