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
	ABI: "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_options\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"voterCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"voterCode\",\"type\":\"string\"}],\"name\":\"getVote\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"}],\"name\":\"getVoteCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a2838038062001a288339818101604052810190620000379190620004bf565b620000576200004b6200007760201b60201c565b6200007f60201b60201c565b80600390805190602001906200006f92919062000143565b505062000842565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b82805482825590600052602060002090810192821562000190579160200282015b828111156200018f5782518290816200017e91906200075b565b509160200191906001019062000164565b5b5090506200019f9190620001a3565b5090565b5b80821115620001c75760008181620001bd9190620001cb565b50600101620001a4565b5090565b508054620001d9906200054a565b6000825580601f10620001ed57506200020e565b601f0160209004906000526020600020908101906200020d919062000211565b5b50565b5b808211156200022c57600081600090555060010162000212565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620002948262000249565b810181811067ffffffffffffffff82111715620002b657620002b56200025a565b5b80604052505050565b6000620002cb62000230565b9050620002d9828262000289565b919050565b600067ffffffffffffffff821115620002fc57620002fb6200025a565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff8211156200033557620003346200025a565b5b620003408262000249565b9050602081019050919050565b60005b838110156200036d57808201518184015260208101905062000350565b60008484015250505050565b6000620003906200038a8462000317565b620002bf565b905082815260208101848484011115620003af57620003ae62000312565b5b620003bc8482856200034d565b509392505050565b600082601f830112620003dc57620003db62000244565b5b8151620003ee84826020860162000379565b91505092915050565b60006200040e6200040884620002de565b620002bf565b905080838252602082019050602084028301858111156200043457620004336200030d565b5b835b818110156200048257805167ffffffffffffffff8111156200045d576200045c62000244565b5b8086016200046c8982620003c4565b8552602085019450505060208101905062000436565b5050509392505050565b600082601f830112620004a457620004a362000244565b5b8151620004b6848260208601620003f7565b91505092915050565b600060208284031215620004d857620004d76200023a565b5b600082015167ffffffffffffffff811115620004f957620004f86200023f565b5b62000507848285016200048c565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200056357607f821691505b6020821081036200057957620005786200051b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005e37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005a4565b620005ef8683620005a4565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200063c62000636620006308462000607565b62000611565b62000607565b9050919050565b6000819050919050565b62000658836200061b565b62000670620006678262000643565b848454620005b1565b825550505050565b600090565b6200068762000678565b620006948184846200064d565b505050565b5b81811015620006bc57620006b06000826200067d565b6001810190506200069a565b5050565b601f8211156200070b57620006d5816200057f565b620006e08462000594565b81016020851015620006f0578190505b62000708620006ff8562000594565b83018262000699565b50505b505050565b600082821c905092915050565b6000620007306000198460080262000710565b1980831691505092915050565b60006200074b83836200071d565b9150826002028217905092915050565b620007668262000510565b67ffffffffffffffff8111156200078257620007816200025a565b5b6200078e82546200054a565b6200079b828285620006c0565b600060209050601f831160018114620007d35760008415620007be578287015190505b620007ca85826200073d565b8655506200083a565b601f198416620007e3866200057f565b60005b828110156200080d57848901518255600182019150602085019450602081019050620007e6565b868310156200082d578489015162000829601f8916826200071d565b8355505b6001600288020188555050505b505050505050565b6111d680620008526000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063069953a714610067578063715018a6146100975780638da5cb5b146100a1578063baa40e5c146100bf578063c479fc91146100ef578063f2fde38b1461011f575b600080fd5b610081600480360381019061007c9190610827565b61013b565b60405161008e919061088f565b60405180910390f35b61009f6101c3565b005b6100a96101d7565b6040516100b691906108eb565b60405180910390f35b6100d960048036038101906100d49190610827565b610200565b6040516100e69190610985565b60405180910390f35b610109600480360381019061010491906109a7565b610300565b6040516101169190610985565b60405180910390f35b61013960048036038101906101349190610a4b565b610447565b005b60006101456104ca565b61014e82610548565b61018d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018490610aea565b60405180910390fd5b60028260405161019d9190610b46565b908152602001604051809103902060009054906101000a900463ffffffff169050919050565b6101cb6104ca565b6101d560006105cb565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606061020a6104ca565b6102138261068f565b610252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024990610ba9565b60405180910390fd5b6001826040516102629190610b46565b9081526020016040518091039020805461027b90610bf8565b80601f01602080910402602001604051908101604052809291908181526020018280546102a790610bf8565b80156102f45780601f106102c9576101008083540402835291602001916102f4565b820191906000526020600020905b8154815290600101906020018083116102d757829003601f168201915b50505050509050919050565b606061030a6104ca565b61031382610548565b61033757604051806060016040528060218152602001611157602191399050610441565b6103408361068f565b15610382576040518060400160405280601d81526020017f566f74653a20566f7465722068617320616c726561647920766f7465640000008152509050610441565b816001846040516103939190610b46565b908152602001604051809103902090816103ad9190610ddf565b5060016002836040516103c09190610b46565b908152602001604051809103902060009054906101000a900463ffffffff166103e99190610ee0565b6002836040516103f99190610b46565b908152602001604051809103902060006101000a81548163ffffffff021916908363ffffffff1602179055506040518060600160405280602181526020016111366021913990505b92915050565b61044f6104ca565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036104be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b590610f8a565b60405180910390fd5b6104c7816105cb565b50565b6104d26106c5565b73ffffffffffffffffffffffffffffffffffffffff166104f06101d7565b73ffffffffffffffffffffffffffffffffffffffff1614610546576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053d90610ff6565b60405180910390fd5b565b600080600090505b6003805490508160ff1610156105c057828051906020012060038260ff168154811061057f5761057e611016565b5b9060005260206000200160405161059691906110e8565b6040518091039020036105ad5760019150506105c6565b80806105b89061110c565b915050610550565b50600090505b919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000806001836040516106a29190610b46565b908152602001604051809103902080546106bb90610bf8565b9050119050919050565b600033905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610734826106eb565b810181811067ffffffffffffffff82111715610753576107526106fc565b5b80604052505050565b60006107666106cd565b9050610772828261072b565b919050565b600067ffffffffffffffff821115610792576107916106fc565b5b61079b826106eb565b9050602081019050919050565b82818337600083830152505050565b60006107ca6107c584610777565b61075c565b9050828152602081018484840111156107e6576107e56106e6565b5b6107f18482856107a8565b509392505050565b600082601f83011261080e5761080d6106e1565b5b813561081e8482602086016107b7565b91505092915050565b60006020828403121561083d5761083c6106d7565b5b600082013567ffffffffffffffff81111561085b5761085a6106dc565b5b610867848285016107f9565b91505092915050565b600063ffffffff82169050919050565b61088981610870565b82525050565b60006020820190506108a46000830184610880565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108d5826108aa565b9050919050565b6108e5816108ca565b82525050565b600060208201905061090060008301846108dc565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610940578082015181840152602081019050610925565b60008484015250505050565b600061095782610906565b6109618185610911565b9350610971818560208601610922565b61097a816106eb565b840191505092915050565b6000602082019050818103600083015261099f818461094c565b905092915050565b600080604083850312156109be576109bd6106d7565b5b600083013567ffffffffffffffff8111156109dc576109db6106dc565b5b6109e8858286016107f9565b925050602083013567ffffffffffffffff811115610a0957610a086106dc565b5b610a15858286016107f9565b9150509250929050565b610a28816108ca565b8114610a3357600080fd5b50565b600081359050610a4581610a1f565b92915050565b600060208284031215610a6157610a606106d7565b5b6000610a6f84828501610a36565b91505092915050565b7f566f74653a2056616c75652073656e74206973206e6f7420616e204f7074696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000610ad4602183610911565b9150610adf82610a78565b604082019050919050565b60006020820190508181036000830152610b0381610ac7565b9050919050565b600081905092915050565b6000610b2082610906565b610b2a8185610b0a565b9350610b3a818560208601610922565b80840191505092915050565b6000610b528284610b15565b915081905092915050565b7f566f74653a20566f746572436f6465206e6f7420666f756e6400000000000000600082015250565b6000610b93601983610911565b9150610b9e82610b5d565b602082019050919050565b60006020820190508181036000830152610bc281610b86565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c1057607f821691505b602082108103610c2357610c22610bc9565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610c8b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610c4e565b610c958683610c4e565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610cdc610cd7610cd284610cad565b610cb7565b610cad565b9050919050565b6000819050919050565b610cf683610cc1565b610d0a610d0282610ce3565b848454610c5b565b825550505050565b600090565b610d1f610d12565b610d2a818484610ced565b505050565b5b81811015610d4e57610d43600082610d17565b600181019050610d30565b5050565b601f821115610d9357610d6481610c29565b610d6d84610c3e565b81016020851015610d7c578190505b610d90610d8885610c3e565b830182610d2f565b50505b505050565b600082821c905092915050565b6000610db660001984600802610d98565b1980831691505092915050565b6000610dcf8383610da5565b9150826002028217905092915050565b610de882610906565b67ffffffffffffffff811115610e0157610e006106fc565b5b610e0b8254610bf8565b610e16828285610d52565b600060209050601f831160018114610e495760008415610e37578287015190505b610e418582610dc3565b865550610ea9565b601f198416610e5786610c29565b60005b82811015610e7f57848901518255600182019150602085019450602081019050610e5a565b86831015610e9c5784890151610e98601f891682610da5565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610eeb82610870565b9150610ef683610870565b9250828201905063ffffffff811115610f1257610f11610eb1565b5b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610f74602683610911565b9150610f7f82610f18565b604082019050919050565b60006020820190508181036000830152610fa381610f67565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610fe0602083610911565b9150610feb82610faa565b602082019050919050565b6000602082019050818103600083015261100f81610fd3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b60008190508160005260206000209050919050565b6000815461107281610bf8565b61107c8186611045565b9450600182166000811461109757600181146110ac576110df565b60ff19831686528115158202860193506110df565b6110b585611050565b60005b838110156110d7578154818901526001820191506020810190506110b8565b838801955050505b50505092915050565b60006110f48284611065565b915081905092915050565b600060ff82169050919050565b6000611117826110ff565b915060ff820361112a57611129610eb1565b5b60018201905091905056fe566f74653a20566f746520726567697374657265642073756363657366756c6c79566f74653a2056616c75652073656e74206973206e6f7420616e204f7074696f6ea2646970667358221220f5c7cb258fa2d2fcf5a62aeb622f9a8e199f69e13ef83075b53e5312e7cad6b364736f6c63782b302e382e31372d646576656c6f702e323032322e392e362b636f6d6d69742e65636463383038652e6d6f64005c",
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
// Solidity: function castVote(string voterCode, string option) returns(string)
func (_Api *ApiTransactor) CastVote(opts *bind.TransactOpts, voterCode string, option string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "castVote", voterCode, option)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns(string)
func (_Api *ApiSession) CastVote(voterCode string, option string) (*types.Transaction, error) {
	return _Api.Contract.CastVote(&_Api.TransactOpts, voterCode, option)
}

// CastVote is a paid mutator transaction binding the contract method 0xc479fc91.
//
// Solidity: function castVote(string voterCode, string option) returns(string)
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
