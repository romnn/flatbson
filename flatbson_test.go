package flatbson

import (
	"testing"

	"github.com/romnn/deepequal"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFlatten(t *testing.T) {
	input := bson.D{
		{Key: "user", Value: bson.D{
			{Key: "email", Value: "test"},
			{Key: "list", Value: bson.A{"a0", bson.M{"a1": bson.E{Key: "will be", Value: "flattened"}}, "a2"}},
		}},
	}
	expected := bson.D{
		{Key: "user.email", Value: "test"},
		{Key: "user.list", Value: bson.A{"a0", bson.M{"a1": bson.E{Key: "will be", Value: "flattened"}}, "a2"}},
	}
	flattened, err := Flattened(input, ".")
	if err != nil {
		t.Errorf("Failed to flatten: %v", err)
	}
	if equal, err := deepequal.DeepEqual(flattened, expected); err != nil || !equal {
		t.Errorf("\nGot: %v\nWant %v\nError: %v", flattened, expected, err)
	}
}

func TestFlattenWithSlices(t *testing.T) {
	input := bson.D{
		{Key: "user", Value: bson.D{
			{Key: "email", Value: "test"},
			{Key: "list", Value: bson.A{"a0", bson.D{{Key: "a1", Value: bson.E{Key: "will be", Value: "flattened"}}}, "a2"}},
		}},
	}
	expected := bson.D{
		{Key: "user.email", Value: "test"},
		{Key: "user.list.0", Value: "a0"},
		{Key: "user.list.1.a1.will be", Value: "flattened"},
		{Key: "user.list.2", Value: "a2"},
	}
	f := &Flattener{FlattenSlices: true}
	flattened, err := f.Flattened(input)
	if err != nil {
		t.Errorf("Failed to flatten: %v", err)
	}
	if equal, err := deepequal.DeepEqual(flattened, expected); err != nil || !equal {
		t.Errorf("\nGot: %v\nWant %v\nError: %v", flattened, expected, err)
	}
}

func TestSeparators(t *testing.T) {
	input := bson.D{
		{Key: "user", Value: bson.D{{Key: "email", Value: "test"}}},
	}
	tests := map[string]bson.D{
		".": bson.D{{Key: "user.email", Value: "test"}},
		" ": bson.D{{Key: "user email", Value: "test"}},
	}
	for sep, expected := range tests {
		flattened, err := Flattened(input, sep)
		if err != nil {
			t.Errorf("Failed to flatten: %v", err)
		}
		if equal, err := deepequal.DeepEqual(flattened, expected); err != nil || !equal {
			t.Errorf("\nGot: %v\nWant %v\nError: %v", flattened, expected, err)
		}
	}
}
