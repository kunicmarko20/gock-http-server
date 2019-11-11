package propertyaccess

import "errors"

type PropertyAccessor struct{}

func NewPropertyAccessor() *PropertyAccessor {
	return &PropertyAccessor{}
}

func (p PropertyAccessor) GetValue(data map[string]interface{}, property *Property) (interface{}, error) {
	paths := property.Path()
	pathsLength := len(paths)

	if pathsLength == 1 {
		return data[paths[0]], nil
	}

	for index, path := range property.Path() {
		if index == (pathsLength - 1) {
			return data[path], nil
		}

		data = data[path].(map[string]interface{})
	}

	return nil, errors.New("value not found")
}
