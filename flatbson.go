package flatbson

import (
	"fmt"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// Version is incremented using bump2version
const Version = "0.1.0"

var defaultSeparator = "."

// Flattener represents flattening options
type Flattener struct {
	FlattenSlices bool
	Separator     *string
	flattened     *bson.D
}

func (f *Flattener) flatten(input interface{}, lkey string) error {
	var err error
	var inputMap bson.D
	switch input.(type) {
	case bson.D:
		inputMap = input.(bson.D)
	case bson.E:
		inputMap = bson.D{input.(bson.E)}
	case bson.M:
		for k, v := range input.(bson.M) {
			inputMap = append(inputMap, bson.E{Key: k, Value: v})
		}
	default:
		return fmt.Errorf("Cannot flatten %v of type %v", input, reflect.TypeOf(input))
	}
	for i := 0; i < len(inputMap); i++ {
		// for rkey, value := range inputMap {
		rkey := inputMap[i].Key
		value := inputMap[i].Value
		key := lkey + rkey
		// Check for nested document
		if nested, ok := value.(bson.D); ok {
			err = f.flatten(nested, key+*f.Separator)
		} else if nested, ok := value.(bson.M); ok {
			err = f.flatten(nested, key+*f.Separator)
		} else if nested, ok := value.(bson.E); ok {
			err = f.flatten(nested, key+*f.Separator)
		} else if nested, ok := value.(bson.A); ok && f.FlattenSlices {
			// Convert slice to map using indices as keys
			mapped := bson.D{}
			for i, v := range nested {
				mapped = append(mapped, bson.E{Key: strconv.Itoa(i), Value: v})
			}
			err = f.flatten(mapped, key+*f.Separator)
		} else {
			// Add the value
			*f.flattened = append(*f.flattened, bson.E{Key: key, Value: value})
		}
	}
	return err
}

// Flatten flattens a bson D or M document
func (f *Flattener) Flatten(input interface{}, flattened *bson.D) error {
	f.flattened = flattened
	if f.Separator == nil {
		f.Separator = &defaultSeparator
	}
	return f.flatten(input, "")
}

// Flattened returns a new, flattened bson.D from a bson D or M document
func (f *Flattener) Flattened(input interface{}) (bson.D, error) {
	var flattened bson.D
	err := f.Flatten(input, &flattened)
	return flattened, err
}

// Flatten flattens a bson D or M document
func Flatten(input interface{}, flattened *bson.D, separator string) error {
	f := &Flattener{
		Separator: &separator,
	}
	return f.Flatten(input, flattened)
}

// Flattened returns a new, flattened bson.D from a bson D or M document
func Flattened(input interface{}, separator string) (bson.D, error) {
	var flattened bson.D
	err := Flatten(input, &flattened, separator)
	return flattened, err
}
