package rpc

import "fmt"

// RPCError contains error details.
// Docs https://docs.near.org/docs/api/rpc/contracts#what-could-go-wrong
type RPCError struct {
	Name  string `json:"name"`
	Cause struct {
		Name string                 `json:"name"`
		Info map[string]interface{} `json:"info"`
	} `json:"cause"`
	Code    float64 `json:"code"`
	Message string  `json:"message"`
	Data    string  `json:"data"`
}

func (e RPCError) Error() string {
	return fmt.Sprintf("RPC error: name %s, code %d, message %s, data %s", e.Name, int(e.Code), e.Message, e.Data)
}

func IsRPCError(err error) bool {
	_, ok := err.(*RPCError)
	return ok
}
