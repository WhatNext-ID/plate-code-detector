package regionplate

import "time"

type RegionCode struct {
	Code         string  `json:"code"`
	Area         string  `json:"area"`
	Note         *string `json:"note"`
	CodePosition string  `json:"position"`
}

type RegionCodeResponse struct {
	RegionCode  string     `json:"regionCode"`
	RegionArea  string     `json:"regionArea"`
	RegionNote  *string    `json:"regionNote"`
	RegionAdded *time.Time `json:"regionAdded"`
}
