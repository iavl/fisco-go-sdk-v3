// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package file_box_v2

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// FileBoxV2File is an auto generated low-level Go binding around an user-defined struct.
type FileBoxV2File struct {
	FileHash            [32]byte
	Owner               common.Address
	Deleted             bool
	StorageSpaceType    string
	StorageSpaceAddress common.Address
	Extra               string
	UploadTimestamp     *big.Int
	UpdatedTimestamp    *big.Int
}

// FileBoxV2User is an auto generated low-level Go binding around an user-defined struct.
type FileBoxV2User struct {
	User             common.Address
	Deleted          bool
	UserType         string
	CreatedTimestamp *big.Int
	UpdatedTimestamp *big.Int
}

// FileBoxV2ABI is the input ABI used to generate the binding from.
const FileBoxV2ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FileDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FileRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"shareType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FileShared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"storageSpaceType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"storageSpaceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FileUploaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"UserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"UserDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"userType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"deleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"deleteUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"getFile\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"storageSpaceType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"storageSpaceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"uploadTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structFileBoxV2.File\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"userType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structFileBoxV2.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"recoverFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"shareType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"shareFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"storageSpaceType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"storageSpaceAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"extra\",\"type\":\"string\"}],\"name\":\"uploadFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FileBoxV2Bin is the compiled bytecode used for deploying new contracts.
var FileBoxV2Bin = "0x608060405234801561001057600080fd5b5061262f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063836904831161005b57806383690483146100fd578063d587d68314610119578063e9247f7b14610149578063f68bee571461016557610088565b80631e1f13531461008d5780634f09077f146100a957806364b92f7c146100c55780637a6bb489146100e1575b600080fd5b6100a760048036038101906100a29190611b34565b610195565b005b6100c360048036038101906100be9190611d5b565b6105a1565b005b6100df60048036038101906100da9190611b34565b61089e565b005b6100fb60048036038101906100f69190611e3c565b610d41565b005b61011760048036038101906101129190611eab565b610f33565b005b610133600480360381019061012e9190611eeb565b611103565b6040516101409190612059565b60405180910390f35b610163600480360381019061015e919061207b565b611272565b005b61017f600480360381019061017a9190611b34565b6116b7565b60405161018c91906121f4565b60405180910390f35b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060405180610100016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff1615151515815260200160028201805461027e90612245565b80601f01602080910402602001604051908101604052809291908181526020018280546102aa90612245565b80156102f75780601f106102cc576101008083540402835291602001916102f7565b820191906000526020600020905b8154815290600101906020018083116102da57829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201805461036690612245565b80601f016020809104026020016040519081016040528092919081815260200182805461039290612245565b80156103df5780601f106103b4576101008083540402835291602001916103df565b820191906000526020600020905b8154815290600101906020018083116103c257829003601f168201915b5050505050815260200160058201548152602001600682015481525050905060008160c0015111610445576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161043c906122d4565b60405180910390fd5b80604001511561048a576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161048190612340565b60405180910390fd5b60018060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060010160146101000a81548160ff02191690831515021790555042600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060060181905550818373ffffffffffffffffffffffffffffffffffffffff167f3a16f0dd0cb4562a9ffc943265e5891f14b543bec0b4625a1073248aa12406c242604051610594919061236f565b60405180910390a3505050565b8560008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015411801561063f57506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16155b61067e576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610675906123d6565b60405180910390fd5b6040518061010001604052808781526020018873ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020018581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200142815260200142815250600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008881526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160020190805190602001906107c492919061192c565b5060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600401908051906020019061082892919061192c565b5060c0820151816005015560e08201518160060155905050858773ffffffffffffffffffffffffffffffffffffffff167f077335aa0b46ad39d339be64b159f600f2cc888f5465886abf7eeb2bfb8acb278686864260405161088d949392919061243e565b60405180910390a350505050505050565b8160008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015411801561093c57506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16155b61097b576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610972906123d6565b60405180910390fd5b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060405180610100016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff16151515158152602001600282018054610a6490612245565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9090612245565b8015610add5780601f10610ab257610100808354040283529160200191610add565b820191906000526020600020905b815481529060010190602001808311610ac057829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482018054610b4c90612245565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7890612245565b8015610bc55780601f10610b9a57610100808354040283529160200191610bc5565b820191906000526020600020905b815481529060010190602001808311610ba857829003601f168201915b505050505081526020016005820154815260200160068201548152505090508060400151610c28576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610c1f906124dd565b60405180910390fd5b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060010160146101000a81548160ff02191690831515021790555042600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060060181905550828473ffffffffffffffffffffffffffffffffffffffff167f6a9c4ad9969ae4489ca9b1d27b2fab987e715f2cdda31088760b751f9f57d50342604051610d33919061236f565b60405180910390a350505050565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015414610dc5576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610dbc90612549565b60405180910390fd5b6040518060a001604052808473ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001838152602001428152602001428152506000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff0219169083151502179055506040820151816001019080519060200190610ec692919061192c565b5060608201518160020155608082015181600301559050508273ffffffffffffffffffffffffffffffffffffffff167fd21dc4fd6530a6f2c798a104bddf57bed720e3a70a6a384a511384f94f792ec48383604051610f26929190612569565b60405180910390a2505050565b8160008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154118015610fd157506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16155b611010576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401611007906123d6565b60405180910390fd5b60016000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548160ff021916908315150217905550426000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301819055508273ffffffffffffffffffffffffffffffffffffffff167f595678fabc53ee285ad70de61dcd0ef440fc67d79b33b254d0ee3953980d2911836040516110f69190612599565b60405180910390a2505050565b61110b6119b2565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff161515151581526020016001820180546111d590612245565b80601f016020809104026020016040519081016040528092919081815260200182805461120190612245565b801561124e5780601f106112235761010080835404028352916020019161124e565b820191906000526020600020905b81548152906001019060200180831161123157829003601f168201915b50505050508152602001600282015481526020016003820154815250509050919050565b8460008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015411801561131057506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16155b61134f576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401611346906123d6565b60405180910390fd5b6000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060405180610100016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff1615151515815260200160028201805461143890612245565b80601f016020809104026020016040519081016040528092919081815260200182805461146490612245565b80156114b15780601f10611486576101008083540402835291602001916114b1565b820191906000526020600020905b81548152906001019060200180831161149457829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201805461152090612245565b80601f016020809104026020016040519081016040528092919081815260200182805461154c90612245565b80156115995780601f1061156e57610100808354040283529160200191611599565b820191906000526020600020905b81548152906001019060200180831161157c57829003601f168201915b5050505050815260200160058201548152602001600682015481525050905060008160c00151116115ff576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016115f6906122d4565b60405180910390fd5b806040015115611644576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161163b90612340565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff16868873ffffffffffffffffffffffffffffffffffffffff167f3003bb4135091ad587e61adc7784f6675f5fa1065e1c0bbe1411c05a277b96c68787426040516116a6939291906125b4565b60405180910390a450505050505050565b6116bf6119f9565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060405180610100016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff161515151581526020016002820180546117a690612245565b80601f01602080910402602001604051908101604052809291908181526020018280546117d290612245565b801561181f5780601f106117f45761010080835404028352916020019161181f565b820191906000526020600020905b81548152906001019060200180831161180257829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201805461188e90612245565b80601f01602080910402602001604051908101604052809291908181526020018280546118ba90612245565b80156119075780601f106118dc57610100808354040283529160200191611907565b820191906000526020600020905b8154815290600101906020018083116118ea57829003601f168201915b5050505050815260200160058201548152602001600682015481525050905092915050565b82805461193890612245565b90600052602060002090601f01602090048101928261195a57600085556119a1565b82601f1061197357805160ff19168380011785556119a1565b828001600101855582156119a1579182015b828111156119a0578251825591602001919060010190611985565b5b5090506119ae9190611a6f565b5090565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020016060815260200160008152602001600081525090565b60405180610100016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160001515815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160008152602001600081525090565b5b80821115611a88576000816000905550600101611a70565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611acb82611aa0565b9050919050565b611adb81611ac0565b8114611ae657600080fd5b50565b600081359050611af881611ad2565b92915050565b6000819050919050565b611b1181611afe565b8114611b1c57600080fd5b50565b600081359050611b2e81611b08565b92915050565b60008060408385031215611b4b57611b4a611a96565b5b6000611b5985828601611ae9565b9250506020611b6a85828601611b1f565b9150509250929050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7fb95aa35500000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611bc782611b7e565b810181811067ffffffffffffffff82111715611be657611be5611b8f565b5b80604052505050565b6000611bf9611a8c565b9050611c058282611bbe565b919050565b600067ffffffffffffffff821115611c2557611c24611b8f565b5b611c2e82611b7e565b9050602081019050919050565b82818337600083830152505050565b6000611c5d611c5884611c0a565b611bef565b905082815260208101848484011115611c7957611c78611b79565b5b611c84848285611c3b565b509392505050565b600082601f830112611ca157611ca0611b74565b5b8135611cb1848260208601611c4a565b91505092915050565b600067ffffffffffffffff821115611cd557611cd4611b8f565b5b611cde82611b7e565b9050602081019050919050565b6000611cfe611cf984611cba565b611bef565b905082815260208101848484011115611d1a57611d19611b79565b5b611d25848285611c3b565b509392505050565b600082601f830112611d4257611d41611b74565b5b8135611d52848260208601611ceb565b91505092915050565b60008060008060008060c08789031215611d7857611d77611a96565b5b6000611d8689828a01611ae9565b9650506020611d9789828a01611b1f565b955050604087013567ffffffffffffffff811115611db857611db7611a9b565b5b611dc489828a01611c8c565b945050606087013567ffffffffffffffff811115611de557611de4611a9b565b5b611df189828a01611d2d565b9350506080611e0289828a01611ae9565b92505060a087013567ffffffffffffffff811115611e2357611e22611a9b565b5b611e2f89828a01611d2d565b9150509295509295509295565b600080600060608486031215611e5557611e54611a96565b5b6000611e6386828701611ae9565b935050602084013567ffffffffffffffff811115611e8457611e83611a9b565b5b611e9086828701611d2d565b9250506040611ea186828701611ae9565b9150509250925092565b60008060408385031215611ec257611ec1611a96565b5b6000611ed085828601611ae9565b9250506020611ee185828601611ae9565b9150509250929050565b600060208284031215611f0157611f00611a96565b5b6000611f0f84828501611ae9565b91505092915050565b611f2181611ac0565b82525050565b60008115159050919050565b611f3c81611f27565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611f7c578082015181840152602081019050611f61565b83811115611f8b576000848401525b50505050565b6000611f9c82611f42565b611fa68185611f4d565b9350611fb6818560208601611f5e565b611fbf81611b7e565b840191505092915050565b6000819050919050565b611fdd81611fca565b82525050565b600060a083016000830151611ffb6000860182611f18565b50602083015161200e6020860182611f33565b50604083015184820360408601526120268282611f91565b915050606083015161203b6060860182611fd4565b50608083015161204e6080860182611fd4565b508091505092915050565b600060208201905081810360008301526120738184611fe3565b905092915050565b600080600080600060a0868803121561209757612096611a96565b5b60006120a588828901611ae9565b95505060206120b688828901611b1f565b94505060406120c788828901611ae9565b935050606086013567ffffffffffffffff8111156120e8576120e7611a9b565b5b6120f488828901611d2d565b925050608086013567ffffffffffffffff81111561211557612114611a9b565b5b61212188828901611d2d565b9150509295509295909350565b61213781611afe565b82525050565b600061010083016000830151612156600086018261212e565b5060208301516121696020860182611f18565b50604083015161217c6040860182611f33565b50606083015184820360608601526121948282611f91565b91505060808301516121a96080860182611f18565b5060a083015184820360a08601526121c18282611f91565b91505060c08301516121d660c0860182611fd4565b5060e08301516121e960e0860182611fd4565b508091505092915050565b6000602082019050818103600083015261220e818461213d565b905092915050565b7fb95aa35500000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061225d57607f821691505b6020821081141561227157612270612216565b5b50919050565b600082825260208201905092915050565b7f46696c65206e6f7420666f756e64000000000000000000000000000000000000600082015250565b60006122be600e83612277565b91506122c982612288565b602082019050919050565b600060208201905081810360008301526122ed816122b1565b9050919050565b7f46696c652069732064656c657465640000000000000000000000000000000000600082015250565b600061232a600f83612277565b9150612335826122f4565b602082019050919050565b600060208201905081810360008301526123598161231d565b9050919050565b61236981611fca565b82525050565b60006020820190506123846000830184612360565b92915050565b7f55736572206e6f7420666f756e64000000000000000000000000000000000000600082015250565b60006123c0600e83612277565b91506123cb8261238a565b602082019050919050565b600060208201905081810360008301526123ef816123b3565b9050919050565b600061240182611f42565b61240b8185612277565b935061241b818560208601611f5e565b61242481611b7e565b840191505092915050565b61243881611ac0565b82525050565b6000608082019050818103600083015261245881876123f6565b9050612467602083018661242f565b818103604083015261247981856123f6565b90506124886060830184612360565b95945050505050565b7f46696c65206973206e6f742064656c6574656400000000000000000000000000600082015250565b60006124c7601383612277565b91506124d282612491565b602082019050919050565b600060208201905081810360008301526124f6816124ba565b9050919050565b7f5573657220616c72656164792065786973747300000000000000000000000000600082015250565b6000612533601383612277565b915061253e826124fd565b602082019050919050565b6000602082019050818103600083015261256281612526565b9050919050565b6000604082019050818103600083015261258381856123f6565b9050612592602083018461242f565b9392505050565b60006020820190506125ae600083018461242f565b92915050565b600060608201905081810360008301526125ce81866123f6565b905081810360208301526125e281856123f6565b90506125f16040830184612360565b94935050505056fea264697066735822122054ff805fc1f0fee6e77849f0ff00f7975a522f7d04474e0a15089c6f33712e5864736f6c634300080b0033"
var FileBoxV2SMBin = "0x"

