package out

import (
	"time"

	"github.com/6dainas/6dainas-os-Common/utils"
	"github.com/6dainas/6dainas-message-bus/codegen"
	"github.com/6dainas/6dainas-message-bus/model"
)

func EventAdapter(event model.Event) codegen.Event {
	// properties := make([]codegen.Property, 0)
	// for _, property := range event.Properties {
	// 	properties = append(properties, PropertyAdapter(property))
	// }

	return codegen.Event{
		SourceID:   event.SourceID,
		Name:       event.Name,
		Properties: event.Properties,
		Timestamp:  utils.Ptr(time.Unix(event.Timestamp, 0)),
		Uuid:       &event.UUID,
	}
}
