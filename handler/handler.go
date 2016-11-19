package handler

import (
	"github.com/suicidegang/onesignal-srv/proto/onesignal"
	"golang.org/x/net/context"
)

const (
	ONESIGNAL_ENDPOINT string = "https://onesignal.com/api/v1/"
)

var (
	ApiKey string
	AppID  string
)

type Onesignal struct{}

func (srv *Onesignal) SendAll(ctx context.Context, req *onesignal.SendAllRequest, res *onesignal.SendAllResponse) error {
	push, err := RequestPost("notifications", map[string]interface{}{
		"included_segments": []string{"All"},
		"contents": map[string]interface{}{
			"en": req.Message,
		},
		"data": req.Variables,
	})

	if err != nil {
		return err
	}

	res.MessageID = "unknown"

	if mid, exists := push["id"]; exists && len(mid.(string)) > 0 {
		res.MessageID = mid.(string)
	}

	return nil
}

func (srv *Onesignal) Send(ctx context.Context, req *onesignal.SendRequest, res *onesignal.SendResponse) error {
	params := map[string]interface{}{
		"contents": map[string]interface{}{
			"en": req.Message,
		},
		"data": req.Variables,
	}

	if len(req.Ids) > 0 {
		params["include_player_ids"] = req.Ids
	}

	if len(req.Segments) > 0 {
		params["included_segments"] = req.Segments
	}

	push, err := RequestPost("notifications", params)

	if err != nil {
		return err
	}

	res.MessageID = "unknown"

	if mid, exists := push["id"]; exists && len(mid.(string)) > 0 {
		res.MessageID = mid.(string)
	}

	return nil
}
