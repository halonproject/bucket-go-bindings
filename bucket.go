// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/cpurta/go-ethereum"
	"github.com/cpurta/go-ethereum/accounts/abi"
	"github.com/cpurta/go-ethereum/accounts/abi/bind"
	"github.com/cpurta/go-ethereum/common"
	"github.com/cpurta/go-ethereum/core/types"
	"github.com/cpurta/go-ethereum/event"
)

// BucketABI is the input ABI used to generate the binding from.
const BucketABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"}],\"name\":\"getPermissionAddresses\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"}],\"name\":\"getFile\",\"outputs\":[{\"name\":\"storageRef\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"fileSize\",\"type\":\"uint256\"},{\"name\":\"isPublic\",\"type\":\"bool\"},{\"name\":\"isDeleted\",\"type\":\"bool\"},{\"name\":\"fileOwner\",\"type\":\"address\"},{\"name\":\"isOwner\",\"type\":\"bool\"},{\"name\":\"lastModified\",\"type\":\"uint256\"},{\"name\":\"permissionAddresses\",\"type\":\"address[]\"},{\"name\":\"writeAccess\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_storageRef\",\"type\":\"string\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_fileSize\",\"type\":\"uint256\"},{\"name\":\"_isPublic\",\"type\":\"bool\"}],\"name\":\"addFile\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"},{\"name\":\"_storageRef\",\"type\":\"string\"},{\"name\":\"_fileSize\",\"type\":\"uint256\"}],\"name\":\"setFileContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"},{\"name\":\"_entity\",\"type\":\"address\"},{\"name\":\"_permission\",\"type\":\"bool\"}],\"name\":\"setReadPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"},{\"name\":\"_entity\",\"type\":\"address\"}],\"name\":\"hasReadAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"}],\"name\":\"removeFile\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"},{\"name\":\"_newName\",\"type\":\"string\"}],\"name\":\"setFileName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"},{\"name\":\"_entity\",\"type\":\"address\"}],\"name\":\"getPermission\",\"outputs\":[{\"name\":\"write\",\"type\":\"bool\"},{\"name\":\"read\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastFileId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"},{\"name\":\"_entity\",\"type\":\"address\"}],\"name\":\"hasWriteAccess\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fileId\",\"type\":\"uint256\"},{\"name\":\"_entity\",\"type\":\"address\"},{\"name\":\"_permission\",\"type\":\"bool\"}],\"name\":\"setWritePermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"FileRename\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"FileContentUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"NewFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"NewWritePermission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"NewReadPermission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fileId\",\"type\":\"uint256\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BucketBin is the compiled bytecode used for deploying new contracts.
const BucketBin = `0x6080604052600060025534801561001557600080fd5b506000805433600160a060020a0319918216811783556001805490921617905561141990819061004590396000f3006080604052600436106100e55763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663104a2a6081146100ea5780632bfda3131461015257806331c02db2146102dc578063397c2eff1461038c5780635b11e4c9146103ee57806369366e5714610417578063715018a61461044f57806383197ef0146104645780638da5cb5b146104795780639926b03f146104aa578063b19902d3146104c2578063e0737dfd14610520578063e14cabcc1461055f578063e3f906cd14610574578063eb3624fe14610598578063f2fde38b146105c1575b600080fd5b3480156100f657600080fd5b506101026004356105e2565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561013e578181015183820152602001610126565b505050509050019250505060405180910390f35b34801561015e57600080fd5b5061016a600435610651565b6040518080602001806020018b81526020018a1515151581526020018915151515815260200188600160a060020a0316600160a060020a0316815260200187151515158152602001868152602001806020018515151515815260200184810384528e818151815260200191508051906020019080838360005b838110156101fb5781810151838201526020016101e3565b50505050905090810190601f1680156102285780820380516001836020036101000a031916815260200191505b5084810383528d5181528d516020918201918f019080838360005b8381101561025b578181015183820152602001610243565b50505050905090810190601f1680156102885780820380516001836020036101000a031916815260200191505b508481038252865181528651602091820191808901910280838360005b838110156102bd5781810151838201526020016102a5565b505050509050019d505050505050505050505050505060405180910390f35b3480156102e857600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261037a94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750508435955050505050602001351515610897565b60408051918252519081900360200190f35b34801561039857600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526103ec9583359536956044949193909101919081908401838280828437509497505093359450610a649350505050565b005b3480156103fa57600080fd5b506103ec600435600160a060020a03602435166044351515610b40565b34801561042357600080fd5b5061043b600435600160a060020a0360243516610c99565b604080519115158252519081900360200190f35b34801561045b57600080fd5b506103ec610cdb565b34801561047057600080fd5b506103ec610d80565b34801561048557600080fd5b5061048e610dde565b60408051600160a060020a039092168252519081900360200190f35b3480156104b657600080fd5b506103ec600435610ded565b3480156104ce57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526103ec958335953695604494919390910191908190840183828082843750949750610eae9650505050505050565b34801561052c57600080fd5b50610544600435600160a060020a0360243516610f86565b60408051921515835290151560208301528051918290030190f35b34801561056b57600080fd5b5061037a610fbd565b34801561058057600080fd5b5061043b600435600160a060020a0360243516610fc3565b3480156105a457600080fd5b506103ec600435600160a060020a03602435166044351515611009565b3480156105cd57600080fd5b506103ec600160a060020a0360043516611167565b60008181526003602090815260409182902060060180548351818402810184019094528084526060939283018282801561064557602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610627575b50505050509050919050565b600081815260036020908152604080832080548251601f6002600019610100600186161502019093169290920491820185900485028101850190935280835260609485949093849384938493849384938a9385939192918391908301828280156106fc5780601f106106d1576101008083540402835291602001916106fc565b820191906000526020600020905b8154815290600101906020018083116106df57829003601f168201915b50505050509a50806001018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561079b5780601f106107705761010080835404028352916020019161079b565b820191906000526020600020905b81548152906001019060200180831161077e57829003601f168201915b50505050509950806002015498508060030160009054906101000a900460ff1697508060030160019054906101000a900460ff1696508060030160029054906101000a9004600160a060020a03169550600160009054906101000a9004600160a060020a0316600160a060020a031633600160a060020a0316149450806004015493508060060180548060200260200160405190810160405280929190818152602001828054801561087657602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610858575b505050505092506108878c33610fc3565b9150509193959799509193959799565b60008054600160a060020a031633146108e8576040805160e560020a62461bcd02815260206004820152601c60248201526000805160206113ce833981519152604482015290519081900360640190fd5b600280546001908101918290556040805161010081018252888152602080820189905281830188905286151560608301526000608083018190529354600160a060020a031660a08301524260c08301528251848152808201845260e083015293835260038452912081518051929391926109659284920190611293565b50602082810151805161097e9260018501920190611293565b50604082015160028201556060820151600382018054608085015160a0860151600160a060020a0316620100000275ffffffffffffffffffffffffffffffffffffffff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c0820151600482015560e08201518051610a10916006840191602090910190611311565b50506001546002546040805191825251600160a060020a0390921692507f2c618b3515ba4be58b653279244757f1b863bf68901fee1deb8acc4112200a1f919081900360200190a250600254949350505050565b610a6e8333610fc3565b1515610ac4576040805160e560020a62461bcd02815260206004820152601560248201527f73656e646572206e6f7420617574686f72697a65640000000000000000000000604482015290519081900360640190fd5b60008381526003602090815260409091208351610ae392850190611293565b50600083815260036020908152604091829020600280820185905542600490920191909155548251908152915133927fbff1f689d7fd5c80d3f93cbce8d1c1add1b8312c5d6ac5b885f469bedee3f6a892908290030190a2505050565b600054600160a060020a03163314610b90576040805160e560020a62461bcd02815260206004820152601c60248201526000805160206113ce833981519152604482015290519081900360640190fd5b6000838152600360209081526040808320600160a060020a038616845260050190915290205462010000900460ff161515610c2d576000838152600360209081526040808320600681018054600181018255908552838520018054600160a060020a03881673ffffffffffffffffffffffffffffffffffffffff19909116811790915584526005019091529020805462ff00001916620100001790555b6000838152600360209081526040808320600160a060020a0386168452600501825291829020805460ff19168415151790556002548251908152915133927f9a6884615d185a700778b977ce626f0444ca77e3537754baded1de887d06a2be92908290030190a2505050565b6000610ca48261127f565b80610cd457506000838152600360209081526040808320600160a060020a038616845260050190915290205460ff165b9392505050565b600054600160a060020a03163314610d2b576040805160e560020a62461bcd02815260206004820152601c60248201526000805160206113ce833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a03163314610dd0576040805160e560020a62461bcd02815260206004820152601c60248201526000805160206113ce833981519152604482015290519081900360640190fd5b600154600160a060020a0316ff5b600154600160a060020a031681565b600054600160a060020a03163314610e3d576040805160e560020a62461bcd02815260206004820152601c60248201526000805160206113ce833981519152604482015290519081900360640190fd5b600081815260036020818152604092839020918201805461ff0019166101001790554260049092019190915560015460025483519081529251600160a060020a03909116927f264a48949e72079f96655477155fb6e35181831bdb63563d610bb34a575c3f6c92908290030190a250565b610eb88233610fc3565b1515610f0e576040805160e560020a62461bcd02815260206004820152601560248201527f73656e646572206e6f7420617574686f72697a65640000000000000000000000604482015290519081900360640190fd5b60008281526003602090815260409091208251610f3392600190920191840190611293565b50600082815260036020908152604091829020426004909101556002548251908152915133927fbff1f689d7fd5c80d3f93cbce8d1c1add1b8312c5d6ac5b885f469bedee3f6a892908290030190a25050565b6000918252600360209081526040808420600160a060020a0390931684526005909201905290205460ff6101008204811692911690565b60025481565b6000610fce8261127f565b80610cd45750506000918252600360209081526040808420600160a060020a0393909316845260059092019052902054610100900460ff1690565b600054600160a060020a03163314611059576040805160e560020a62461bcd02815260206004820152601c60248201526000805160206113ce833981519152604482015290519081900360640190fd5b6000838152600360209081526040808320600160a060020a038616845260050190915290205462010000900460ff1615156110f6576000838152600360209081526040808320600681018054600181018255908552838520018054600160a060020a03881673ffffffffffffffffffffffffffffffffffffffff19909116811790915584526005019091529020805462ff00001916620100001790555b6000838152600360209081526040808320600160a060020a0386168452600501825291829020805461ff001916610100851515021790556002548251908152915133927f9a6884615d185a700778b977ce626f0444ca77e3537754baded1de887d06a2be92908290030190a2505050565b600054600160a060020a031633146111b7576040805160e560020a62461bcd02815260206004820152601c60248201526000805160206113ce833981519152604482015290519081900360640190fd5b600160a060020a0381161515611217576040805160e560020a62461bcd02815260206004820152601e60248201527f63616e6e6f74206d616b6520302061646472657373206173206f776e65720000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a0390811691161490565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112d457805160ff1916838001178555611301565b82800160010185558215611301579182015b828111156113015782518255916020019190600101906112e6565b5061130d92915061137f565b5090565b828054828255906000526020600020908101928215611373579160200282015b82811115611373578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190611331565b5061130d92915061139c565b61139991905b8082111561130d5760008155600101611385565b90565b61139991905b8082111561130d57805473ffffffffffffffffffffffffffffffffffffffff191681556001016113a2560073656e646572206973206e6f7420636f6e7472616374206f776e657200000000a165627a7a723058205827dbeb240caf3b42bbab3ebe474f4ea8b4c7ffb6bf265c3f5a2416bc6eab620029`

// DeployBucket deploys a new Ethereum contract, binding an instance of Bucket to it.
func DeployBucket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bucket, error) {
	parsed, err := abi.JSON(strings.NewReader(BucketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BucketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bucket{BucketCaller: BucketCaller{contract: contract}, BucketTransactor: BucketTransactor{contract: contract}, BucketFilterer: BucketFilterer{contract: contract}}, nil
}

// Bucket is an auto generated Go binding around an Ethereum contract.
type Bucket struct {
	BucketCaller     // Read-only binding to the contract
	BucketTransactor // Write-only binding to the contract
	BucketFilterer   // Log filterer for contract events
}

// BucketCaller is an auto generated read-only Go binding around an Ethereum contract.
type BucketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BucketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BucketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BucketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BucketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BucketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BucketSession struct {
	Contract     *Bucket           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BucketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BucketCallerSession struct {
	Contract *BucketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BucketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BucketTransactorSession struct {
	Contract     *BucketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BucketRaw is an auto generated low-level Go binding around an Ethereum contract.
type BucketRaw struct {
	Contract *Bucket // Generic contract binding to access the raw methods on
}

// BucketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BucketCallerRaw struct {
	Contract *BucketCaller // Generic read-only contract binding to access the raw methods on
}

// BucketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BucketTransactorRaw struct {
	Contract *BucketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBucket creates a new instance of Bucket, bound to a specific deployed contract.
func NewBucket(address common.Address, backend bind.ContractBackend) (*Bucket, error) {
	contract, err := bindBucket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bucket{BucketCaller: BucketCaller{contract: contract}, BucketTransactor: BucketTransactor{contract: contract}, BucketFilterer: BucketFilterer{contract: contract}}, nil
}

// NewBucketCaller creates a new read-only instance of Bucket, bound to a specific deployed contract.
func NewBucketCaller(address common.Address, caller bind.ContractCaller) (*BucketCaller, error) {
	contract, err := bindBucket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BucketCaller{contract: contract}, nil
}

// NewBucketTransactor creates a new write-only instance of Bucket, bound to a specific deployed contract.
func NewBucketTransactor(address common.Address, transactor bind.ContractTransactor) (*BucketTransactor, error) {
	contract, err := bindBucket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BucketTransactor{contract: contract}, nil
}

// NewBucketFilterer creates a new log filterer instance of Bucket, bound to a specific deployed contract.
func NewBucketFilterer(address common.Address, filterer bind.ContractFilterer) (*BucketFilterer, error) {
	contract, err := bindBucket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BucketFilterer{contract: contract}, nil
}

// bindBucket binds a generic wrapper to an already deployed contract.
func bindBucket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BucketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bucket *BucketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bucket.Contract.BucketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bucket *BucketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bucket.Contract.BucketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bucket *BucketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bucket.Contract.BucketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bucket *BucketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bucket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bucket *BucketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bucket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bucket *BucketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bucket.Contract.contract.Transact(opts, method, params...)
}

// GetFile is a free data retrieval call binding the contract method 0x2bfda313.
//
// Solidity: function getFile(_fileId uint256) constant returns(storageRef string, name string, fileSize uint256, isPublic bool, isDeleted bool, fileOwner address, isOwner bool, lastModified uint256, permissionAddresses address[], writeAccess bool)
func (_Bucket *BucketCaller) GetFile(opts *bind.CallOpts, _fileId *big.Int) (struct {
	StorageRef          string
	Name                string
	FileSize            *big.Int
	IsPublic            bool
	IsDeleted           bool
	FileOwner           common.Address
	IsOwner             bool
	LastModified        *big.Int
	PermissionAddresses []common.Address
	WriteAccess         bool
}, error) {
	ret := new(struct {
		StorageRef          string
		Name                string
		FileSize            *big.Int
		IsPublic            bool
		IsDeleted           bool
		FileOwner           common.Address
		IsOwner             bool
		LastModified        *big.Int
		PermissionAddresses []common.Address
		WriteAccess         bool
	})
	out := ret
	err := _Bucket.contract.Call(opts, out, "getFile", _fileId)
	return *ret, err
}

// GetFile is a free data retrieval call binding the contract method 0x2bfda313.
//
// Solidity: function getFile(_fileId uint256) constant returns(storageRef string, name string, fileSize uint256, isPublic bool, isDeleted bool, fileOwner address, isOwner bool, lastModified uint256, permissionAddresses address[], writeAccess bool)
func (_Bucket *BucketSession) GetFile(_fileId *big.Int) (struct {
	StorageRef          string
	Name                string
	FileSize            *big.Int
	IsPublic            bool
	IsDeleted           bool
	FileOwner           common.Address
	IsOwner             bool
	LastModified        *big.Int
	PermissionAddresses []common.Address
	WriteAccess         bool
}, error) {
	return _Bucket.Contract.GetFile(&_Bucket.CallOpts, _fileId)
}

// GetFile is a free data retrieval call binding the contract method 0x2bfda313.
//
// Solidity: function getFile(_fileId uint256) constant returns(storageRef string, name string, fileSize uint256, isPublic bool, isDeleted bool, fileOwner address, isOwner bool, lastModified uint256, permissionAddresses address[], writeAccess bool)
func (_Bucket *BucketCallerSession) GetFile(_fileId *big.Int) (struct {
	StorageRef          string
	Name                string
	FileSize            *big.Int
	IsPublic            bool
	IsDeleted           bool
	FileOwner           common.Address
	IsOwner             bool
	LastModified        *big.Int
	PermissionAddresses []common.Address
	WriteAccess         bool
}, error) {
	return _Bucket.Contract.GetFile(&_Bucket.CallOpts, _fileId)
}

// GetPermission is a free data retrieval call binding the contract method 0xe0737dfd.
//
// Solidity: function getPermission(_fileId uint256, _entity address) constant returns(write bool, read bool)
func (_Bucket *BucketCaller) GetPermission(opts *bind.CallOpts, _fileId *big.Int, _entity common.Address) (struct {
	Write bool
	Read  bool
}, error) {
	ret := new(struct {
		Write bool
		Read  bool
	})
	out := ret
	err := _Bucket.contract.Call(opts, out, "getPermission", _fileId, _entity)
	return *ret, err
}

// GetPermission is a free data retrieval call binding the contract method 0xe0737dfd.
//
// Solidity: function getPermission(_fileId uint256, _entity address) constant returns(write bool, read bool)
func (_Bucket *BucketSession) GetPermission(_fileId *big.Int, _entity common.Address) (struct {
	Write bool
	Read  bool
}, error) {
	return _Bucket.Contract.GetPermission(&_Bucket.CallOpts, _fileId, _entity)
}

// GetPermission is a free data retrieval call binding the contract method 0xe0737dfd.
//
// Solidity: function getPermission(_fileId uint256, _entity address) constant returns(write bool, read bool)
func (_Bucket *BucketCallerSession) GetPermission(_fileId *big.Int, _entity common.Address) (struct {
	Write bool
	Read  bool
}, error) {
	return _Bucket.Contract.GetPermission(&_Bucket.CallOpts, _fileId, _entity)
}

// GetPermissionAddresses is a free data retrieval call binding the contract method 0x104a2a60.
//
// Solidity: function getPermissionAddresses(_fileId uint256) constant returns(addresses address[])
func (_Bucket *BucketCaller) GetPermissionAddresses(opts *bind.CallOpts, _fileId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Bucket.contract.Call(opts, out, "getPermissionAddresses", _fileId)
	return *ret0, err
}

// GetPermissionAddresses is a free data retrieval call binding the contract method 0x104a2a60.
//
// Solidity: function getPermissionAddresses(_fileId uint256) constant returns(addresses address[])
func (_Bucket *BucketSession) GetPermissionAddresses(_fileId *big.Int) ([]common.Address, error) {
	return _Bucket.Contract.GetPermissionAddresses(&_Bucket.CallOpts, _fileId)
}

// GetPermissionAddresses is a free data retrieval call binding the contract method 0x104a2a60.
//
// Solidity: function getPermissionAddresses(_fileId uint256) constant returns(addresses address[])
func (_Bucket *BucketCallerSession) GetPermissionAddresses(_fileId *big.Int) ([]common.Address, error) {
	return _Bucket.Contract.GetPermissionAddresses(&_Bucket.CallOpts, _fileId)
}

// HasReadAccess is a free data retrieval call binding the contract method 0x69366e57.
//
// Solidity: function hasReadAccess(_fileId uint256, _entity address) constant returns(bool)
func (_Bucket *BucketCaller) HasReadAccess(opts *bind.CallOpts, _fileId *big.Int, _entity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bucket.contract.Call(opts, out, "hasReadAccess", _fileId, _entity)
	return *ret0, err
}

// HasReadAccess is a free data retrieval call binding the contract method 0x69366e57.
//
// Solidity: function hasReadAccess(_fileId uint256, _entity address) constant returns(bool)
func (_Bucket *BucketSession) HasReadAccess(_fileId *big.Int, _entity common.Address) (bool, error) {
	return _Bucket.Contract.HasReadAccess(&_Bucket.CallOpts, _fileId, _entity)
}

// HasReadAccess is a free data retrieval call binding the contract method 0x69366e57.
//
// Solidity: function hasReadAccess(_fileId uint256, _entity address) constant returns(bool)
func (_Bucket *BucketCallerSession) HasReadAccess(_fileId *big.Int, _entity common.Address) (bool, error) {
	return _Bucket.Contract.HasReadAccess(&_Bucket.CallOpts, _fileId, _entity)
}

// HasWriteAccess is a free data retrieval call binding the contract method 0xe3f906cd.
//
// Solidity: function hasWriteAccess(_fileId uint256, _entity address) constant returns(bool)
func (_Bucket *BucketCaller) HasWriteAccess(opts *bind.CallOpts, _fileId *big.Int, _entity common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bucket.contract.Call(opts, out, "hasWriteAccess", _fileId, _entity)
	return *ret0, err
}

// HasWriteAccess is a free data retrieval call binding the contract method 0xe3f906cd.
//
// Solidity: function hasWriteAccess(_fileId uint256, _entity address) constant returns(bool)
func (_Bucket *BucketSession) HasWriteAccess(_fileId *big.Int, _entity common.Address) (bool, error) {
	return _Bucket.Contract.HasWriteAccess(&_Bucket.CallOpts, _fileId, _entity)
}

// HasWriteAccess is a free data retrieval call binding the contract method 0xe3f906cd.
//
// Solidity: function hasWriteAccess(_fileId uint256, _entity address) constant returns(bool)
func (_Bucket *BucketCallerSession) HasWriteAccess(_fileId *big.Int, _entity common.Address) (bool, error) {
	return _Bucket.Contract.HasWriteAccess(&_Bucket.CallOpts, _fileId, _entity)
}

// LastFileId is a free data retrieval call binding the contract method 0xe14cabcc.
//
// Solidity: function lastFileId() constant returns(uint256)
func (_Bucket *BucketCaller) LastFileId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bucket.contract.Call(opts, out, "lastFileId")
	return *ret0, err
}

// LastFileId is a free data retrieval call binding the contract method 0xe14cabcc.
//
// Solidity: function lastFileId() constant returns(uint256)
func (_Bucket *BucketSession) LastFileId() (*big.Int, error) {
	return _Bucket.Contract.LastFileId(&_Bucket.CallOpts)
}

// LastFileId is a free data retrieval call binding the contract method 0xe14cabcc.
//
// Solidity: function lastFileId() constant returns(uint256)
func (_Bucket *BucketCallerSession) LastFileId() (*big.Int, error) {
	return _Bucket.Contract.LastFileId(&_Bucket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bucket *BucketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bucket.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bucket *BucketSession) Owner() (common.Address, error) {
	return _Bucket.Contract.Owner(&_Bucket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bucket *BucketCallerSession) Owner() (common.Address, error) {
	return _Bucket.Contract.Owner(&_Bucket.CallOpts)
}

// AddFile is a paid mutator transaction binding the contract method 0x31c02db2.
//
// Solidity: function addFile(_storageRef string, _name string, _fileSize uint256, _isPublic bool) returns(uint256)
func (_Bucket *BucketTransactor) AddFile(opts *bind.TransactOpts, _storageRef string, _name string, _fileSize *big.Int, _isPublic bool) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "addFile", _storageRef, _name, _fileSize, _isPublic)
}

// AddFile is a paid mutator transaction binding the contract method 0x31c02db2.
//
// Solidity: function addFile(_storageRef string, _name string, _fileSize uint256, _isPublic bool) returns(uint256)
func (_Bucket *BucketSession) AddFile(_storageRef string, _name string, _fileSize *big.Int, _isPublic bool) (*types.Transaction, error) {
	return _Bucket.Contract.AddFile(&_Bucket.TransactOpts, _storageRef, _name, _fileSize, _isPublic)
}

// AddFile is a paid mutator transaction binding the contract method 0x31c02db2.
//
// Solidity: function addFile(_storageRef string, _name string, _fileSize uint256, _isPublic bool) returns(uint256)
func (_Bucket *BucketTransactorSession) AddFile(_storageRef string, _name string, _fileSize *big.Int, _isPublic bool) (*types.Transaction, error) {
	return _Bucket.Contract.AddFile(&_Bucket.TransactOpts, _storageRef, _name, _fileSize, _isPublic)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Bucket *BucketTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Bucket *BucketSession) Destroy() (*types.Transaction, error) {
	return _Bucket.Contract.Destroy(&_Bucket.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Bucket *BucketTransactorSession) Destroy() (*types.Transaction, error) {
	return _Bucket.Contract.Destroy(&_Bucket.TransactOpts)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(_fileId uint256) returns()
func (_Bucket *BucketTransactor) RemoveFile(opts *bind.TransactOpts, _fileId *big.Int) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "removeFile", _fileId)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(_fileId uint256) returns()
func (_Bucket *BucketSession) RemoveFile(_fileId *big.Int) (*types.Transaction, error) {
	return _Bucket.Contract.RemoveFile(&_Bucket.TransactOpts, _fileId)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(_fileId uint256) returns()
func (_Bucket *BucketTransactorSession) RemoveFile(_fileId *big.Int) (*types.Transaction, error) {
	return _Bucket.Contract.RemoveFile(&_Bucket.TransactOpts, _fileId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bucket *BucketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bucket *BucketSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bucket.Contract.RenounceOwnership(&_Bucket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bucket *BucketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bucket.Contract.RenounceOwnership(&_Bucket.TransactOpts)
}

// SetFileContent is a paid mutator transaction binding the contract method 0x397c2eff.
//
// Solidity: function setFileContent(_fileId uint256, _storageRef string, _fileSize uint256) returns()
func (_Bucket *BucketTransactor) SetFileContent(opts *bind.TransactOpts, _fileId *big.Int, _storageRef string, _fileSize *big.Int) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "setFileContent", _fileId, _storageRef, _fileSize)
}

// SetFileContent is a paid mutator transaction binding the contract method 0x397c2eff.
//
// Solidity: function setFileContent(_fileId uint256, _storageRef string, _fileSize uint256) returns()
func (_Bucket *BucketSession) SetFileContent(_fileId *big.Int, _storageRef string, _fileSize *big.Int) (*types.Transaction, error) {
	return _Bucket.Contract.SetFileContent(&_Bucket.TransactOpts, _fileId, _storageRef, _fileSize)
}

// SetFileContent is a paid mutator transaction binding the contract method 0x397c2eff.
//
// Solidity: function setFileContent(_fileId uint256, _storageRef string, _fileSize uint256) returns()
func (_Bucket *BucketTransactorSession) SetFileContent(_fileId *big.Int, _storageRef string, _fileSize *big.Int) (*types.Transaction, error) {
	return _Bucket.Contract.SetFileContent(&_Bucket.TransactOpts, _fileId, _storageRef, _fileSize)
}

// SetFileName is a paid mutator transaction binding the contract method 0xb19902d3.
//
// Solidity: function setFileName(_fileId uint256, _newName string) returns()
func (_Bucket *BucketTransactor) SetFileName(opts *bind.TransactOpts, _fileId *big.Int, _newName string) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "setFileName", _fileId, _newName)
}

// SetFileName is a paid mutator transaction binding the contract method 0xb19902d3.
//
// Solidity: function setFileName(_fileId uint256, _newName string) returns()
func (_Bucket *BucketSession) SetFileName(_fileId *big.Int, _newName string) (*types.Transaction, error) {
	return _Bucket.Contract.SetFileName(&_Bucket.TransactOpts, _fileId, _newName)
}

// SetFileName is a paid mutator transaction binding the contract method 0xb19902d3.
//
// Solidity: function setFileName(_fileId uint256, _newName string) returns()
func (_Bucket *BucketTransactorSession) SetFileName(_fileId *big.Int, _newName string) (*types.Transaction, error) {
	return _Bucket.Contract.SetFileName(&_Bucket.TransactOpts, _fileId, _newName)
}

// SetReadPermission is a paid mutator transaction binding the contract method 0x5b11e4c9.
//
// Solidity: function setReadPermission(_fileId uint256, _entity address, _permission bool) returns()
func (_Bucket *BucketTransactor) SetReadPermission(opts *bind.TransactOpts, _fileId *big.Int, _entity common.Address, _permission bool) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "setReadPermission", _fileId, _entity, _permission)
}

