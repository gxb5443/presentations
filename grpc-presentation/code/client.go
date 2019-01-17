package main

// START OMIT
import (
	"context"
	"log"
	"time"

	pb "go.jet.network/phaser/core-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Assign(ctx, &pb.AssignmentRequest{DrOrpheusHeader: "Dr. O: Phaser!"})
	if err != nil {
		log.Fatalf("could not assign: %v", err)
	}
}

// END OMIT
