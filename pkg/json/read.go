package json

import (
	"encoding/json"
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/pkg/errors"
)

func Parse[T any](bs []byte) (*T, error) {
	var t T
	if err := json.Unmarshal(bs, &t); err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal json")
	}
	return &t, nil
}

func ParseFile[T any](path string) (*T, error) {
	bytes, err := file.Read(path)
	if err != nil {
		return nil, err
	}
	return Parse[T](bytes)
}

func ParseString[T any](contents string) (*T, error) {
	return Parse[T]([]byte(contents))
}