// SetReadPermission is a paid mutator transaction binding the contract method 0x5b11e4c9.
//
// Solidity: function setReadPermission(_fileId uint256, _entity address, _permission bool) returns()
func (_Bucket *BucketSession) SetReadPermission(_fileId *big.Int, _entity common.Address, _permission bool) (*types.Transaction, error) {
	return _Bucket.Contract.SetReadPermission(&_Bucket.TransactOpts, _fileId, _entity, _permission)
}

// SetReadPermission is a paid mutator transaction binding the contract method 0x5b11e4c9.
//
// Solidity: function setReadPermission(_fileId uint256, _entity address, _permission bool) returns()
func (_Bucket *BucketTransactorSession) SetReadPermission(_fileId *big.Int, _entity common.Address, _permission bool) (*types.Transaction, error) {
	return _Bucket.Contract.SetReadPermission(&_Bucket.TransactOpts, _fileId, _entity, _permission)
}

// SetWritePermission is a paid mutator transaction binding the contract method 0xeb3624fe.
//
// Solidity: function setWritePermission(_fileId uint256, _entity address, _permission bool) returns()
func (_Bucket *BucketTransactor) SetWritePermission(opts *bind.TransactOpts, _fileId *big.Int, _entity common.Address, _permission bool) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "setWritePermission", _fileId, _entity, _permission)
}

