package blockchain

import "time"

type (
	// Blockchain model.
	Blockchain struct {
		Chain               []Block       `json:"chain"`
		CurrentTransactions []Transaction `json:"current_transactions"`
		Nodes               []string      `json:"nodes"`
	}

	// Block model.
	Block struct {
		Index        uint64        `json:"index"`
		Timestamp    time.Time     `json:"timestamp"`
		Transactions []Transaction `json:"transactions"`
		Proof        uint64        `json:"proof"`
		PreviousHash string        `json:"previous_hash"`
	}

	// Transaction model.
	Transaction struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		Amount    uint64 `json:"amount"`
	}

	// Mine model.
	Mine struct {
		Message      string        `json:"message"`
		Index        uint64        `json:"index"`
		Transactions []Transaction `json:"transactions"`
		Proof        uint64        `json:"proof"`
		PreviousHash string        `json:"previous_hash"`
	}

	// chainResponse model.
	chainResponse struct {
		Length int     `json:"length"`
		Chain  []Block `json:"chain"`
	}
)
