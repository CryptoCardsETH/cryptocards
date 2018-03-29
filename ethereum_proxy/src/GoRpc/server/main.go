package main

import (
	cb "GoRpc/contracts/cardbase"
	pb "GoRpc/rpcServer"
	"big"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type server struct{}

type ContractAddr struct {
	ContractName    string `json:"name,omitempty"`
	Addr            string `json:"address,omitempty"`
	TransactionHash string `json:"transaction_hash,omitempty"`
}

func (s *server) GetBlank(ctx context.Context, in *pb.BlankRequest) (*pb.BlankReply, error) {
	return &pb.BlankReply{Message: "Hello" + in.Name}, nil
}

func (s *server) GetCardsByOwner(ctx context.Context, in *pb.CardsRequest) (*pb.CardsReply, error) {
	conn := Dial()
	addr := GetCardbaseAddr()
	cardBase := InitializeCardbaseContract(addr, conn)
	cardsOwned, err := cardBase.TokensOfOwner(nil, addr)
	if err != nil {
		log.Fatalf("Error on TokensOfOwner")
		panic(err)
	}
	log.Printf("%v", cardsOwned)

	return &pb.CardsReply{CreationTime: 5555, BattleCooldownEnd: 5555, CreationBattleID: 10, CurrentBattleID: 10, Attributes: "maybe tokens man idk"}, nil
}

func Dial() *ethclient.Client {
	//TODO Remove hardcoded address, move to .env
	conn, err := ethclient.Dial("localhost:8545")
	if err != nil {
		log.Fatalf("Could not dial Ganache: %v", err)
		panic(err)
	}
	log.Printf("Connected to Ganache")
	return conn
}

func GetCardbaseAddr() string {
	//TODO: Just pretend this works for right now
	resp, err := http.Get("localhost:4000/cardbase")
	if err != nil {
		log.Fatalf("Could not contact Laravel")
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Body delivered by Laravel is somehow broken")
		panic(err)
	}
	var contractLocation = new(ContractAddr)
	jsonerr := json.Unmarshal(body, &contractLocation)
	if jsonerr != nil {
		log.Fatalf("Json was not as expected format")
		panic(err)
	}
	return contractLocation.Addr
}

func InitializeCardbaseContract(addr string, conn *ethclient.Client) *cb.CardOwnership {
	contract, err := cb.NewCardOwnership(common.HexToAddress(addr), conn)
	if err != nil {
		log.Fatalf("Could not grab contract from the blockchain")
		panic(err)
	}
	return contract
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
