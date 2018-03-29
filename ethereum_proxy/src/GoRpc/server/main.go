package main

import (
	pb "GoRpc/rpcServer"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
)

type server struct{}

type Client struct {
	c *rpc.Client
}

type ContractAddr struct {
	ContractName    string `json:"name,omitempty"`
	Addr            string `json:"address,omitempty"`
	TransactionHash string `json:"transaction_hash,omitempty"`
}

func (s *server) GetBlank(ctx context.Context, in *pb.BlankRequest) (*pb.BlankReply, error) {
	return &pb.BlankReply{Message: "Hello" + in.Name}, nil
}

func (s *server) GetCardsByOwner(ctx context.Context, in *pb.CardsRequest) (*pb.CardsReply, error) {
	return &pb.CardsReply{CreationTime: 5555, BattleCooldownEnd: 5555, CreationBattleID: 10, CurrentBattleID: 10, Attributes: "maybe tokens man idk"}, nil
}

func testDial() *Client {
	//TODO Remove hardcoded address, move to .env
	conn, err := ethclient.Dial("localhost:8545")
	if err != nil {
		log.Fatalf("Could not dial Ganache: %v", err)
		panic(err)
	}
	log.Printf("Connected to Ganache")
	log.Printf(conn)
	return conn
}

func getCardbaseAddr() string {
	//TODO: Just pretend this works for right now
	resp, err := http.Get("localhost:4000/cardbase")
	if err != nil {
		log.Fatalf("Could not contact Laravel")
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.All(resp.Body)
	if err != nil {
		log.Fatalf("Body delivered by Laravel is somehow broken")
		panic(err)
	}
	var contractLocation = new(ContractAddr)
	err := json.Unmarshal(body, &contractLocation)
	if err != nil {
		log.Fatalf("Json was not as expected format")
		panic(err)
	}
	return contractLocation.Addr
}

func main() {
	log.Printf("Attempting to start Go-RPC Server")

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "50051"
	}
	port = ":" + port

	log.Printf("Attempting to connect to Ganache Test Network")
	c := testDial()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	log.Printf("RPC Server listening on %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	address := getCardbaseAddr()
}
