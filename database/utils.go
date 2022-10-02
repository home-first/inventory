package database

type Minimizable interface {
	Minimize() interface{}
}

// Take and array of objects and return a new array of the same type with the same length but with the minimize method called on each object.
func Minimize(objects []Minimizable) []interface{} {
	ret := make([]interface{}, len(objects))
	for i, object := range objects {
		ret[i] = object.Minimize()
	}
	return ret
}
