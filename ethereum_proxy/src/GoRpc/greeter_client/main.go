package main

import (
	pb "GoRpc/rpcServer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.GetBlank(context.Background(), &pb.BlankRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	rep, err := c.GetCardsByOwner(context.Background(), &pb.CardsRequest{Address: "0x1"})
	if err != nil {
		log.Fatalf("could not get blank card: %v", err)
	}
	log.Printf("Template Card: %d, %d, %d, %d, %s", rep.CreationTime, rep.BattleCooldownEnd, rep.CreationBattleID, rep.CurrentBattleID, rep.Attributes)
}
