package app

import "blockchain/pkg/blockchain"

// InitBlockchain inits blockchain.
func initBlockchain(a *App) {
	a.blockchain = blockchain.New()
}
