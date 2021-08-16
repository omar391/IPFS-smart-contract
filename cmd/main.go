package main

import (
	"context"
	"fmt"
	"ipfs_smart_contract/api"
	"ipfs_smart_contract/internal/clients"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//load env vars
	loadEnvs()

	// init the IPFS node
	node := runIPFSNode(ctx)

	// upload and get a CID string
	cid := upload(ctx, node)

	//deploy the smart contract into the Ganache instance
	scClient := &clients.SmartContractClient{
		Addr:       os.Getenv("GANACHE_ADDR"),
		Privatekey: os.Getenv("ADDRESS_PRIVATE_KEY"),
	}
	api := deploy(scClient)

	// Store our CID into blockchain
	storeCID(cid, api, scClient)

	// Retrieve our CID from blockchain
	r_cid := retriveCID(api)
	fmt.Printf("We have successfully stored and retrieved the CID: %s", r_cid)
}

func loadEnvs() {
	d, _ := os.Getwd()
	fmt.Print(d)
	err := godotenv.Load("config/dev.env")
	if err != nil {
		log.Fatal("cfg: ", err)
	}
}

func runIPFSNode(ctx context.Context) icore.CoreAPI {
	fmt.Println("\n1.-- Running an embedded IPFS node on a tmp directory-- ")
	ipfs, err := clients.SpawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn a node: %s", err))
	}

	fmt.Println("-- IPFS node is running successfully --")
	return ipfs
}

func upload(ctx context.Context, node icore.CoreAPI) string {
	fmt.Println("\n2.-- Adding test.pdf from the etc dir --")
	cidFile, err := clients.UploadNGetCID(ctx, os.Getenv("UPLOAD_FILE"), node)
	if err != nil {
		panic(err)
	}

	cid := cidFile.String()
	fmt.Printf("Added file to IPFS with CID: %s\n", cid)
	return cid
}

func deploy(scClient *clients.SmartContractClient) *api.Api {
	fmt.Println("\n3.-- Deploying our CID storage smart contract into Ganache --")
	deployedAddress, api := scClient.Deploy()
	fmt.Printf("Deployed address: %v --", deployedAddress.Hex())
	return api
}

func storeCID(cid string, api *api.Api, scClient *clients.SmartContractClient) *types.Transaction {
	fmt.Println("\n\n4.-- Storing our CID into the blockchain --")
	auth := scClient.GetUpdatedOpts()
	tx, err := api.Store(&auth, cid)
	if err != nil {
		log.Fatalf("Error storing: %v", err)
	}

	return tx
}

func retriveCID(api *api.Api) string {
	fmt.Println("\n5.-- Retrieving our CID from the blockchain --")
	cid, err := api.Retrieve(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Error retrieving: %v", err)
	}

	return cid
}
