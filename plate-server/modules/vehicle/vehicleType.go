package vehicle

import (
	"plate-server/utils"
	"time"

	"github.com/google/uuid"
)

type VehicleType struct {
	VehicleType string `json:"vehicleType"`
}

type VehicleEngine struct {
	EngineType string `json:"engineType"`
}

type VehicleCategory struct {
	VehicleType   string   `json:"vehicleType"`
	VehicleEngine string   `json:"vehicleEngine"`
	ColorCriteria []string `json:"colorCriteria"`
}

type VehicleData struct {
	IdVehicleCategory uuid.UUID                 `json:"idVehicleCat"`
	VehicleType       string                    `json:"vehicleType"`
	VehicleEngine     string                    `json:"vehicleEngine"`
	ColorCriteria     utils.StringArrayResponse `json:"vehicleColorCriteria"`
	CreatedAt         *time.Time                `json:"createdAt"`
}

type VehicleResponse struct {
	VehicleType          string     `json:"vehicleType"`
	VehicleEngine        string     `json:"vehicleEngine"`
	VehicleColorCriteria string     `json:"vehicleColorCriteria"`
	CreatedAt            *time.Time `json:"createdAt"`
}
