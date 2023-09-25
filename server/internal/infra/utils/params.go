package utils

type Param[Object any] func(*Object)

func ApplyParams[Object any](object *Object, params []Param[Object]) *Object {
	for _, param := range params {
		param(object)
	}
	return object
}
