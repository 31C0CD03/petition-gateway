package petitions

type CreatePetitionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreatePetitionResponse struct {
	Created bool `json:"created"`
}