// DeployFileBoxV2 deploys a new contract, binding an instance of FileBoxV2 to it.
func DeployFileBoxV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *FileBoxV2, error) {
	parsed, err := abi.JSON(strings.NewReader(FileBoxV2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(FileBoxV2SMBin)
	} else {
		bytecode = common.FromHex(FileBoxV2Bin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, FileBoxV2ABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &FileBoxV2{FileBoxV2Caller: FileBoxV2Caller{contract: contract}, FileBoxV2Transactor: FileBoxV2Transactor{contract: contract}, FileBoxV2Filterer: FileBoxV2Filterer{contract: contract}}, nil
}

func AsyncDeployFileBoxV2(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(FileBoxV2ABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(FileBoxV2SMBin)
	} else {
		bytecode = common.FromHex(FileBoxV2Bin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, FileBoxV2ABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// FileBoxV2 is an auto generated Go binding around a Solidity contract.
type FileBoxV2 struct {
	FileBoxV2Caller     // Read-only binding to the contract
	FileBoxV2Transactor // Write-only binding to the contract
	FileBoxV2Filterer   // Log filterer for contract events
}

// FileBoxV2Caller is an auto generated read-only Go binding around a Solidity contract.
type FileBoxV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileBoxV2Transactor is an auto generated write-only Go binding around a Solidity contract.
type FileBoxV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileBoxV2Filterer is an auto generated log filtering Go binding around a Solidity contract events.
type FileBoxV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileBoxV2Session is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type FileBoxV2Session struct {
	Contract     *FileBoxV2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileBoxV2CallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type FileBoxV2CallerSession struct {
	Contract *FileBoxV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FileBoxV2TransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type FileBoxV2TransactorSession struct {
	Contract     *FileBoxV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FileBoxV2Raw is an auto generated low-level Go binding around a Solidity contract.
type FileBoxV2Raw struct {
	Contract *FileBoxV2 // Generic contract binding to access the raw methods on
}

// FileBoxV2CallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type FileBoxV2CallerRaw struct {
	Contract *FileBoxV2Caller // Generic read-only contract binding to access the raw methods on
}

// FileBoxV2TransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type FileBoxV2TransactorRaw struct {
	Contract *FileBoxV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFileBoxV2 creates a new instance of FileBoxV2, bound to a specific deployed contract.
func NewFileBoxV2(address common.Address, backend bind.ContractBackend) (*FileBoxV2, error) {
	contract, err := bindFileBoxV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileBoxV2{FileBoxV2Caller: FileBoxV2Caller{contract: contract}, FileBoxV2Transactor: FileBoxV2Transactor{contract: contract}, FileBoxV2Filterer: FileBoxV2Filterer{contract: contract}}, nil
}

// NewFileBoxV2Caller creates a new read-only instance of FileBoxV2, bound to a specific deployed contract.
func NewFileBoxV2Caller(address common.Address, caller bind.ContractCaller) (*FileBoxV2Caller, error) {
	contract, err := bindFileBoxV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileBoxV2Caller{contract: contract}, nil
}

// NewFileBoxV2Transactor creates a new write-only instance of FileBoxV2, bound to a specific deployed contract.
func NewFileBoxV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FileBoxV2Transactor, error) {
	contract, err := bindFileBoxV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileBoxV2Transactor{contract: contract}, nil
}

// NewFileBoxV2Filterer creates a new log filterer instance of FileBoxV2, bound to a specific deployed contract.
func NewFileBoxV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FileBoxV2Filterer, error) {
	contract, err := bindFileBoxV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileBoxV2Filterer{contract: contract}, nil
}

// bindFileBoxV2 binds a generic wrapper to an already deployed contract.
func bindFileBoxV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileBoxV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileBoxV2 *FileBoxV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileBoxV2.Contract.FileBoxV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileBoxV2 *FileBoxV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.FileBoxV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileBoxV2 *FileBoxV2Raw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.FileBoxV2Transactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileBoxV2 *FileBoxV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileBoxV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileBoxV2 *FileBoxV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileBoxV2 *FileBoxV2TransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetFile is a free data retrieval call binding the contract method 0xf68bee57.
//
// Solidity: function getFile(address user, bytes32 fileHash) constant returns(FileBoxV2File)
func (_FileBoxV2 *FileBoxV2Caller) GetFile(opts *bind.CallOpts, user common.Address, fileHash [32]byte) (FileBoxV2File, error) {
	var (
		ret0 = new(FileBoxV2File)
	)
	out := ret0
	err := _FileBoxV2.contract.Call(opts, out, "getFile", user, fileHash)
	return *ret0, err
}

// GetFile is a free data retrieval call binding the contract method 0xf68bee57.
//
// Solidity: function getFile(address user, bytes32 fileHash) constant returns(FileBoxV2File)
func (_FileBoxV2 *FileBoxV2Session) GetFile(user common.Address, fileHash [32]byte) (FileBoxV2File, error) {
	return _FileBoxV2.Contract.GetFile(&_FileBoxV2.CallOpts, user, fileHash)
}

// GetFile is a free data retrieval call binding the contract method 0xf68bee57.
//
// Solidity: function getFile(address user, bytes32 fileHash) constant returns(FileBoxV2File)
func (_FileBoxV2 *FileBoxV2CallerSession) GetFile(user common.Address, fileHash [32]byte) (FileBoxV2File, error) {
	return _FileBoxV2.Contract.GetFile(&_FileBoxV2.CallOpts, user, fileHash)
}

// GetUser is a free data retrieval call binding the contract method 0xd587d683.
//
// Solidity: function getUser(address user_) constant returns(FileBoxV2User)
func (_FileBoxV2 *FileBoxV2Caller) GetUser(opts *bind.CallOpts, user_ common.Address) (FileBoxV2User, error) {
	var (
		ret0 = new(FileBoxV2User)
	)
	out := ret0
	err := _FileBoxV2.contract.Call(opts, out, "getUser", user_)
	return *ret0, err
}

// GetUser is a free data retrieval call binding the contract method 0xd587d683.
//
// Solidity: function getUser(address user_) constant returns(FileBoxV2User)
func (_FileBoxV2 *FileBoxV2Session) GetUser(user_ common.Address) (FileBoxV2User, error) {
	return _FileBoxV2.Contract.GetUser(&_FileBoxV2.CallOpts, user_)
}

// GetUser is a free data retrieval call binding the contract method 0xd587d683.
//
// Solidity: function getUser(address user_) constant returns(FileBoxV2User)
func (_FileBoxV2 *FileBoxV2CallerSession) GetUser(user_ common.Address) (FileBoxV2User, error) {
	return _FileBoxV2.Contract.GetUser(&_FileBoxV2.CallOpts, user_)
}

// AddUser is a paid mutator transaction binding the contract method 0x7a6bb489.
//
// Solidity: function addUser(address user, string userType, address admin) returns()
func (_FileBoxV2 *FileBoxV2Transactor) AddUser(opts *bind.TransactOpts, user common.Address, userType string, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileBoxV2.contract.TransactWithResult(opts, out, "addUser", user, userType, admin)
	return transaction, receipt, err
}

func (_FileBoxV2 *FileBoxV2Transactor) AsyncAddUser(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address, userType string, admin common.Address) (*types.Transaction, error) {
	return _FileBoxV2.contract.AsyncTransact(opts, handler, "addUser", user, userType, admin)
}

// AddUser is a paid mutator transaction binding the contract method 0x7a6bb489.
//
// Solidity: function addUser(address user, string userType, address admin) returns()
func (_FileBoxV2 *FileBoxV2Session) AddUser(user common.Address, userType string, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.AddUser(&_FileBoxV2.TransactOpts, user, userType, admin)
}

func (_FileBoxV2 *FileBoxV2Session) AsyncAddUser(handler func(*types.Receipt, error), user common.Address, userType string, admin common.Address) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncAddUser(handler, &_FileBoxV2.TransactOpts, user, userType, admin)
}

// AddUser is a paid mutator transaction binding the contract method 0x7a6bb489.
//
// Solidity: function addUser(address user, string userType, address admin) returns()
func (_FileBoxV2 *FileBoxV2TransactorSession) AddUser(user common.Address, userType string, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.AddUser(&_FileBoxV2.TransactOpts, user, userType, admin)
}

func (_FileBoxV2 *FileBoxV2TransactorSession) AsyncAddUser(handler func(*types.Receipt, error), user common.Address, userType string, admin common.Address) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncAddUser(handler, &_FileBoxV2.TransactOpts, user, userType, admin)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x1e1f1353.
//
// Solidity: function deleteFile(address user, bytes32 fileHash) returns()
func (_FileBoxV2 *FileBoxV2Transactor) DeleteFile(opts *bind.TransactOpts, user common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileBoxV2.contract.TransactWithResult(opts, out, "deleteFile", user, fileHash)
	return transaction, receipt, err
}

func (_FileBoxV2 *FileBoxV2Transactor) AsyncDeleteFile(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileBoxV2.contract.AsyncTransact(opts, handler, "deleteFile", user, fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x1e1f1353.
//
// Solidity: function deleteFile(address user, bytes32 fileHash) returns()
func (_FileBoxV2 *FileBoxV2Session) DeleteFile(user common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.DeleteFile(&_FileBoxV2.TransactOpts, user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Session) AsyncDeleteFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncDeleteFile(handler, &_FileBoxV2.TransactOpts, user, fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x1e1f1353.
//
// Solidity: function deleteFile(address user, bytes32 fileHash) returns()
func (_FileBoxV2 *FileBoxV2TransactorSession) DeleteFile(user common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.DeleteFile(&_FileBoxV2.TransactOpts, user, fileHash)
}

func (_FileBoxV2 *FileBoxV2TransactorSession) AsyncDeleteFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncDeleteFile(handler, &_FileBoxV2.TransactOpts, user, fileHash)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x83690483.
//
// Solidity: function deleteUser(address user, address admin) returns()
func (_FileBoxV2 *FileBoxV2Transactor) DeleteUser(opts *bind.TransactOpts, user common.Address, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileBoxV2.contract.TransactWithResult(opts, out, "deleteUser", user, admin)
	return transaction, receipt, err
}

func (_FileBoxV2 *FileBoxV2Transactor) AsyncDeleteUser(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address, admin common.Address) (*types.Transaction, error) {
	return _FileBoxV2.contract.AsyncTransact(opts, handler, "deleteUser", user, admin)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x83690483.
//
// Solidity: function deleteUser(address user, address admin) returns()
func (_FileBoxV2 *FileBoxV2Session) DeleteUser(user common.Address, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.DeleteUser(&_FileBoxV2.TransactOpts, user, admin)
}

func (_FileBoxV2 *FileBoxV2Session) AsyncDeleteUser(handler func(*types.Receipt, error), user common.Address, admin common.Address) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncDeleteUser(handler, &_FileBoxV2.TransactOpts, user, admin)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x83690483.
//
// Solidity: function deleteUser(address user, address admin) returns()
func (_FileBoxV2 *FileBoxV2TransactorSession) DeleteUser(user common.Address, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.DeleteUser(&_FileBoxV2.TransactOpts, user, admin)
}

func (_FileBoxV2 *FileBoxV2TransactorSession) AsyncDeleteUser(handler func(*types.Receipt, error), user common.Address, admin common.Address) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncDeleteUser(handler, &_FileBoxV2.TransactOpts, user, admin)
}

// RecoverFile is a paid mutator transaction binding the contract method 0x64b92f7c.
//
// Solidity: function recoverFile(address user, bytes32 fileHash) returns()
func (_FileBoxV2 *FileBoxV2Transactor) RecoverFile(opts *bind.TransactOpts, user common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileBoxV2.contract.TransactWithResult(opts, out, "recoverFile", user, fileHash)
	return transaction, receipt, err
}

func (_FileBoxV2 *FileBoxV2Transactor) AsyncRecoverFile(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileBoxV2.contract.AsyncTransact(opts, handler, "recoverFile", user, fileHash)
}

// RecoverFile is a paid mutator transaction binding the contract method 0x64b92f7c.
//
// Solidity: function recoverFile(address user, bytes32 fileHash) returns()
func (_FileBoxV2 *FileBoxV2Session) RecoverFile(user common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.RecoverFile(&_FileBoxV2.TransactOpts, user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Session) AsyncRecoverFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncRecoverFile(handler, &_FileBoxV2.TransactOpts, user, fileHash)
}

// RecoverFile is a paid mutator transaction binding the contract method 0x64b92f7c.
//
// Solidity: function recoverFile(address user, bytes32 fileHash) returns()
func (_FileBoxV2 *FileBoxV2TransactorSession) RecoverFile(user common.Address, fileHash [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.RecoverFile(&_FileBoxV2.TransactOpts, user, fileHash)
}

func (_FileBoxV2 *FileBoxV2TransactorSession) AsyncRecoverFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncRecoverFile(handler, &_FileBoxV2.TransactOpts, user, fileHash)
}

// ShareFile is a paid mutator transaction binding the contract method 0xe9247f7b.
//
// Solidity: function shareFile(address user, bytes32 fileHash, address to, string shareType, string extra) returns()
func (_FileBoxV2 *FileBoxV2Transactor) ShareFile(opts *bind.TransactOpts, user common.Address, fileHash [32]byte, to common.Address, shareType string, extra string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileBoxV2.contract.TransactWithResult(opts, out, "shareFile", user, fileHash, to, shareType, extra)
	return transaction, receipt, err
}

func (_FileBoxV2 *FileBoxV2Transactor) AsyncShareFile(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address, fileHash [32]byte, to common.Address, shareType string, extra string) (*types.Transaction, error) {
	return _FileBoxV2.contract.AsyncTransact(opts, handler, "shareFile", user, fileHash, to, shareType, extra)
}

// ShareFile is a paid mutator transaction binding the contract method 0xe9247f7b.
//
// Solidity: function shareFile(address user, bytes32 fileHash, address to, string shareType, string extra) returns()
func (_FileBoxV2 *FileBoxV2Session) ShareFile(user common.Address, fileHash [32]byte, to common.Address, shareType string, extra string) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.ShareFile(&_FileBoxV2.TransactOpts, user, fileHash, to, shareType, extra)
}

func (_FileBoxV2 *FileBoxV2Session) AsyncShareFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte, to common.Address, shareType string, extra string) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncShareFile(handler, &_FileBoxV2.TransactOpts, user, fileHash, to, shareType, extra)
}

// ShareFile is a paid mutator transaction binding the contract method 0xe9247f7b.
//
// Solidity: function shareFile(address user, bytes32 fileHash, address to, string shareType, string extra) returns()
func (_FileBoxV2 *FileBoxV2TransactorSession) ShareFile(user common.Address, fileHash [32]byte, to common.Address, shareType string, extra string) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.ShareFile(&_FileBoxV2.TransactOpts, user, fileHash, to, shareType, extra)
}

func (_FileBoxV2 *FileBoxV2TransactorSession) AsyncShareFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte, to common.Address, shareType string, extra string) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncShareFile(handler, &_FileBoxV2.TransactOpts, user, fileHash, to, shareType, extra)
}

// UploadFile is a paid mutator transaction binding the contract method 0x4f09077f.
//
// Solidity: function uploadFile(address user, bytes32 fileHash, bytes signature, string storageSpaceType, address storageSpaceAddress, string extra) returns()
func (_FileBoxV2 *FileBoxV2Transactor) UploadFile(opts *bind.TransactOpts, user common.Address, fileHash [32]byte, signature []byte, storageSpaceType string, storageSpaceAddress common.Address, extra string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _FileBoxV2.contract.TransactWithResult(opts, out, "uploadFile", user, fileHash, signature, storageSpaceType, storageSpaceAddress, extra)
	return transaction, receipt, err
}

func (_FileBoxV2 *FileBoxV2Transactor) AsyncUploadFile(handler func(*types.Receipt, error), opts *bind.TransactOpts, user common.Address, fileHash [32]byte, signature []byte, storageSpaceType string, storageSpaceAddress common.Address, extra string) (*types.Transaction, error) {
	return _FileBoxV2.contract.AsyncTransact(opts, handler, "uploadFile", user, fileHash, signature, storageSpaceType, storageSpaceAddress, extra)
}

// UploadFile is a paid mutator transaction binding the contract method 0x4f09077f.
//
// Solidity: function uploadFile(address user, bytes32 fileHash, bytes signature, string storageSpaceType, address storageSpaceAddress, string extra) returns()
func (_FileBoxV2 *FileBoxV2Session) UploadFile(user common.Address, fileHash [32]byte, signature []byte, storageSpaceType string, storageSpaceAddress common.Address, extra string) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.UploadFile(&_FileBoxV2.TransactOpts, user, fileHash, signature, storageSpaceType, storageSpaceAddress, extra)
}

func (_FileBoxV2 *FileBoxV2Session) AsyncUploadFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte, signature []byte, storageSpaceType string, storageSpaceAddress common.Address, extra string) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncUploadFile(handler, &_FileBoxV2.TransactOpts, user, fileHash, signature, storageSpaceType, storageSpaceAddress, extra)
}

