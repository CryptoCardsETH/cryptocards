package main

import (
	conts "GoRpc/contracts"
	pb "GoRpc/rpcServer"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/big"
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
func getCoreContractInstance(a *pb.CoreContractAddress) *conts.CryptoCardsCore {
	client := getEthClientConnection()
	coreContractAddr := common.HexToAddress(a.Address)
	core, err := conts.NewCryptoCardsCore(coreContractAddr, client)
	if err != nil {
		log.Fatalf("Error getting Core Contract instance %v", err)
	}
	return core

}

//Returns public + private keypair, based on environment vars
func getKeypairForTransactions() (common.Address, *ecdsa.PrivateKey) {
	pubStr, _ := os.LookupEnv("ADMIN_PUBKEY")
	privStr, _ := os.LookupEnv("ADMIN_PRIVKEY")

	pub := common.HexToAddress(pubStr)
	priv, _ := crypto.HexToECDSA(privStr)

	return pub, priv
}

// Creates a Card (next incremental CardID) on the blockchain, with the owner set as specified
func (s *server) CreateCard(ctx context.Context, in *pb.CreateCardRequest) (*pb.BlankReply, error) {
	ownerAddress := in.OwnerAddress
	log.Printf("Creating a card for owner: %v\n", ownerAddress)

	ganachePublicKey, ganachePrivateKey := getKeypairForTransactions()

	signer := func(signer types.Signer, address common.Address, txn *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(txn, signer, ganachePrivateKey)
	}

	sesh := &conts.CryptoCardsCoreSession{
		Contract: getCoreContractInstance(in.CoreAddress),
		TransactOpts: bind.TransactOpts{
			From:   ganachePublicKey,
			Signer: signer,
		},
	}

	a, err := sesh.CreateCard(common.HexToAddress(ownerAddress), big.NewInt(3))
	if err != nil {
		log.Fatalf("Error creating card %v", err)
	}
	log.Printf("CreateCard txn: %v", a.Hash().Hex())

	return &pb.BlankReply{Message: a.Hash().Hex()}, nil
}

// Get `NewCard` events
func (s *server) RequestCardInfo(ctx context.Context, in *pb.CardInfoRequest) (*pb.CardInfoReply, error) {
	core := getCoreContractInstance(in.Contract)
	it, err := core.CryptoCardsCoreFilterer.FilterNewCard(&bind.FilterOpts{})
	if err != nil {
		log.Fatalf("error filtering events: %v", err)
	}
	var cards []*pb.CardInfo
	for it.Next() {
		newCard := it.Event
		cards = append(cards, &pb.CardInfo{
			OwnerAddress:     newCard.Owner.Hex(),
			Id:               newCard.CardID.Uint64(),
			CreationBattleId: newCard.CreationBattleID.Uint64(),
			Attributes:       newCard.Attributes.Uint64(),
		})
	}
	return &pb.CardInfoReply{Items: cards}, nil
}
func (s *server) RequestBattleGroupInfo(ctx context.Context, in *pb.BattleGroupInfoRequest) (*pb.BattleGroupInfoReply, error) {

	core := getCoreContractInstance(in.Contract)
	battleGroupsContractAddr, err := core.BattleGroupContract(&bind.CallOpts{})

	client := getEthClientConnection()
	battleGroupsContract, err := conts.NewBattleGroups(battleGroupsContractAddr, client)

	it, err := battleGroupsContract.BattleGroupsFilterer.FilterNewBattleGroup(&bind.FilterOpts{})
	if err != nil {
		log.Fatalf("error filtering events: %v", err)
	}

	var groups []*pb.BattleGroupInfo

	for it.Next() {
		log.Println("NewBattleGroup Event log:")
		newBattleGroupEvent := it.Event
		log.Printf("\towner: %v\n", newBattleGroupEvent.Owner.Hex())
		log.Printf("\tBG id: %v\n", newBattleGroupEvent.BattleGroupID)
		log.Printf("\tCards: %v\n", newBattleGroupEvent.Cards)

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
