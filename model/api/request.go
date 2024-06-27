package api

type InsertRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type FindRequest struct {
	Source string `json:"source"`
}

type DeleteRequest struct {
	Source string `json:"source"`
}

type UpdateRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}