// UploadFile is a paid mutator transaction binding the contract method 0x4f09077f.
//
// Solidity: function uploadFile(address user, bytes32 fileHash, bytes signature, string storageSpaceType, address storageSpaceAddress, string extra) returns()
func (_FileBoxV2 *FileBoxV2TransactorSession) UploadFile(user common.Address, fileHash [32]byte, signature []byte, storageSpaceType string, storageSpaceAddress common.Address, extra string) (*types.Transaction, *types.Receipt, error) {
	return _FileBoxV2.Contract.UploadFile(&_FileBoxV2.TransactOpts, user, fileHash, signature, storageSpaceType, storageSpaceAddress, extra)
}

func (_FileBoxV2 *FileBoxV2TransactorSession) AsyncUploadFile(handler func(*types.Receipt, error), user common.Address, fileHash [32]byte, signature []byte, storageSpaceType string, storageSpaceAddress common.Address, extra string) (*types.Transaction, error) {
	return _FileBoxV2.Contract.AsyncUploadFile(handler, &_FileBoxV2.TransactOpts, user, fileHash, signature, storageSpaceType, storageSpaceAddress, extra)
}

// FileBoxV2FileDeleted represents a FileDeleted event raised by the FileBoxV2 contract.
type FileBoxV2FileDeleted struct {
	User      common.Address
	FileHash  [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// WatchFileDeleted is a free log subscription operation binding the contract event 0x3a16f0dd0cb4562a9ffc943265e5891f14b543bec0b4625a1073248aa12406c2.
//
// Solidity: event FileDeleted(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) WatchFileDeleted(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileDeleted", user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Filterer) WatchAllFileDeleted(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileDeleted")
}

// ParseFileDeleted is a log parse operation binding the contract event 0x3a16f0dd0cb4562a9ffc943265e5891f14b543bec0b4625a1073248aa12406c2.
//
// Solidity: event FileDeleted(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) ParseFileDeleted(log types.Log) (*FileBoxV2FileDeleted, error) {
	event := new(FileBoxV2FileDeleted)
	if err := _FileBoxV2.contract.UnpackLog(event, "FileDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchFileDeleted is a free log subscription operation binding the contract event 0x3a16f0dd0cb4562a9ffc943265e5891f14b543bec0b4625a1073248aa12406c2.
//
// Solidity: event FileDeleted(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) WatchFileDeleted(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte) (string, error) {
	return _FileBoxV2.Contract.WatchFileDeleted(fromBlock, handler, user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Session) WatchAllFileDeleted(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.Contract.WatchAllFileDeleted(fromBlock, handler)
}

// ParseFileDeleted is a log parse operation binding the contract event 0x3a16f0dd0cb4562a9ffc943265e5891f14b543bec0b4625a1073248aa12406c2.
//
// Solidity: event FileDeleted(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) ParseFileDeleted(log types.Log) (*FileBoxV2FileDeleted, error) {
	return _FileBoxV2.Contract.ParseFileDeleted(log)
}

// FileBoxV2FileRecovered represents a FileRecovered event raised by the FileBoxV2 contract.
type FileBoxV2FileRecovered struct {
	User      common.Address
	FileHash  [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// WatchFileRecovered is a free log subscription operation binding the contract event 0x6a9c4ad9969ae4489ca9b1d27b2fab987e715f2cdda31088760b751f9f57d503.
//
// Solidity: event FileRecovered(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) WatchFileRecovered(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileRecovered", user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Filterer) WatchAllFileRecovered(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileRecovered")
}

// ParseFileRecovered is a log parse operation binding the contract event 0x6a9c4ad9969ae4489ca9b1d27b2fab987e715f2cdda31088760b751f9f57d503.
//
// Solidity: event FileRecovered(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) ParseFileRecovered(log types.Log) (*FileBoxV2FileRecovered, error) {
	event := new(FileBoxV2FileRecovered)
	if err := _FileBoxV2.contract.UnpackLog(event, "FileRecovered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchFileRecovered is a free log subscription operation binding the contract event 0x6a9c4ad9969ae4489ca9b1d27b2fab987e715f2cdda31088760b751f9f57d503.
//
// Solidity: event FileRecovered(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) WatchFileRecovered(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte) (string, error) {
	return _FileBoxV2.Contract.WatchFileRecovered(fromBlock, handler, user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Session) WatchAllFileRecovered(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.Contract.WatchAllFileRecovered(fromBlock, handler)
}

// ParseFileRecovered is a log parse operation binding the contract event 0x6a9c4ad9969ae4489ca9b1d27b2fab987e715f2cdda31088760b751f9f57d503.
//
// Solidity: event FileRecovered(address indexed user, bytes32 indexed fileHash, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) ParseFileRecovered(log types.Log) (*FileBoxV2FileRecovered, error) {
	return _FileBoxV2.Contract.ParseFileRecovered(log)
}

// FileBoxV2FileShared represents a FileShared event raised by the FileBoxV2 contract.
type FileBoxV2FileShared struct {
	User      common.Address
	FileHash  [32]byte
	To        common.Address
	ShareType string
	Extra     string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// WatchFileShared is a free log subscription operation binding the contract event 0x3003bb4135091ad587e61adc7784f6675f5fa1065e1c0bbe1411c05a277b96c6.
//
// Solidity: event FileShared(address indexed user, bytes32 indexed fileHash, address indexed to, string shareType, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) WatchFileShared(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte, to common.Address) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileShared", user, fileHash, to)
}

func (_FileBoxV2 *FileBoxV2Filterer) WatchAllFileShared(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileShared")
}

// ParseFileShared is a log parse operation binding the contract event 0x3003bb4135091ad587e61adc7784f6675f5fa1065e1c0bbe1411c05a277b96c6.
//
// Solidity: event FileShared(address indexed user, bytes32 indexed fileHash, address indexed to, string shareType, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) ParseFileShared(log types.Log) (*FileBoxV2FileShared, error) {
	event := new(FileBoxV2FileShared)
	if err := _FileBoxV2.contract.UnpackLog(event, "FileShared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchFileShared is a free log subscription operation binding the contract event 0x3003bb4135091ad587e61adc7784f6675f5fa1065e1c0bbe1411c05a277b96c6.
//
// Solidity: event FileShared(address indexed user, bytes32 indexed fileHash, address indexed to, string shareType, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) WatchFileShared(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte, to common.Address) (string, error) {
	return _FileBoxV2.Contract.WatchFileShared(fromBlock, handler, user, fileHash, to)
}

func (_FileBoxV2 *FileBoxV2Session) WatchAllFileShared(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.Contract.WatchAllFileShared(fromBlock, handler)
}

// ParseFileShared is a log parse operation binding the contract event 0x3003bb4135091ad587e61adc7784f6675f5fa1065e1c0bbe1411c05a277b96c6.
//
// Solidity: event FileShared(address indexed user, bytes32 indexed fileHash, address indexed to, string shareType, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) ParseFileShared(log types.Log) (*FileBoxV2FileShared, error) {
	return _FileBoxV2.Contract.ParseFileShared(log)
}

// FileBoxV2FileUploaded represents a FileUploaded event raised by the FileBoxV2 contract.
type FileBoxV2FileUploaded struct {
	User                common.Address
	FileHash            [32]byte
	StorageSpaceType    string
	StorageSpaceAddress common.Address
	Extra               string
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0x077335aa0b46ad39d339be64b159f600f2cc888f5465886abf7eeb2bfb8acb27.
//
// Solidity: event FileUploaded(address indexed user, bytes32 indexed fileHash, string storageSpaceType, address storageSpaceAddress, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) WatchFileUploaded(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileUploaded", user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Filterer) WatchAllFileUploaded(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "FileUploaded")
}

// ParseFileUploaded is a log parse operation binding the contract event 0x077335aa0b46ad39d339be64b159f600f2cc888f5465886abf7eeb2bfb8acb27.
//
// Solidity: event FileUploaded(address indexed user, bytes32 indexed fileHash, string storageSpaceType, address storageSpaceAddress, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Filterer) ParseFileUploaded(log types.Log) (*FileBoxV2FileUploaded, error) {
	event := new(FileBoxV2FileUploaded)
	if err := _FileBoxV2.contract.UnpackLog(event, "FileUploaded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0x077335aa0b46ad39d339be64b159f600f2cc888f5465886abf7eeb2bfb8acb27.
//
// Solidity: event FileUploaded(address indexed user, bytes32 indexed fileHash, string storageSpaceType, address storageSpaceAddress, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) WatchFileUploaded(fromBlock *int64, handler func(int, []types.Log), user common.Address, fileHash [32]byte) (string, error) {
	return _FileBoxV2.Contract.WatchFileUploaded(fromBlock, handler, user, fileHash)
}

func (_FileBoxV2 *FileBoxV2Session) WatchAllFileUploaded(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.Contract.WatchAllFileUploaded(fromBlock, handler)
}

// ParseFileUploaded is a log parse operation binding the contract event 0x077335aa0b46ad39d339be64b159f600f2cc888f5465886abf7eeb2bfb8acb27.
//
// Solidity: event FileUploaded(address indexed user, bytes32 indexed fileHash, string storageSpaceType, address storageSpaceAddress, string extra, uint256 timestamp)
func (_FileBoxV2 *FileBoxV2Session) ParseFileUploaded(log types.Log) (*FileBoxV2FileUploaded, error) {
	return _FileBoxV2.Contract.ParseFileUploaded(log)
}

// FileBoxV2UserAdded represents a UserAdded event raised by the FileBoxV2 contract.
type FileBoxV2UserAdded struct {
	User     common.Address
	UserType string
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchUserAdded is a free log subscription operation binding the contract event 0xd21dc4fd6530a6f2c798a104bddf57bed720e3a70a6a384a511384f94f792ec4.
//
// Solidity: event UserAdded(address indexed user, string userType, address admin)
func (_FileBoxV2 *FileBoxV2Filterer) WatchUserAdded(fromBlock *int64, handler func(int, []types.Log), user common.Address) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "UserAdded", user)
}

