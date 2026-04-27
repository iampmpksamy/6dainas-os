package in

import (
	"github.com/6dainas/6dainas-message-bus/codegen"
	"github.com/6dainas/6dainas-message-bus/model"
)

func PropertyTypeAdapter(propertyType codegen.PropertyType) model.PropertyType {
	return model.PropertyType{
		Name: propertyType.Name,
	}
}
