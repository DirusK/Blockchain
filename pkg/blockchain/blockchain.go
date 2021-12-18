package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// New - blockchain constructor.
func New() *Blockchain {
	bc := new(Blockchain)
	bc.NewBlock(100, "1")

	return bc
}

func (bc *Blockchain) RegisterNode(address string) {
	parsedURL, err := url.Parse(address)
	if err != nil {
		log.Println(err)
		return
	}

	host := parsedURL.Host
	if host == "" {
		return
	}

	for _, node := range bc.Nodes {
		if node == host {
			return
		}
	}

	bc.Nodes = append(bc.Nodes, host)
}

func (bc *Blockchain) ValidChain(chain []Block) bool {
	lastBlock := chain[0]
	currentIndex := 1

	for currentIndex < len(chain) {

		block := chain[currentIndex]
		if block.PreviousHash != Hash(lastBlock) {
			return false
		}
		if !ValidProof(lastBlock.Proof, block.Proof, lastBlock.PreviousHash) {
			return false
		}

		lastBlock = block
		currentIndex++
	}

	return true
}

func (bc *Blockchain) ResolveConflicts() bool {
	var newChain []Block
	maxLength := len(bc.Chain)

	for _, node := range bc.Nodes {
		url := fmt.Sprintf("http://%s/v1/blockchain/explore", node)
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
			return false
		}

		defer res.Body.Close()

		byteArr, _ := ioutil.ReadAll(res.Body)

		var response chainResponse
		if err = json.Unmarshal(byteArr, &response); err != nil {
			log.Println(err)
			return false
		}

		if response.Length > maxLength && bc.ValidChain(response.Chain) {
			maxLength = response.Length
			newChain = response.Chain
		}
	}

	if len(newChain) > 0 {
		bc.Chain = newChain
		return true
	}

	return false
}

func (bc *Blockchain) NewBlock(proof uint64, previousHash string) Block {
	block := Block{
		Index:        uint64(len(bc.Chain) + 1),
		Timestamp:    time.Now().UTC(),
		Transactions: bc.CurrentTransactions,
		Proof:        proof,
		PreviousHash: previousHash,
	}

	bc.CurrentTransactions = nil
	bc.Chain = append(bc.Chain, block)

	return block
}

func (bc *Blockchain) NewTransaction(sender, recipient string, amount uint64) uint64 {
	transaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	bc.CurrentTransactions = append(bc.CurrentTransactions, transaction)

	return bc.LastBlock().Index + 1
}

func (bc *Blockchain) ProofOfWork(lastBlock Block) uint64 {
	var (
		lastProof = lastBlock.Proof
		lastHash  = Hash(lastBlock)
		proof     = uint64(0)
	)

	for ValidProof(lastProof, proof, lastHash) == false {
		proof++
	}

	return proof
}

func ValidProof(lastProof, proof uint64, lastHash string) bool {
	guess := fmt.Sprintf("%x%x%x", lastProof, proof, lastHash)
	guessBytes := sha256.Sum256([]byte(guess))
	guessHash := hex.EncodeToString(guessBytes[:])

	return guessHash[:4] == "0000"
}

func (bc *Blockchain) LastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func Hash(block Block) string {
	jsonBytes, _ := json.Marshal(block)
	hashBytes := sha256.Sum256(jsonBytes)

	return hex.EncodeToString(hashBytes[:])
}
