package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

var isUseNameTag = false

const (
	tagName = "name"
)

type CustomMarshaller struct {
	Value interface{}
}

func (c CustomMarshaller) MarshalWithNameTag() ([]byte, error) {
	isUseNameTag = true
	data, err := c.MarshalJSON()
	isUseNameTag = false

	return data, err
}

func (c CustomMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	if isUseNameTag {
		fmt.Println("Is using name tag")
		var intf interface{}
		if err = json.Unmarshal(marshalled, &intf); err != nil {
			goto end
		}

		m := intf.(map[string]interface{})
		_m := make(map[string]interface{}, len(m))

		typ := reflect.TypeOf(c.Value)

		idx := 0
		for _, v := range m {
			tag, found := typ.Field(idx).Tag.Lookup(tagName)
			if !found {
				continue
			}

			tag = strings.ReplaceAll(tag, " ", "_")
			tag = strings.ReplaceAll(tag, "-", "")
			tag = strings.ToLower(tag)

			_m[tag] = v
			idx++
		}

		marshalled, err = json.Marshal(_m)
	}

end:
	return marshalled, err
}
