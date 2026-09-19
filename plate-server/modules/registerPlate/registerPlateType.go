package registerplate

import "time"

type RegisterCode struct {
	RegisterCode string  `json:"registerCode"`
	RegisterCity string  `json:"registerCity"`
	CodePosition *int    `json:"codePosition"`
	Note         *string `json:"note"`
}

type RegisterCodeResponse struct {
	RegisterCode  string     `json:"registerCode"`
	RegisterCity  string     `json:"registerCity"`
	RegisterNote  *string    `json:"registerNote"`
	RegisterAdded *time.Time `json:"registerAdded"`
}
