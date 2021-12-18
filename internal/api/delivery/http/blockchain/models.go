package blockchain

import "blockchain/pkg/blockchain"

type (
	// TransactionRequest request model.
	TransactionRequest struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		Amount    uint64 `json:"amount"`
	}

	// TransactionResponse response model.
	TransactionResponse struct {
		Message string `json:"message"`
	}

	// ExploreResponse response model.
	ExploreResponse struct {
		Chain  []blockchain.Block `json:"chain"`
		Length int                `json:"length"`
	}

	// NodesRegisterRequest request model.
	NodesRegisterRequest struct {
		Nodes []string `json:"nodes"`
	}

	// NodesRegisterResponse response model.
	NodesRegisterResponse struct {
		Message    string   `json:"message"`
		TotalNodes []string `json:"total_nodes"`
	}

	// NodesResolveResponse response model.
	NodesResolveResponse struct {
		Message string             `json:"message"`
		Chain   []blockchain.Block `json:"chain"`
	}
)
