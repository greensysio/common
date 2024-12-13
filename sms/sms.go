package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/greensysio/common/log"
	"github.com/nyaruka/phonenumbers"
)

func init() {
	log.InitLogger(false)
}

type (
	// SMSI is interface for sending SMS
	SMSI interface {
		SendVerifyCode(VerifyCodeInfo) bool
		SendSMS(r SMSRequest) bool
	}

	// SMS is struct for sending SMS
	SMS struct{}
)

// SendVerifyCode : Using API from 3rd to send sms
func (s *SMS) SendVerifyCode(info VerifyCodeInfo) bool {
	if info.VerifyCode == "" {
		return false
	}
	return s.SendSMS(SMSRequest{
		CountryCode: info.CountryCode,
		Phone:       info.Phone,
		Content:     fmt.Sprintf("%s là mã xác minh GreenSys của bạn.", info.VerifyCode),
	})
}

// SendSMS : Send SMS to user
func (s *SMS) SendSMS(r SMSRequest) bool {
	if r.CountryCode == "" {
		r.CountryCode = "VN"
	}
	if len(r.Content) == 0 || len(r.Phone) == 0 {
		log.Error("SendSMS() invalid request. Missing info, data=", r)
		return false
	}

	// CountryCode must be Upper when using phonenumbers lib
	parsedPhonenumber, err := phonenumbers.Parse(r.Phone, strings.ToUpper(r.CountryCode))
	if err != nil {
		args := map[string]interface{}{
			"phone":        "0772662222",
			"country_code": r.CountryCode,
		}
		log.Error("Error when parsing phone for sending sms! Args: %+v", args)
		return false
	}

	realPhone := phonenumbers.Format(parsedPhonenumber, phonenumbers.NATIONAL)

	data := Payload{
		To:      []string{realPhone},
		Content: r.Content,
		SmsType: "3",           // Loại tin nhắn là Brandname
		Sender:  "GREENSYS.IO", // Brandname đăng ký với speedsms
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Error("SMS Error at payloadBytes: ", err)
		return false
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://api.speedsms.vn/index.php/sms/send", body)
	if err != nil {
		log.Error("SMS Error at Request: ", err)
		return false
	}
	req.SetBasicAuth("4_r50fYyVQTu0NuDICOSpgrblYj3yEqx", "x")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	log.Info("Sending sms to phone ", r.Phone)
	if err != nil {
		log.Error("SMS Error at Response: ", err)
		return false
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataRes = new(DataRes)
	if err := json.Unmarshal(bodyBytes, &dataRes); err != nil {
		log.Error("Cannot Unmarshal bodyBytes in SendSMS() func.", err)
	}
	if dataRes.Code == "00" {
		log.Info("Sent sms to phone ", r.Phone)
	} else {
		log.Error(dataRes.Message)
	}
	return true
}
