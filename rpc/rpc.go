package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	server string
	debug  bool
	client *http.Client
}

// New returns NEAR RPC client.
// Related docs https://docs.near.org/docs/api/rpc.
func New(server string, debug bool) *Client {
	return &Client{
		server: server,
		debug:  debug,
		client: &http.Client{},
	}
}

func (c Client) sendRPC(ctx context.Context, method string, params interface{}) (*http.Response, error) {
	id := "imcare-" + strconv.Itoa(int(time.Now().UnixNano()))

	var req request

	req = request{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, fmt.Errorf("encoding into JSON: %w", err)
	}

	// if c.debug {
	// 	started := time.Now().UTC()
	// 	b, _ := json.Marshal(params)
	// 	log.Printf("[NEAR RPC] calling method: %s(%s) id=%s", method, string(b), id)
	// 	defer func() {
	// 		elapsed := time.Now().UTC().Sub(started)
	// 		if data.Result != nil {
	// 			b, _ := json.Marshal(data.Result)
	// 			log.Printf("[NEAR RPC] finished with result: %s in %s, id=%s", string(b), elapsed.String(), data.ID)
	// 		} else if data.Error != nil {
	// 			b, _ := json.Marshal(data.Error)
	// 			log.Printf("[NEAR RPC] finished with error: %s in %s, id=%s", string(b), elapsed.String(), data.ID)
	// 		}
	// 	}()
	// }

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.server, &body)
	if err != nil {
		return nil, fmt.Errorf("constructing request: %w", err)
	}

	httpReq.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("response plain body: %s", string(buf))

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	return resp, nil
}
