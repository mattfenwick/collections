package json

import (
	"encoding/json"
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/pkg/errors"
)

// SortOptions is used to get keys from a struct into a sorted order.
//   See https://stackoverflow.com/a/61887446/894284.
//   Apparently, golang's json library sorts keys from
//   maps, but NOT from structs.  So this function works by reading json into a
//   generic structure of maps, then marshaling back into completely sorted json.
//   NOTE: this only works for a json object.
func SortOptions(contents []byte, escapeHtml bool, indent bool) ([]byte, error) {
	obj, err := ParseJson[map[string]interface{}](contents)
	if err != nil {
		return nil, err
	}
	return MarshalHelper(obj, escapeHtml, indent)
}

func SortFileOptions(path string, escapeHtml bool, indent bool) error {
	bytes, err := file.ReadFileBytes(path)
	if err != nil {
		return err
	}
	sortedBytes, err := SortOptions(bytes, escapeHtml, indent)
	if err != nil {
		return err
	}
	return file.WriteFileBytes(path, sortedBytes, 0644)
}

// Remarshal is of questionable utility.  It first marshals, then unmarshals
//   into an `interface{}`, which throws away struct types (struct types cause
//   problems for predictability of order during marshaling).
func Remarshal(obj interface{}) (interface{}, error) {
	bs, err := MarshalHelper(obj, false, false)
	if err != nil {
		return nil, err
	}

	var out interface{}
	err = json.Unmarshal(bs, &out)
	return &out, errors.Wrapf(err, "unable to unmarshal json")
}
