package propertyaccess

import "regexp"

type Property struct {
	path []string
}

func NewProperty(propertyPath string) *Property {
	regex, _ := regexp.Compile("[\\w+]+")

	return &Property{
		path: regex.FindAllString(propertyPath, -1),
	}
}

func (p Property) Path() []string {
	return p.path
}
