package rpc

// response represents RPC response common for all methods.
type response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  interface{} `json:"result"`
	Error   *RPCError   `json:"error"`
}

type QueryResult struct {
	BlockHeight int      `json:"block_height"`
	BlockHash   string   `json:"block_hash"`
	Logs        []string `json:"logs,omitempty"`
	Result      []byte   `json:"result"`
	Nonce       uint64   `json:"nonce,omitempty"`
	Permission  string   `json:"permission,omitempty"`
}

// NetworkInfoResult is used for RPC method "network_info".
type NetworkInfoResult struct {
	ActivePeers []struct {
		ID        string  `json:"id"`
		Addr      string  `json:"addr"`
		AccountID *string `json:"account_id"`
	} `json:"active_peers"`
	NumActivePeers      int `json:"num_active_peers"`
	PeerMaxCount        int `json:"peer_max_count"`
	SentBytesPerSec     int `json:"sent_bytes_per_sec"`
	ReceivedBytesPerSec int `json:"received_bytes_per_sec"`
	KnownProducers      []struct {
		AccountID string      `json:"account_id"`
		Addr      interface{} `json:"addr"`
		PeerID    string      `json:"peer_id"`
	} `json:"known_producers"`
}

// CommitBroadcastTx is used for RPC method "broadcast_tx_commit".
type CommitBroadcastTx struct {
	Status struct {
		SuccessValue string `json:"SuccessValue"`
	} `json:"status"`
	Transaction struct {
		Hash string `json:"hash"`
	} `json:"transaction"`
}

// AsyncBroadcastTxResult is used for RPC method "broadcast_tx_async".
type AsyncBroadcastTxResult string

type ViewAccountResult struct {
	Amount        string `json:"amount"`
	BlockHash     string `json:"block_hash"`
	BlockHeight   int    `json:"block_height"`
	CodeHash      string `json:"code_hash"`
	Locked        string `json:"locked"`
	StoragePaidAt int    `json:"storage_paid_at"`
	StorageUsage  int    `json:"storage_usage"`
}
