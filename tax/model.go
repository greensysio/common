package tax

import (
	"time"
)

type (
	// FuncI is interface for getting information by tax code
	FuncI interface {
		GetInfo(string) CompanyJSON
	}

	// Tax model
	Tax struct {
		ServerName string
	}

	// CompanyJSON for response from 3rd API
	CompanyJSON struct {
		TaxCode              string           `json:"MaSoThue"`
		CreatedDate          *time.Time       `json:"NgayCap"`
		Name                 string           `json:"Title"`
		EnName               string           `json:"TitleEn"`
		Address              string           `json:"DiaChiCongTy"`
		Owner                string           `json:"ChuSoHuu"`
		Director             string           `json:"GiamDoc"`
		RegisteredPlacePhone string           `json:"NoiDangKyQuanLy_DienThoai"`
		Domain               string           `json:"NganhNgheTitle"`
		SubCompanies         []SubCompanyJSON `json:"LtsDoanhNghiepTrucThuoc"`
	}

	// SubCompanyJSON of comapny by taxcode
	SubCompanyJSON struct {
		TaxCode  string `json:"MaSoThue"`
		Relation int    `json:"QuanHe"`
		Name     string `json:"Title"`
		Address  string `json:"DiaChi"`
	}
)
