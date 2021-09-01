package rpc

import (
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/ebellocchia/go-base58"
	"github.com/near/borsh-go"
)

const (
	callTxFunctionGas = 10_000_000_000_000
)

var base58BTC = base58.New(base58.AlphabetBitcoin)

func (c *Client) CreateAndSignCallFunctionTransaction(
	ctx context.Context,
	privateKey string,
	contractID string,
	signerID string,
	methodName string,
	args map[string]interface{},
) ([]byte, error) {
	privateKeyB58Decoded, err := base58BTC.Decode(strings.Split(privateKey, ":")[1])
	if err != nil {
		return nil, err
	}

	prv := ed25519.PrivateKey(privateKeyB58Decoded)
	publicKey := prv.Public().(ed25519.PublicKey)
	publicKeyData := [32]byte{}
	copy(publicKeyData[:], publicKey)

	accessKey, err := c.Query(ctx, RequestParams{
		"request_type": "view_access_key",
		"finality":     "final",
		"account_id":   signerID,
		"public_key":   base58BTC.Encode(publicKey),
	})
	if err != nil {
		return nil, err
	}

	blockHashData := [32]byte{}
	blockHashBase58Decoded, err := base58BTC.Decode(accessKey.BlockHash)
	if err != nil {
		return nil, err
	} else {
		copy(blockHashData[:], blockHashBase58Decoded)
	}

	argsJSONBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	transaction := TransactionSchema{
		SignerID: signerID,
		PublicKey: PublicKeySchema{
			KeyType: 0,
			Data:    publicKeyData,
		},
		Nonce:      accessKey.Nonce + 1,
		ReceiverID: contractID,
		BlockHash:  blockHashData,
		Actions: []ActionSchema{
			{
				Enum: 2,
				ActionFunctionCall: ActionFunctionCall{
					MethodName: methodName,
					Args:       argsJSONBytes,
					Gas:        callTxFunctionGas,
				},
			},
		},
	}

	serializedTx, err := borsh.Serialize(transaction)
	if err != nil {
		return nil, err
	}
	serializedTxHash := sha256.Sum256(serializedTx)
	signedSerialized := ed25519.Sign(prv, serializedTxHash[:])

	signatureData := [64]byte{}
	copy(signatureData[:], signedSerialized)

	signedTransaction := SignedTransactionSchema{
		Transaction: transaction,
		Signature: SignatureSchema{
			KeyType: 0,
			Data:    signatureData,
		},
	}

	signedSerializedTx, err := borsh.Serialize(signedTransaction)
	if err != nil {
		return nil, err
	}

	return []byte(base64.StdEncoding.EncodeToString(signedSerializedTx)), nil
}
