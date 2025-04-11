package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
)

func main() {
	privateKey, _ := hex.DecodeString("75941985862045597874049342503894607692969701661690202790539931515978786679111")
	// disable ssl of node rpc
	config := &client.Config{IsSMCrypto: true, GroupID: "group0", DisableSsl: false,
		PrivateKey: privateKey, Host: "157.173.207.120/", Port: 20200, TLSCaFile: "./conf/sm_ca.crt", TLSKeyFile: "./conf/sm_sdk.key", TLSCertFile: "./conf/sm_sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=================DeployFileBox===============")
	address, receipt, _, err := DeployFileBox(client.GetTransactOpts(), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
	fmt.Println("transaction hash: ", receipt.TransactionHash)

}
