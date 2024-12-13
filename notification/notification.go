package notification

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	cmContext "github.com/greensysio/common/context"
	"github.com/greensysio/common/log"
	"github.com/greensysio/common/model/res"
)

func init() {
	log.InitLogger(false)
}

// SendNotification send request to  API firebase.exgo.vn
func SendNotification(c *cmContext.CustomContext, pl Payload) bool {
	ctx, cncl := cmContext.InitNewCtxFromCustomCtx(c)
	defer cncl()

	// Set payload sender.
	tokenInfo := c.GetTokenInfo()
	if tokenInfo.UserID != "" {
		pl.SenderId = tokenInfo.UserID
	}
	if len(tokenInfo.GetRoles()) != 0 {
		pl.SenderRole = c.GetTokenInfo().GetRole()
	}

	plByte, _ := json.Marshal(pl)
	req, err := http.NewRequest("POST", os.Getenv("FB_PUSH_NOTIFICATION"), strings.NewReader(string(plByte)))
	if err != nil {
		log.ErrorfCtx(c, "Error when call notification server: %+v", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	log.InfofCtx(c, "Push notification to %s . Args: %+v", req.URL.String(), pl)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.ErrorfCtx(c, "Response from Notification API gets error: %+v", err)
		return false
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	dataRes := new(res.Response)
	if err := json.Unmarshal(bodyBytes, &dataRes); err != nil {
		log.ErrorfCtx(c, "Can not Unmarshal in SendNotification() func. Response: %+v . Error: %+v", string(bodyBytes), err)
		return false
	}
	if dataRes.Status != http.StatusOK {
		log.WarnfCtx(c, "Push notification fail! Respone: %+v", dataRes)
		return false
	}
	return true
}

// SendNotificationForSchedule send request to notify service
func SendNotificationForSchedule(c *cmContext.CustomContext, pl Payload) bool {
	ctx, cncl := cmContext.InitNewCtxFromCustomCtx(c)
	defer cncl()

	plByte, _ := json.Marshal(pl)
	req, err := http.NewRequest("POST", os.Getenv("FB_PUSH_NOTIFICATION")+"/schedule", strings.NewReader(string(plByte)))
	if err != nil {
		log.ErrorfCtx(c, "Error when call notification of schedule to %s . Error: %+v", req.URL.String(), err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	log.InfofCtx(c, "Push notification of schedule to %s . Args: %+v", req.URL.String(), pl)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.ErrorfCtx(c, "Response from Schedule Notification API gets error: %+v", err)
		return false
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	dataRes := new(res.Response)
	if err := json.Unmarshal(bodyBytes, &dataRes); err != nil {
		log.ErrorfCtx(c, "Can not Unmarshal in SendNotification() func. Error: %+v", err)
		return false
	}
	if dataRes.Status != http.StatusOK {
		log.WarnfCtx(c, "Push notification Schedule fail! Respone: %+v", dataRes)
		return false
	}
	return true
}
