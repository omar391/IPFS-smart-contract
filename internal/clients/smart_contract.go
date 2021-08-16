package clients

import (
	"context"
	"crypto/ecdsa"
	"ipfs_smart_contract/api"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SmartContractClient struct {
	Addr       string
	Privatekey string
	client     *ethclient.Client
	address    *common.Address
	rootAuth   *bind.TransactOpts
}

func (sm *SmartContractClient) lazyInit() {
	if sm.client == nil {
		client, err := ethclient.Dial(sm.Addr)
		sm.client = client
		if err != nil {
			panic(err)
		}

		privateKey, err := crypto.HexToECDSA(sm.Privatekey)
		if err != nil {
			panic(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			panic("invalid key")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		sm.address = &fromAddress
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			panic(err)
		}

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			panic(err)
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			panic(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)      // in wei
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = big.NewInt(1000000)
		sm.rootAuth = auth
	}
}

func (sm *SmartContractClient) GetUpdatedOpts() bind.TransactOpts {
	sm.lazyInit()
	nonce, err := sm.client.PendingNonceAt(context.Background(), *sm.address)
	if err != nil {
		panic(err)
	}

	auth := *sm.rootAuth
	auth.Nonce = big.NewInt(int64(nonce))

	return auth
}

func (sm *SmartContractClient) Deploy() (common.Address, *api.Api) {
	sm.lazyInit()
	address, _, instance, err := api.DeployApi(sm.rootAuth, sm.client)
	if err != nil {
		panic(err)
	}

	return address, instance
}
