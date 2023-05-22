package ethereum

import (
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seedchamp45/satang-go/config"
	"gopkg.in/yaml.v3"
)

// EthereumClient represents the Ethereum client connection
type EthereumClient struct {
	*ethclient.Client
}

// NewEthereumClient creates a new EthereumClient instance and connects to the Ethereum network
func NewEthereumClient() (*EthereumClient, error) {
	// Read the YAML file
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the YAML content
	var ethClient config.ETHClient
	err = yaml.Unmarshal(yamlFile, &ethClient)
	if err != nil {
		log.Fatal(err)
	}

	clientConfig := ethClient.Client
	log.Println(clientConfig)

	client, err := ethclient.Dial(clientConfig)
	if err != nil {
		return nil, err
	}

	return &EthereumClient{Client: client}, nil
}
