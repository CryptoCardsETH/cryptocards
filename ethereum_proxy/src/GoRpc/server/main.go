package main

import (
	conts "GoRpc/contracts"
	pb "GoRpc/rpcServer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type server struct{}

func (s *server) GetBlank(ctx context.Context, in *pb.BlankRequest) (*pb.BlankReply, error) {
	log.Println(in)
	return &pb.BlankReply{Message: "Hello" + in.Name}, nil
}

func (s *server) GetCardsByOwner(ctx context.Context, in *pb.CardsRequest) (*pb.CardsReply, error) {
	return &pb.CardsReply{CreationTime: 5555, BattleCooldownEnd: 5555, CreationBattleID: 10, CurrentBattleID: 10, Attributes: "maybe tokens man idk"}, nil
}
func (s *server) PerformECRecover(ctx context.Context, in *pb.ECRecoverRequest) (*pb.ECRecoverReply, error) {
	address := in.Address
	signed := in.Signed
	plaintext := in.Plaintext
	log.Printf("a: %v, s: %v, p: %v", address, signed, plaintext)

	status := verifySig(
		address,
		signed,
		[]byte("CryptoCards"),
	)

	return &pb.ECRecoverReply{Success: status}, nil
}

func (s *server) RequestBattleGroupInfo(ctx context.Context, in *pb.BattleGroupInfoRequest) (*pb.BattleGroupInfoReply, error) {

	client := getEthClientConnection()

	coreContractAddr := common.HexToAddress(in.Contract.Address)
	log.Printf("core addr, %v\n", coreContractAddr.Hex())

	core, err := conts.NewCryptoCardsCore(coreContractAddr, client)

	a := &bind.CallOpts{}
	battleGroupsContractAddr, err := core.BattleGroupContract(a)

	battleGroupsContract, err := conts.NewBattleGroups(battleGroupsContractAddr, client)

	opt := &bind.FilterOpts{}
	it, err := battleGroupsContract.BattleGroupsFilterer.FilterNewBattleGroup(opt)
	if err != nil {
		log.Fatalf("error filtering events: %v", err)
	}

	var groups []*pb.BattleGroupInfo

	notEmpty := true
	for notEmpty {
		notEmpty = it.Next()
		if notEmpty {
			log.Println("event log:", it)
			newBattleGroupEvent := it.Event
			log.Printf("owner: %v\n", newBattleGroupEvent.Owner.Hex())
			log.Printf("BG id: %v\n", newBattleGroupEvent.BattleGroupID)
			log.Printf("Cards: %v\n", newBattleGroupEvent.Cards)

			cardsField := make([]uint64, len(newBattleGroupEvent.Cards))
			for i, card := range newBattleGroupEvent.Cards {
				cardsField[i] = card.Uint64()
			}

			groups = append(
				groups,
				&pb.BattleGroupInfo{
					OwnerAddress: newBattleGroupEvent.Owner.Hex(),
					Id:           newBattleGroupEvent.BattleGroupID.Uint64(),
					Cards:        cardsField,
				},
			)
		}
	}

	return &pb.BattleGroupInfoReply{Items: groups}, nil

}

func getEthClientConnection() *ethclient.Client {
	RPCHOST, _ := os.LookupEnv("RPC_HOST")
	RPCPORT, _ := os.LookupEnv("RPC_PORT")

	client, err := ethclient.Dial("http://" + RPCHOST + ":" + RPCPORT)
	if err != nil {
		log.Fatalf("Could not dial RPC server: %v", err)
		panic(err)
	}
	return client

}

func main() {
	log.Printf("hello")

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "50051"
	}
	port = ":" + port

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
}
