package blockchain

import (
	"net/http"
	"strconv"

	errors "blockchain/internal/api/delivery/http"
	"blockchain/pkg/blockchain"

	"github.com/gofiber/fiber/v2"
)

type (
	// Handler defines a handler for HTTP requests for working with blockchain.
	Handler struct {
		chain  *blockchain.Blockchain
		nodeID string
	}
)

// NewHandler defines a handler constructor.
func NewHandler(blockchain *blockchain.Blockchain, nodeID string) *Handler {
	return &Handler{
		chain:  blockchain,
		nodeID: nodeID,
	}
}

func (h *Handler) Mine(ctx *fiber.Ctx) error {
	lastBlock := h.chain.LastBlock()
	proof := h.chain.ProofOfWork(lastBlock)

	h.chain.NewTransaction("0", h.nodeID, 1)
	previousHash := blockchain.Hash(lastBlock)
	block := h.chain.NewBlock(proof, previousHash)

	mine := blockchain.Mine{
		Message:      "New Block Forged",
		Index:        block.Index,
		Transactions: block.Transactions,
		Proof:        block.Proof,
		PreviousHash: block.PreviousHash,
	}

	return h.respond(ctx, http.StatusOK, mine)
}

func (h *Handler) CreateTransaction(ctx *fiber.Ctx) error {
	var transaction TransactionRequest

	err := ctx.BodyParser(&transaction)
	if err != nil {
		return h.respond(ctx, http.StatusInternalServerError, errors.ErrResponse{Error: "cannot unmarshal request"})
	}

	index := h.chain.NewTransaction(transaction.Sender, transaction.Recipient, transaction.Amount)
	response := TransactionResponse{
		Message: "Transaction will be added to Block " + strconv.FormatUint(index, 10),
	}

	return h.respond(ctx, http.StatusOK, response)
}

func (h *Handler) Explore(ctx *fiber.Ctx) error {
	responseExplore := ExploreResponse{
		Chain:  h.chain.Chain,
		Length: len(h.chain.Chain),
	}

	return h.respond(ctx, http.StatusOK, responseExplore)
}

func (h *Handler) NodesRegister(ctx *fiber.Ctx) error {
	var nodesRegister NodesRegisterRequest

	err := ctx.BodyParser(&nodesRegister)
	if err != nil {
		return h.respond(ctx, http.StatusInternalServerError, errors.ErrResponse{Error: "cannot unmarshal request"})
	}

	for _, node := range nodesRegister.Nodes {
		h.chain.RegisterNode(node)
	}

	responseNodesRegister := NodesRegisterResponse{
		Message:    "New nodes have been added",
		TotalNodes: h.chain.Nodes,
	}

	return h.respond(ctx, http.StatusOK, responseNodesRegister)
}

func (h *Handler) NodesResolve(ctx *fiber.Ctx) error {
	var resNodesResolve NodesResolveResponse

	replaced := h.chain.ResolveConflicts()
	if replaced {
		resNodesResolve = NodesResolveResponse{
			Message: "Our chain was replaced",
			Chain:   h.chain.Chain,
		}
	} else {
		resNodesResolve = NodesResolveResponse{
			Message: "Our chain is authoritative",
			Chain:   h.chain.Chain,
		}
	}

	return h.respond(ctx, http.StatusOK, resNodesResolve)
}
