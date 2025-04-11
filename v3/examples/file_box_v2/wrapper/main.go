package main

import (
	"context"
	"encoding/hex"
	"log"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	filebox "github.com/FISCO-BCOS/go-sdk/v3/examples/file_box_v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)


func main() {
	privateKey, _ := hex.DecodeString("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	fileBoxContractAddress := common.HexToAddress("0xf3bac27b6e968c51a83f999b3a72a3df9bc9d9f6")
	gateway := common.HexToAddress("0x3075db338861e9bb340ed00e1a5b3a5b6fb4b77e")
	alice := common.HexToAddress("0x089103734d0bf26591a953c07d36e18fe4b67d89")
	bob := common.HexToAddress("0x7e56f853015b610f0901a6cc2b6749197b3a7f5a")

	// connect to the node rpc
	config := &client.Config{
		IsSMCrypto: true,
		GroupID: "group0", 
		DisableSsl: false,
		PrivateKey: privateKey, 
		Host: "127.0.0.1", 
		Port: 20200, 
		TLSCaFile: "./conf/sm_ca.crt", 
		TLSKeyFile: "./conf/sm_sdk.key", 
		TLSCertFile:"./conf/sm_sdk.crt", 
		TLSSmEnKeyFile: "./conf/sm_ensdk.key", 
		TLSSmEnCertFile: "./conf/sm_ensdk.crt",
	}
	client, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal("DialContext failed", err)
	}

	// load the contract
	instance, err := filebox.NewFileBoxV2(fileBoxContractAddress, client)
	if err != nil {
	    log.Fatal("NewFileBoxV2 failed", err)
	}
	fileBoxSession := &filebox.FileBoxV2Session{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	// 1. add user
	userType := "admin" // super admin, admin, member, guest
	_, receipt, err := fileBoxSession.AddUser(alice, userType, gateway)
	if err != nil {
		log.Fatal("AddUser failed", err)
	}
	log.Printf("addUser transaction hash: %s\n", receipt.GetTransactionHash())

	// 2. get user
	user, err := fileBoxSession.GetUser(alice)
	if err != nil {
		log.Fatal("GetUser failed", err)
	}
	log.Printf("getUser: %+v\n", user)

	// 3. upload file for alice
	fileHash := crypto.Keccak256Hash([]byte("0xe71a80263ac5c6918bac2366c6889453653c9f7cf24c91259950fc83c0d24e31"))
	var userSignature []byte
	copy(userSignature[:], "this is the user signature signed by user's private key")
	storageSpaceType := "organization" // organization, group, person
	storageSpaceAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	extra := "extra_info"
	_, receipt, err = fileBoxSession.UploadFile(alice, fileHash, userSignature, storageSpaceType, storageSpaceAddress, extra)
	if err != nil {
		log.Fatal("UploadFile failed", err)
	}
	log.Printf("uploadFile transaction hash: %s\n", receipt.GetTransactionHash())

	// 4. get uploaded file
	file, err := fileBoxSession.GetFile(alice, fileHash)
	if err != nil {
		log.Fatal("GetFile failed", err)
	}
	log.Printf("getFile: %+v\n", file)

	// 5. alice shares file to bob
	shareType := "owner" // shareType
	extra = "extra_info"
	_, receipt, err = fileBoxSession.ShareFile(alice, fileHash, bob, shareType, extra)
	if err != nil {
		log.Fatal("ShareFile failed", err)
	}
	log.Printf("shareFile transaction hash: %s\n", receipt.GetTransactionHash())


	// 6. delete file
	_, receipt, err = fileBoxSession.DeleteFile(alice, fileHash)
	if err != nil {
		log.Fatal("DeleteFile failed", err)
	}
	log.Printf("deleteFile transaction hash: %s\n", receipt.GetTransactionHash())
	// get file after delete
	file, err = fileBoxSession.GetFile(alice, fileHash)
	if err != nil {
		log.Fatal("GetFile failed", err)
	}
	log.Printf("file deleted: %+v\n", file.Deleted)

	// 7. recover file
	_, receipt, err = fileBoxSession.RecoverFile(alice, fileHash)
	if err != nil {
		log.Fatal("RecoverFile failed", err)
	}
	log.Printf("recoverFile transaction hash: %s\n", receipt.GetTransactionHash())
	// get file after recover
	file, err = fileBoxSession.GetFile(alice, fileHash)
	if err != nil {
		log.Fatal("GetFile failed", err)
	}
	log.Printf("file recovered: %+v\n", !file.Deleted)

	// 8. delete user
	_, receipt, err = fileBoxSession.DeleteUser(alice, gateway)
	if err != nil {
		log.Fatal("DeleteUser failed", err)
	}
	log.Printf("deleteUser transaction hash: %s\n", receipt.GetTransactionHash())
	// 9. get user after delete
	user, err = fileBoxSession.GetUser(alice)
	if err != nil {
		log.Fatal("GetUser failed", err)
	}
	log.Printf("user deleted: %+v\n", user.Deleted)
}
