package json

import (
	"encoding/json"
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/pkg/errors"
)

func ParseJson[T any](bs []byte) (*T, error) {
	var t T
	if err := json.Unmarshal(bs, &t); err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal json")
	}
	return &t, nil
}

func ParseJsonFromFile[T any](path string) (*T, error) {
	bytes, err := file.ReadFileBytes(path)
	if err != nil {
		return nil, err
	}
	return ParseJson[T](bytes)
}

func ParseJsonFromString[T any](contents string) (*T, error) {
	return ParseJson[T]([]byte(contents))
}
