package checkplate

import "github.com/google/uuid"

type DataCode struct {
	BasePlateColor       string  `json:"basePlateColor"`
	TextPlateColor       string  `json:"textPlateColor"`
	AdditionalPlateColor *string `json:"additionalPlateColor"`
	RegionCode           string  `json:"regionCode"`
	RegisterFirstCode    string  `json:"registerFirstCode"`
	RegisterLastCode     string  `json:"registerLastCode"`
	RegisterCode         string  `json:"registerCode"`
}

type DataVehicleStatus struct {
	BasePlateColor       string  `json:"basePlateColor"`
	TextPlateColor       string  `json:"textPlateColor"`
	AdditionalPlateColor *string `json:"additionalPlateColor"`
}

type VehicleRegionQuery struct {
	IdRegionCode uuid.UUID `gorm:"column:id_region_code" json:"id_region_code"`
	RegionCode   string    `gorm:"column:region_code" json:"region_code"`
	RegionArea   string    `gorm:"column:region_area" json:"region_area"`
	Note         *string   `gorm:"column:note" json:"note"`
}

type VehicleRegionResponse struct {
	RegionCode string  `json:"regionCode"`
	RegionArea string  `json:"regionArea"`
	Note       *string `json:"note"`
}

type VehicleRegisterParam struct {
	RegisterCode      string `json:"registerCode"`
	RegisterFirstCode string `json:"registerFirstCode"`
	RegisterLastCode  string `json:"registerLastCode"`
}

type VehicleRegister struct {
	RegisterCode string  `gorm:"register_code" json:"registerCode"`
	RegisterCity string  `gorm:"register_city" json:"registerCity"`
	Note         *string `gorm:"column:note" json:"note"`
}

type VehicleStatus struct {
	VehicleType   string `gorm:"column:vehicle_type" json:"vehicleType"`
	VehicleEngine string `gorm:"column:vehicle_engine" json:"vehicleEngine"`
}