func (_FileBoxV2 *FileBoxV2Filterer) WatchAllUserAdded(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "UserAdded")
}

// ParseUserAdded is a log parse operation binding the contract event 0xd21dc4fd6530a6f2c798a104bddf57bed720e3a70a6a384a511384f94f792ec4.
//
// Solidity: event UserAdded(address indexed user, string userType, address admin)
func (_FileBoxV2 *FileBoxV2Filterer) ParseUserAdded(log types.Log) (*FileBoxV2UserAdded, error) {
	event := new(FileBoxV2UserAdded)
	if err := _FileBoxV2.contract.UnpackLog(event, "UserAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchUserAdded is a free log subscription operation binding the contract event 0xd21dc4fd6530a6f2c798a104bddf57bed720e3a70a6a384a511384f94f792ec4.
//
// Solidity: event UserAdded(address indexed user, string userType, address admin)
func (_FileBoxV2 *FileBoxV2Session) WatchUserAdded(fromBlock *int64, handler func(int, []types.Log), user common.Address) (string, error) {
	return _FileBoxV2.Contract.WatchUserAdded(fromBlock, handler, user)
}

func (_FileBoxV2 *FileBoxV2Session) WatchAllUserAdded(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.Contract.WatchAllUserAdded(fromBlock, handler)
}

// ParseUserAdded is a log parse operation binding the contract event 0xd21dc4fd6530a6f2c798a104bddf57bed720e3a70a6a384a511384f94f792ec4.
//
// Solidity: event UserAdded(address indexed user, string userType, address admin)
func (_FileBoxV2 *FileBoxV2Session) ParseUserAdded(log types.Log) (*FileBoxV2UserAdded, error) {
	return _FileBoxV2.Contract.ParseUserAdded(log)
}

// FileBoxV2UserDeleted represents a UserDeleted event raised by the FileBoxV2 contract.
type FileBoxV2UserDeleted struct {
	User  common.Address
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchUserDeleted is a free log subscription operation binding the contract event 0x595678fabc53ee285ad70de61dcd0ef440fc67d79b33b254d0ee3953980d2911.
//
// Solidity: event UserDeleted(address indexed user, address admin)
func (_FileBoxV2 *FileBoxV2Filterer) WatchUserDeleted(fromBlock *int64, handler func(int, []types.Log), user common.Address) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "UserDeleted", user)
}

func (_FileBoxV2 *FileBoxV2Filterer) WatchAllUserDeleted(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.contract.WatchLogs(fromBlock, handler, "UserDeleted")
}

// ParseUserDeleted is a log parse operation binding the contract event 0x595678fabc53ee285ad70de61dcd0ef440fc67d79b33b254d0ee3953980d2911.
//
// Solidity: event UserDeleted(address indexed user, address admin)
func (_FileBoxV2 *FileBoxV2Filterer) ParseUserDeleted(log types.Log) (*FileBoxV2UserDeleted, error) {
	event := new(FileBoxV2UserDeleted)
	if err := _FileBoxV2.contract.UnpackLog(event, "UserDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchUserDeleted is a free log subscription operation binding the contract event 0x595678fabc53ee285ad70de61dcd0ef440fc67d79b33b254d0ee3953980d2911.
//
// Solidity: event UserDeleted(address indexed user, address admin)
func (_FileBoxV2 *FileBoxV2Session) WatchUserDeleted(fromBlock *int64, handler func(int, []types.Log), user common.Address) (string, error) {
	return _FileBoxV2.Contract.WatchUserDeleted(fromBlock, handler, user)
}

func (_FileBoxV2 *FileBoxV2Session) WatchAllUserDeleted(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _FileBoxV2.Contract.WatchAllUserDeleted(fromBlock, handler)
}

// ParseUserDeleted is a log parse operation binding the contract event 0x595678fabc53ee285ad70de61dcd0ef440fc67d79b33b254d0ee3953980d2911.
//
// Solidity: event UserDeleted(address indexed user, address admin)
func (_FileBoxV2 *FileBoxV2Session) ParseUserDeleted(log types.Log) (*FileBoxV2UserDeleted, error) {
	return _FileBoxV2.Contract.ParseUserDeleted(log)
}
