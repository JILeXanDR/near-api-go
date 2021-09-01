package rpc

import (
	"math/big"

	"github.com/near/borsh-go"
)

type PublicKeySchema struct {
	KeyType uint8
	Data    [32]byte
}

type ActionSchema struct {
	Enum               borsh.Enum         `borsh_enum:"true"`
	CreateAccount      CreateAccountCall  // 0
	DeployContract     DeployContractCall // 1
	ActionFunctionCall ActionFunctionCall // 2
}

type CreateAccountCall struct {
}

type DeployContractCall struct {
}

type ActionFunctionCall struct {
	MethodName string
	Args       []uint8
	Gas        uint64
	Deposit    big.Int
}

type TransactionSchema struct {
	SignerID   string
	PublicKey  PublicKeySchema
	Nonce      uint64
	ReceiverID string
	BlockHash  [32]byte
	Actions    []ActionSchema
}

type SignedTransactionSchema struct {
	Transaction TransactionSchema
	Signature   SignatureSchema
}

type SignatureSchema struct {
	KeyType uint8
	Data    [64]byte
}