// SetWritePermission is a paid mutator transaction binding the contract method 0xeb3624fe.
//
// Solidity: function setWritePermission(_fileId uint256, _entity address, _permission bool) returns()
func (_Bucket *BucketSession) SetWritePermission(_fileId *big.Int, _entity common.Address, _permission bool) (*types.Transaction, error) {
	return _Bucket.Contract.SetWritePermission(&_Bucket.TransactOpts, _fileId, _entity, _permission)
}

// SetWritePermission is a paid mutator transaction binding the contract method 0xeb3624fe.
//
// Solidity: function setWritePermission(_fileId uint256, _entity address, _permission bool) returns()
func (_Bucket *BucketTransactorSession) SetWritePermission(_fileId *big.Int, _entity common.Address, _permission bool) (*types.Transaction, error) {
	return _Bucket.Contract.SetWritePermission(&_Bucket.TransactOpts, _fileId, _entity, _permission)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bucket *BucketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bucket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bucket *BucketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bucket.Contract.TransferOwnership(&_Bucket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bucket *BucketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bucket.Contract.TransferOwnership(&_Bucket.TransactOpts, newOwner)
}

// BucketDeleteFileIterator is returned from FilterDeleteFile and is used to iterate over the raw logs and unpacked data for DeleteFile events raised by the Bucket contract.
type BucketDeleteFileIterator struct {
	Event *BucketDeleteFile // Event containing the contract specifics and raw log

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
func (it *BucketDeleteFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketDeleteFile)
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
		it.Event = new(BucketDeleteFile)
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
func (it *BucketDeleteFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketDeleteFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketDeleteFile represents a DeleteFile event raised by the Bucket contract.
type BucketDeleteFile struct {
	Entity common.Address
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeleteFile is a free log retrieval operation binding the contract event 0x264a48949e72079f96655477155fb6e35181831bdb63563d610bb34a575c3f6c.
//
// Solidity: e DeleteFile(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) FilterDeleteFile(opts *bind.FilterOpts, entity []common.Address) (*BucketDeleteFileIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "DeleteFile", entityRule)
	if err != nil {
		return nil, err
	}
	return &BucketDeleteFileIterator{contract: _Bucket.contract, event: "DeleteFile", logs: logs, sub: sub}, nil
}

// WatchDeleteFile is a free log subscription operation binding the contract event 0x264a48949e72079f96655477155fb6e35181831bdb63563d610bb34a575c3f6c.
//
// Solidity: e DeleteFile(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) WatchDeleteFile(opts *bind.WatchOpts, sink chan<- *BucketDeleteFile, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "DeleteFile", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketDeleteFile)
				if err := _Bucket.contract.UnpackLog(event, "DeleteFile", log); err != nil {
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

// BucketFileContentUpdateIterator is returned from FilterFileContentUpdate and is used to iterate over the raw logs and unpacked data for FileContentUpdate events raised by the Bucket contract.
type BucketFileContentUpdateIterator struct {
	Event *BucketFileContentUpdate // Event containing the contract specifics and raw log

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
func (it *BucketFileContentUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketFileContentUpdate)
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
		it.Event = new(BucketFileContentUpdate)
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
func (it *BucketFileContentUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketFileContentUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketFileContentUpdate represents a FileContentUpdate event raised by the Bucket contract.
type BucketFileContentUpdate struct {
	Entity common.Address
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFileContentUpdate is a free log retrieval operation binding the contract event 0x9db4d4af9746fa572d0bf7fa1ad5c68dc2136946392b1bd94bd83ffaa8ab078b.
//
// Solidity: e FileContentUpdate(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) FilterFileContentUpdate(opts *bind.FilterOpts, entity []common.Address) (*BucketFileContentUpdateIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "FileContentUpdate", entityRule)
	if err != nil {
		return nil, err
	}
	return &BucketFileContentUpdateIterator{contract: _Bucket.contract, event: "FileContentUpdate", logs: logs, sub: sub}, nil
}

// WatchFileContentUpdate is a free log subscription operation binding the contract event 0x9db4d4af9746fa572d0bf7fa1ad5c68dc2136946392b1bd94bd83ffaa8ab078b.
//
// Solidity: e FileContentUpdate(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) WatchFileContentUpdate(opts *bind.WatchOpts, sink chan<- *BucketFileContentUpdate, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "FileContentUpdate", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketFileContentUpdate)
				if err := _Bucket.contract.UnpackLog(event, "FileContentUpdate", log); err != nil {
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

// BucketFileRenameIterator is returned from FilterFileRename and is used to iterate over the raw logs and unpacked data for FileRename events raised by the Bucket contract.
type BucketFileRenameIterator struct {
	Event *BucketFileRename // Event containing the contract specifics and raw log

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
func (it *BucketFileRenameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketFileRename)
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
		it.Event = new(BucketFileRename)
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
func (it *BucketFileRenameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketFileRenameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketFileRename represents a FileRename event raised by the Bucket contract.
type BucketFileRename struct {
	Entity common.Address
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFileRename is a free log retrieval operation binding the contract event 0xbff1f689d7fd5c80d3f93cbce8d1c1add1b8312c5d6ac5b885f469bedee3f6a8.
//
// Solidity: e FileRename(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) FilterFileRename(opts *bind.FilterOpts, entity []common.Address) (*BucketFileRenameIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "FileRename", entityRule)
	if err != nil {
		return nil, err
	}
	return &BucketFileRenameIterator{contract: _Bucket.contract, event: "FileRename", logs: logs, sub: sub}, nil
}

// WatchFileRename is a free log subscription operation binding the contract event 0xbff1f689d7fd5c80d3f93cbce8d1c1add1b8312c5d6ac5b885f469bedee3f6a8.
//
// Solidity: e FileRename(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) WatchFileRename(opts *bind.WatchOpts, sink chan<- *BucketFileRename, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "FileRename", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketFileRename)
				if err := _Bucket.contract.UnpackLog(event, "FileRename", log); err != nil {
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

// BucketNewFileIterator is returned from FilterNewFile and is used to iterate over the raw logs and unpacked data for NewFile events raised by the Bucket contract.
type BucketNewFileIterator struct {
	Event *BucketNewFile // Event containing the contract specifics and raw log

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
func (it *BucketNewFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketNewFile)
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
		it.Event = new(BucketNewFile)
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
func (it *BucketNewFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketNewFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketNewFile represents a NewFile event raised by the Bucket contract.
type BucketNewFile struct {
	Entity common.Address
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewFile is a free log retrieval operation binding the contract event 0x2c618b3515ba4be58b653279244757f1b863bf68901fee1deb8acc4112200a1f.
//
// Solidity: e NewFile(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) FilterNewFile(opts *bind.FilterOpts, entity []common.Address) (*BucketNewFileIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "NewFile", entityRule)
	if err != nil {
		return nil, err
	}
	return &BucketNewFileIterator{contract: _Bucket.contract, event: "NewFile", logs: logs, sub: sub}, nil
}

// WatchNewFile is a free log subscription operation binding the contract event 0x2c618b3515ba4be58b653279244757f1b863bf68901fee1deb8acc4112200a1f.
//
// Solidity: e NewFile(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) WatchNewFile(opts *bind.WatchOpts, sink chan<- *BucketNewFile, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "NewFile", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketNewFile)
				if err := _Bucket.contract.UnpackLog(event, "NewFile", log); err != nil {
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

// BucketNewReadPermissionIterator is returned from FilterNewReadPermission and is used to iterate over the raw logs and unpacked data for NewReadPermission events raised by the Bucket contract.
type BucketNewReadPermissionIterator struct {
	Event *BucketNewReadPermission // Event containing the contract specifics and raw log

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
func (it *BucketNewReadPermissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketNewReadPermission)
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
		it.Event = new(BucketNewReadPermission)
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
func (it *BucketNewReadPermissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketNewReadPermissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketNewReadPermission represents a NewReadPermission event raised by the Bucket contract.
type BucketNewReadPermission struct {
	Entity common.Address
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewReadPermission is a free log retrieval operation binding the contract event 0x04060b000f2d2b8ce099f0ea3f5542353d7069d78f3e8a2eb0e1c50733026b53.
//
// Solidity: e NewReadPermission(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) FilterNewReadPermission(opts *bind.FilterOpts, entity []common.Address) (*BucketNewReadPermissionIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "NewReadPermission", entityRule)
	if err != nil {
		return nil, err
	}
	return &BucketNewReadPermissionIterator{contract: _Bucket.contract, event: "NewReadPermission", logs: logs, sub: sub}, nil
}

// WatchNewReadPermission is a free log subscription operation binding the contract event 0x04060b000f2d2b8ce099f0ea3f5542353d7069d78f3e8a2eb0e1c50733026b53.
//
// Solidity: e NewReadPermission(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) WatchNewReadPermission(opts *bind.WatchOpts, sink chan<- *BucketNewReadPermission, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "NewReadPermission", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketNewReadPermission)
				if err := _Bucket.contract.UnpackLog(event, "NewReadPermission", log); err != nil {
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

// BucketNewWritePermissionIterator is returned from FilterNewWritePermission and is used to iterate over the raw logs and unpacked data for NewWritePermission events raised by the Bucket contract.
type BucketNewWritePermissionIterator struct {
	Event *BucketNewWritePermission // Event containing the contract specifics and raw log

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
func (it *BucketNewWritePermissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketNewWritePermission)
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
		it.Event = new(BucketNewWritePermission)
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
func (it *BucketNewWritePermissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketNewWritePermissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketNewWritePermission represents a NewWritePermission event raised by the Bucket contract.
type BucketNewWritePermission struct {
	Entity common.Address
	FileId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewWritePermission is a free log retrieval operation binding the contract event 0x9a6884615d185a700778b977ce626f0444ca77e3537754baded1de887d06a2be.
//
// Solidity: e NewWritePermission(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) FilterNewWritePermission(opts *bind.FilterOpts, entity []common.Address) (*BucketNewWritePermissionIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "NewWritePermission", entityRule)
	if err != nil {
		return nil, err
	}
	return &BucketNewWritePermissionIterator{contract: _Bucket.contract, event: "NewWritePermission", logs: logs, sub: sub}, nil
}

// WatchNewWritePermission is a free log subscription operation binding the contract event 0x9a6884615d185a700778b977ce626f0444ca77e3537754baded1de887d06a2be.
//
// Solidity: e NewWritePermission(entity indexed address, fileId uint256)
func (_Bucket *BucketFilterer) WatchNewWritePermission(opts *bind.WatchOpts, sink chan<- *BucketNewWritePermission, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "NewWritePermission", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketNewWritePermission)
				if err := _Bucket.contract.UnpackLog(event, "NewWritePermission", log); err != nil {
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

// BucketOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Bucket contract.
type BucketOwnershipRenouncedIterator struct {
	Event *BucketOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *BucketOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketOwnershipRenounced)
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
		it.Event = new(BucketOwnershipRenounced)
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
func (it *BucketOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketOwnershipRenounced represents a OwnershipRenounced event raised by the Bucket contract.
type BucketOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Bucket *BucketFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*BucketOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BucketOwnershipRenouncedIterator{contract: _Bucket.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Bucket *BucketFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *BucketOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketOwnershipRenounced)
				if err := _Bucket.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// BucketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bucket contract.
type BucketOwnershipTransferredIterator struct {
	Event *BucketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BucketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BucketOwnershipTransferred)
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
		it.Event = new(BucketOwnershipTransferred)
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
func (it *BucketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BucketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BucketOwnershipTransferred represents a OwnershipTransferred event raised by the Bucket contract.
type BucketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Bucket *BucketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BucketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bucket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BucketOwnershipTransferredIterator{contract: _Bucket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Bucket *BucketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BucketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bucket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BucketOwnershipTransferred)
				if err := _Bucket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610325806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610192565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a03600435166101a1565b600054600160a060020a0316331461013d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f73656e646572206973206e6f7420636f6e7472616374206f776e657200000000604482015290519081900360640190fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461021a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f73656e646572206973206e6f7420636f6e7472616374206f776e657200000000604482015290519081900360640190fd5b600160a060020a038116151561029157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f63616e6e6f74206d616b6520302061646472657373206173206f776e65720000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820ddc6cd0d1eb512489336c264404d2d68facfdaf34bdfac8a579e9a906b51cb1d0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
