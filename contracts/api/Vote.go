// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_options\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"voterCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptions\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"voterCode\",\"type\":\"string\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"}],\"name\":\"getVoteCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"voterCode\",\"type\":\"string\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"}],\"name\":\"isValidOption\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001d2b38038062001d2b8339818101604052810190620000379190620004bf565b620000576200004b6200007760201b60201c565b6200007f60201b60201c565b80600390805190602001906200006f92919062000143565b505062000842565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b82805482825590600052602060002090810192821562000190579160200282015b828111156200018f5782518290816200017e91906200075b565b509160200191906001019062000164565b5b5090506200019f9190620001a3565b5090565b5b80821115620001c75760008181620001bd9190620001cb565b50600101620001a4565b5090565b508054620001d9906200054a565b6000825580601f10620001ed57506200020e565b601f0160209004906000526020600020908101906200020d919062000211565b5b50565b5b808211156200022c57600081600090555060010162000212565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620002948262000249565b810181811067ffffffffffffffff82111715620002b657620002b56200025a565b5b80604052505050565b6000620002cb62000230565b9050620002d9828262000289565b919050565b600067ffffffffffffffff821115620002fc57620002fb6200025a565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff8211156200033557620003346200025a565b5b620003408262000249565b9050602081019050919050565b60005b838110156200036d57808201518184015260208101905062000350565b60008484015250505050565b6000620003906200038a8462000317565b620002bf565b905082815260208101848484011115620003af57620003ae62000312565b5b620003bc8482856200034d565b509392505050565b600082601f830112620003dc57620003db62000244565b5b8151620003ee84826020860162000379565b91505092915050565b60006200040e6200040884620002de565b620002bf565b905080838252602082019050602084028301858111156200043457620004336200030d565b5b835b818110156200048257805167ffffffffffffffff8111156200045d576200045c62000244565b5b8086016200046c8982620003c4565b8552602085019450505060208101905062000436565b5050509392505050565b600082601f830112620004a457620004a362000244565b5b8151620004b6848260208601620003f7565b91505092915050565b600060208284031215620004d857620004d76200023a565b5b600082015167ffffffffffffffff811115620004f957620004f86200023f565b5b62000507848285016200048c565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200056357607f821691505b6020821081036200057957620005786200051b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005e37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005a4565b620005ef8683620005a4565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200063c62000636620006308462000607565b62000611565b62000607565b9050919050565b6000819050919050565b62000658836200061b565b62000670620006678262000643565b848454620005b1565b825550505050565b600090565b6200068762000678565b620006948184846200064d565b505050565b5b81811015620006bc57620006b06000826200067d565b6001810190506200069a565b5050565b601f8211156200070b57620006d5816200057f565b620006e08462000594565b81016020851015620006f0578190505b62000708620006ff8562000594565b83018262000699565b50505b505050565b600082821c905092915050565b6000620007306000198460080262000710565b1980831691505092915050565b60006200074b83836200071d565b9150826002028217905092915050565b620007668262000510565b67ffffffffffffffff8111156200078257620007816200025a565b5b6200078e82546200054a565b6200079b828285620006c0565b600060209050601f831160018114620007d35760008415620007be578287015190505b620007ca85826200073d565b8655506200083a565b601f198416620007e3866200057f565b60005b828110156200080d57848901518255600182019150602085019450602081019050620007e6565b868310156200082d578489015162000829601f8916826200071d565b8355505b6001600288020188555050505b505050505050565b6114d980620008526000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063baa40e5c11610066578063baa40e5c14610120578063c479fc9114610150578063cc2ee1961461016c578063cde71b781461018a578063f2fde38b146101ba57610093565b8063069953a71461009857806339bfeae3146100c8578063715018a6146100f85780638da5cb5b14610102575b600080fd5b6100b260048036038101906100ad919061099c565b6101d6565b6040516100bf9190610a04565b60405180910390f35b6100e260048036038101906100dd919061099c565b61025e565b6040516100ef9190610a3a565b60405180910390f35b610100610294565b005b61010a6102a8565b6040516101179190610a96565b60405180910390f35b61013a6004803603810190610135919061099c565b6102d1565b6040516101479190610b30565b60405180910390f35b61016a60048036038101906101659190610b52565b6103d1565b005b610174610511565b6040516101819190610cd6565b60405180910390f35b6101a4600480360381019061019f919061099c565b6105f2565b6040516101b19190610a3a565b60405180910390f35b6101d460048036038101906101cf9190610d24565b610675565b005b60006101e06106f8565b6101e9826105f2565b610228576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021f90610dc3565b60405180910390fd5b6002826040516102389190610e1f565b908152602001604051809103902060009054906101000a900463ffffffff169050919050565b6000806001836040516102719190610e1f565b9081526020016040518091039020805461028a90610e65565b9050119050919050565b61029c6106f8565b6102a66000610776565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606102db6106f8565b6102e48261025e565b610323576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031a90610ee2565b60405180910390fd5b6001826040516103339190610e1f565b9081526020016040518091039020805461034c90610e65565b80601f016020809104026020016040519081016040528092919081815260200182805461037890610e65565b80156103c55780601f1061039a576101008083540402835291602001916103c5565b820191906000526020600020905b8154815290600101906020018083116103a857829003601f168201915b50505050509050919050565b6103d96106f8565b6103e2816105f2565b610421576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041890610dc3565b60405180910390fd5b61042a8261025e565b1561046a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046190610f4e565b60405180910390fd5b8060018360405161047b9190610e1f565b908152602001604051809103902090816104959190611124565b5060016002826040516104a89190610e1f565b908152602001604051809103902060009054906101000a900463ffffffff166104d19190611225565b6002826040516104e19190610e1f565b908152602001604051809103902060006101000a81548163ffffffff021916908363ffffffff1602179055505050565b606061051b6106f8565b6003805480602002602001604051908101604052809291908181526020016000905b828210156105e957838290600052602060002001805461055c90610e65565b80601f016020809104026020016040519081016040528092919081815260200182805461058890610e65565b80156105d55780601f106105aa576101008083540402835291602001916105d5565b820191906000526020600020905b8154815290600101906020018083116105b857829003601f168201915b50505050508152602001906001019061053d565b50505050905090565b600080600090505b6003805490508160ff16101561066a57828051906020012060038260ff16815481106106295761062861125d565b5b90600052602060002001604051610640919061132f565b604051809103902003610657576001915050610670565b808061066290611353565b9150506105fa565b50600090505b919050565b61067d6106f8565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e3906113ee565b60405180910390fd5b6106f581610776565b50565b61070061083a565b73ffffffffffffffffffffffffffffffffffffffff1661071e6102a8565b73ffffffffffffffffffffffffffffffffffffffff1614610774576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076b9061145a565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6108a982610860565b810181811067ffffffffffffffff821117156108c8576108c7610871565b5b80604052505050565b60006108db610842565b90506108e782826108a0565b919050565b600067ffffffffffffffff82111561090757610906610871565b5b61091082610860565b9050602081019050919050565b82818337600083830152505050565b600061093f61093a846108ec565b6108d1565b90508281526020810184848401111561095b5761095a61085b565b5b61096684828561091d565b509392505050565b600082601f83011261098357610982610856565b5b813561099384826020860161092c565b91505092915050565b6000602082840312156109b2576109b161084c565b5b600082013567ffffffffffffffff8111156109d0576109cf610851565b5b6109dc8482850161096e565b91505092915050565b600063ffffffff82169050919050565b6109fe816109e5565b82525050565b6000602082019050610a1960008301846109f5565b92915050565b60008115159050919050565b610a3481610a1f565b82525050565b6000602082019050610a4f6000830184610a2b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a8082610a55565b9050919050565b610a9081610a75565b82525050565b6000602082019050610aab6000830184610a87565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610aeb578082015181840152602081019050610ad0565b60008484015250505050565b6000610b0282610ab1565b610b0c8185610abc565b9350610b1c818560208601610acd565b610b2581610860565b840191505092915050565b60006020820190508181036000830152610b4a8184610af7565b905092915050565b60008060408385031215610b6957610b6861084c565b5b600083013567ffffffffffffffff811115610b8757610b86610851565b5b610b938582860161096e565b925050602083013567ffffffffffffffff811115610bb457610bb3610851565b5b610bc08582860161096e565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000610c1282610ab1565b610c1c8185610bf6565b9350610c2c818560208601610acd565b610c3581610860565b840191505092915050565b6000610c4c8383610c07565b905092915050565b6000602082019050919050565b6000610c6c82610bca565b610c768185610bd5565b935083602082028501610c8885610be6565b8060005b85811015610cc45784840389528151610ca58582610c40565b9450610cb083610c54565b925060208a01995050600181019050610c8c565b50829750879550505050505092915050565b60006020820190508181036000830152610cf08184610c61565b905092915050565b610d0181610a75565b8114610d0c57600080fd5b50565b600081359050610d1e81610cf8565b92915050565b600060208284031215610d3a57610d3961084c565b5b6000610d4884828501610d0f565b91505092915050565b7f566f74653a2056616c75652073656e74206973206e6f7420616e204f7074696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000610dad602183610abc565b9150610db882610d51565b604082019050919050565b60006020820190508181036000830152610ddc81610da0565b9050919050565b600081905092915050565b6000610df982610ab1565b610e038185610de3565b9350610e13818560208601610acd565b80840191505092915050565b6000610e2b8284610dee565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e7d57607f821691505b602082108103610e9057610e8f610e36565b5b50919050565b7f566f74653a20566f746572436f6465206e6f7420666f756e6400000000000000600082015250565b6000610ecc601983610abc565b9150610ed782610e96565b602082019050919050565b60006020820190508181036000830152610efb81610ebf565b9050919050565b7f566f74653a20566f7465722068617320616c726561647920766f746564000000600082015250565b6000610f38601d83610abc565b9150610f4382610f02565b602082019050919050565b60006020820190508181036000830152610f6781610f2b565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610fd07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f93565b610fda8683610f93565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061102161101c61101784610ff2565b610ffc565b610ff2565b9050919050565b6000819050919050565b61103b83611006565b61104f61104782611028565b848454610fa0565b825550505050565b600090565b611064611057565b61106f818484611032565b505050565b5b818110156110935761108860008261105c565b600181019050611075565b5050565b601f8211156110d8576110a981610f6e565b6110b284610f83565b810160208510156110c1578190505b6110d56110cd85610f83565b830182611074565b50505b505050565b600082821c905092915050565b60006110fb600019846008026110dd565b1980831691505092915050565b600061111483836110ea565b9150826002028217905092915050565b61112d82610ab1565b67ffffffffffffffff81111561114657611145610871565b5b6111508254610e65565b61115b828285611097565b600060209050601f83116001811461118e576000841561117c578287015190505b6111868582611108565b8655506111ee565b601f19841661119c86610f6e565b60005b828110156111c45784890151825560018201915060208501945060208101905061119f565b868310156111e157848901516111dd601f8916826110ea565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611230826109e5565b915061123b836109e5565b9250828201905063ffffffff811115611257576112566111f6565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b60008190508160005260206000209050919050565b600081546112b981610e65565b6112c3818661128c565b945060018216600081146112de57600181146112f357611326565b60ff1983168652811515820286019350611326565b6112fc85611297565b60005b8381101561131e578154818901526001820191506020810190506112ff565b838801955050505b50505092915050565b600061133b82846112ac565b915081905092915050565b600060ff82169050919050565b600061135e82611346565b915060ff8203611371576113706111f6565b5b600182019050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006113d8602683610abc565b91506113e38261137c565b604082019050919050565b60006020820190508181036000830152611407816113cb565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611444602083610abc565b915061144f8261140e565b602082019050919050565b6000602082019050818103600083015261147381611437565b905091905056fea2646970667358221220e224a4147feefcc37beaf28a5b489f96deaaa37d8910f143812c1ce50f8d6aae64736f6c63782b302e382e31372d646576656c6f702e323032322e392e362b636f6d6d69742e65636463383038652e6d6f64005c",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _options []string) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _options)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetOptions is a free data retrieval call binding the contract method 0xcc2ee196.
//
// Solidity: function getOptions() view returns(string[])
func (_Api *ApiCaller) GetOptions(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getOptions")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetOptions is a free data retrieval call binding the contract method 0xcc2ee196.
//
// Solidity: function getOptions() view returns(string[])
func (_Api *ApiSession) GetOptions() ([]string, error) {
	return _Api.Contract.GetOptions(&_Api.CallOpts)
}

// GetOptions is a free data retrieval call binding the contract method 0xcc2ee196.
//
// Solidity: function getOptions() view returns(string[])
func (_Api *ApiCallerSession) GetOptions() ([]string, error) {
	return _Api.Contract.GetOptions(&_Api.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string voterCode) view returns(string)
func (_Api *ApiCaller) GetVote(opts *bind.CallOpts, voterCode string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getVote", voterCode)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string voterCode) view returns(string)
func (_Api *ApiSession) GetVote(voterCode string) (string, error) {
	return _Api.Contract.GetVote(&_Api.CallOpts, voterCode)
}

// GetVote is a free data retrieval call binding the contract method 0xbaa40e5c.
//
// Solidity: function getVote(string voterCode) view returns(string)
func (_Api *ApiCallerSession) GetVote(voterCode string) (string, error) {
	return _Api.Contract.GetVote(&_Api.CallOpts, voterCode)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string option) view returns(uint32)
func (_Api *ApiCaller) GetVoteCount(opts *bind.CallOpts, option string) (uint32, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getVoteCount", option)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string option) view returns(uint32)
func (_Api *ApiSession) GetVoteCount(option string) (uint32, error) {
	return _Api.Contract.GetVoteCount(&_Api.CallOpts, option)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string option) view returns(uint32)
func (_Api *ApiCallerSession) GetVoteCount(option string) (uint32, error) {
	return _Api.Contract.GetVoteCount(&_Api.CallOpts, option)
}

// HasVoted is a free data retrieval call binding the contract method 0x39bfeae3.
//
// Solidity: function hasVoted(string voterCode) view returns(bool)
func (_Api *ApiCaller) HasVoted(opts *bind.CallOpts, voterCode string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "hasVoted", voterCode)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x39bfeae3.
//
// Solidity: function hasVoted(string voterCode) view returns(bool)
func (_Api *ApiSession) HasVoted(voterCode string) (bool, error) {
	return _Api.Contract.HasVoted(&_Api.CallOpts, voterCode)
}

// HasVoted is a free data retrieval call binding the contract method 0x39bfeae3.
//
// Solidity: function hasVoted(string voterCode) view returns(bool)
func (_Api *ApiCallerSession) HasVoted(voterCode string) (bool, error) {
	return _Api.Contract.HasVoted(&_Api.CallOpts, voterCode)
}

// IsValidOption is a free data retrieval call binding the contract method 0xcde71b78.
//
// Solidity: function isValidOption(string option) view returns(bool)
func (_Api *ApiCaller) IsValidOption(opts *bind.CallOpts, option string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isValidOption", option)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOption is a free data retrieval call binding the contract method 0xcde71b78.
//
// Solidity: function isValidOption(string option) view returns(bool)
func (_Api *ApiSession) IsValidOption(option string) (bool, error) {
	return _Api.Contract.IsValidOption(&_Api.CallOpts, option)
}

// IsValidOption is a free data retrieval call binding the contract method 0xcde71b78.
//
// Solidity: function isValidOption(string option) view returns(bool)
func (_Api *ApiCallerSession) IsValidOption(option string) (bool, error) {
	return _Api.Contract.IsValidOption(&_Api.CallOpts, option)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns()
func (_Api *ApiTransactor) CastVote(opts *bind.TransactOpts, voterCode string, option string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "castVote", voterCode, option)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns()
func (_Api *ApiSession) CastVote(voterCode string, option string) (*types.Transaction, error) {
	return _Api.Contract.CastVote(&_Api.TransactOpts, voterCode, option)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns()
func (_Api *ApiTransactorSession) CastVote(voterCode string, option string) (*types.Transaction, error) {
	return _Api.Contract.CastVote(&_Api.TransactOpts, voterCode, option)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Api *ApiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Api.Contract.RenounceOwnership(&_Api.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Api *ApiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Api.Contract.TransferOwnership(&_Api.TransactOpts, newOwner)
}

// ApiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Api contract.
type ApiOwnershipTransferredIterator struct {
	Event *ApiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ApiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ApiOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ApiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiOwnershipTransferred represents a OwnershipTransferred event raised by the Api contract.
type ApiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApiOwnershipTransferredIterator{contract: _Api.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiOwnershipTransferred)
				if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Api *ApiFilterer) ParseOwnershipTransferred(log types.Log) (*ApiOwnershipTransferred, error) {
	event := new(ApiOwnershipTransferred)
	if err := _Api.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
