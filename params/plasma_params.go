package params

import (
	"math/big"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
)

var (
	NullAddress = common.Address{0x0000000000000000000000000000000000000000}
	NullKey, _  = crypto.HexToECDSA("00")

	// address of b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291
	Operator = common.HexToAddress("0x71562b71999873DB5b286dF957af199Ec94617F7")

	RequestTxGasPrice        = big.NewInt(1e9)
	RequestTxGasLimit uint64 = 100000
)
