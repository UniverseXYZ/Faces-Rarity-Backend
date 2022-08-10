package dlt

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct {
	Client *ethclient.Client
}

func NewEthereumClient(nodeURL string) (*EthereumClient, error) {
	client, err := ethclient.Dial(nodeURL)

	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		Client: client,
	}, nil
}
