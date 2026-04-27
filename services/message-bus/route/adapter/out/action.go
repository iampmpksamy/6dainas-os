package out

import (
	"time"

	"github.com/6dainas/6dainas-os-Common/utils"
	"github.com/6dainas/6dainas-message-bus/codegen"
	"github.com/6dainas/6dainas-message-bus/model"
)

func ActionAdapter(action model.Action) codegen.Action {
	return codegen.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  utils.Ptr(time.Unix(action.Timestamp, 0)),
	}
}
