package rpc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func (c *Client) Query(ctx context.Context, params interface{}) (*QueryResult, error) {
	resp, err := c.sendRPC(ctx, "query", params)
	if err != nil {
		return nil, fmt.Errorf("calling query RPC method with params %+v: %w", params, err)
	}

	data := response{
		Result: &QueryResult{},
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Error != nil {
		return nil, data.Error
	}

	return data.Result.(*QueryResult), nil
}

// CallFunction is view call.
func (c *Client) CallFunction(ctx context.Context, accountID string, methodName string, args CallFunctionArgs) (*QueryResult, error) {
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	result, err := c.Query(ctx, RequestParams{
		"request_type": "call_function",
		"finality":     "final",
		"account_id":   accountID,
		"method_name":  methodName,
		"args_base64":  base64.StdEncoding.EncodeToString(jsonBytes),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) ViewAccount(ctx context.Context, accountID string) (*QueryResult, error) {
	result, err := c.Query(ctx, RequestParams{
		"request_type": "view_account",
		"finality":     "final",
		"account_id":   accountID,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) BroadcastTxAsync(ctx context.Context, signedTransaction string) (*AsyncBroadcastTxResult, error) {
	params := []string{
		signedTransaction,
	}

	resp, err := c.sendRPC(ctx, "broadcast_tx_async", params)
	if err != nil {
		return nil, fmt.Errorf("calling broadcast_tx_async RPC method with params %+v: %w", signedTransaction, err)
	}

	data := response{
		Result: AsyncBroadcastTxResult(""),
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Error != nil {
		return nil, data.Error
	}

	return data.Result.(*AsyncBroadcastTxResult), nil
}

func (c *Client) BroadcastTxCommit(ctx context.Context, signedTransaction []byte) (*CommitBroadcastTx, error) {
	resp, err := c.sendRPC(ctx, "broadcast_tx_commit", []string{string(signedTransaction)})
	if err != nil {
		return nil, fmt.Errorf("calling broadcast_tx_commit RPC method: %w", err)
	}

	result := &CommitBroadcastTx{}

	data := response{
		Result: result,
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Error != nil {
		return nil, data.Error
	}

	return result, nil
}
