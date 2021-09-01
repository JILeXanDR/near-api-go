package rpc

// request represents RPC request common for all methods.
type request struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type RequestParams map[string]interface{}

type CallFunctionArgs map[string]interface{}
