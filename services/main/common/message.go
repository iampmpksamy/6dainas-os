package common

import (
	"github.com/6dainas/6dainas-os/codegen/message_bus"
)

// devtype -> action -> event
var EventTypes = []message_bus.EventType{
	{Name: "6dainas:system:utilization", SourceID: SERVICENAME, PropertyTypeList: []message_bus.PropertyType{}},
	{Name: "6dainas:file:recover", SourceID: SERVICENAME, PropertyTypeList: []message_bus.PropertyType{}},
	{Name: "6dainas:file:operate", SourceID: SERVICENAME, PropertyTypeList: []message_bus.PropertyType{}},
}
