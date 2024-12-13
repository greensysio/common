package tax

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/greensysio/common/log"
)

func init() {
	log.InitLogger(false)
}

// GetInfo gets information of company from API
func (t Tax) GetInfo(TaxCode string) CompanyJSON {
	url := fmt.Sprintf(t.ServerName, TaxCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error("Error when call: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	log.Info("Getting company information by tax code: ", TaxCode)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Response for API gets error: ", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataRes = new(CompanyJSON)
	if err := json.Unmarshal(bodyBytes, &dataRes); err != nil {
		log.Error("Can not Unmarshal in GetInfo() func.", err)
	}
	if dataRes.TaxCode != TaxCode {
		log.Warn("Can not get information of TaxCode: ", TaxCode)
	}

	return CompanyJSON{
		TaxCode:              dataRes.TaxCode,
		CreatedDate:          dataRes.CreatedDate,
		Name:                 dataRes.Name,
		EnName:               dataRes.EnName,
		Address:              dataRes.Address,
		Owner:                dataRes.Owner,
		Director:             dataRes.Director,
		RegisteredPlacePhone: dataRes.RegisteredPlacePhone,
		Domain:               dataRes.Domain,
		SubCompanies:         dataRes.SubCompanies,
	}
}
