package status

type (
	response struct {
		App     string `json:"app"`
		Version string `json:"version"`
		NodeID  string `json:"node_id"`
		Status  int    `json:"status"`
	}
)
