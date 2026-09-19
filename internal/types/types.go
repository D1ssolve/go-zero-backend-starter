package types

type CreateNoteRequest struct {
	Text string `json:"text"`
}

type Note struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

type ListNotesResponse struct {
	Items []Note `json:"items"`
}

type HealthResponse struct {
	Status string `json:"status"`
}
