package models

type Code struct {
	KodeWilayah     string  `json:"kode_wilayah"`
	KodeAwal        string  `json:"kode_awal"`
	KodeAkhir       string  `json:"kode_akhir"`
	KodeAliasAwal   string  `json:"kode_alias_awal"`
	KodeAliasAkhir  string  `json:"kode_alias_akhir"`
	KodeRegistrasi  string  `json:"kode_registrasi"`
	WilayahHukum    *string `json:"wilayah_hukum"`
	StatusKendaraan string  `json:"status_kendaraan"`
}

type KodeRegisterPost struct {
	KodeAwal        *string `json:"kode_awal"`
	KodeAkhir       *string `json:"kode_akhir"`
	KodeAlias       *string `json:"kode_alias"`
	WilayahHukum    *string `json:"wilayah_hukum"`
	Keterangan      string  `json:"keterangan"`
	KodeWilayah     string  `json:"kode_wilayah"`
	StatusKendaraan string  `json:"status_kendaraan"`
}

type KodeRegisterKhususPost struct {
	KodeRegistrasi  string  `json:"kode_registrasi"`
	WilayahHukum    *string `json:"wilayah_hukum"`
	Keterangan      string  `json:"keterangan"`
	KodeWilayah     string  `json:"kode_wilayah"`
	StatusKendaraan string  `json:"status_kendaraan"`
}

type KendaraanKhususResult struct {
	KodeRegistrasi  string `json:"kode_registrasi"`
	WilayahHukum    string `json:"wilayah_hukum"`
	StatusKendaraan string `json:"status_kendaraan"`
	Keterangan      string `json:"keterangan"`
	NamaProvinsi    string `json:"nama_provinsi"`
}

type KendaraanPribadiResult struct {
	WilayahHukum    string `json:"wilayah_hukum"`
	KodeAwal        string `json:"kode_awal"`
	KodeAkhir       string `json:"kode_akhir"`
	KodeAlias       string `json:"kode_alias"`
	NamaProvinsi    string `json:"nama_provinsi"`
	StatusKendaraan string `json:"status_kendaraan"`
	Keterangan      string `json:"keterangan"`
}
