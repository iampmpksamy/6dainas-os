package in

import (
	"github.com/6dainas/6dainas-message-bus/codegen"
	"github.com/6dainas/6dainas-message-bus/model"
)

func EventTypeAdapter(eventType codegen.EventType) model.EventType {
	propertyTypeList := make([]model.PropertyType, 0)
	for _, propertyType := range eventType.PropertyTypeList {
		propertyTypeList = append(propertyTypeList, PropertyTypeAdapter(propertyType))
	}

	return model.EventType{
		SourceID:         eventType.SourceID,
		Name:             eventType.Name,
		PropertyTypeList: propertyTypeList,
	}
}
