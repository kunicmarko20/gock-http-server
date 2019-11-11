package rules

import (
	"encoding/json"
	"github.com/kunicmarko20/gock-http-server/propertyaccess"
	"log"
	"net/http"
)

type JsonRequestPayloadPropertyEquals struct {
	property         *propertyaccess.Property
	value            interface{}
	propertyAccessor *propertyaccess.PropertyAccessor
}

func NewJsonRequestPayloadPropertyEquals(
	property string,
	value interface{},
	propertyAccessor *propertyaccess.PropertyAccessor,
) *JsonRequestPayloadPropertyEquals {

	return &JsonRequestPayloadPropertyEquals{propertyaccess.NewProperty(property), value, propertyAccessor}
}

func (j *JsonRequestPayloadPropertyEquals) Matches(request *http.Request) bool {
	decoder := json.NewDecoder(request.Body)

	var data map[string]interface{}

	err := decoder.Decode(&data)

	if err != nil {
		log.Println("Error decoding request body in \"JsonRequestPayloadPropertyEquals\".")

		return false
	}

	value, err := j.propertyAccessor.GetValue(data, j.property)

	if err != nil {
		log.Println("Error accessing value of body in \"JsonRequestPayloadPropertyEquals\".")

		return false
	}

	return j.value == value
}
