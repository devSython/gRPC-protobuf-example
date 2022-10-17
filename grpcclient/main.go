package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "proto-example/proto"
)

const (
	address = "localhost:50051"
)

var obj pb.Person
