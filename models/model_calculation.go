package models


// BillRequest is the request to create a calculation
type BillRequest struct {
	Name string `json:"name"`
}

// BillResponse is the response to create a calculation
type BillResponse struct {
	Bill float64 `json:"bill"`
}
