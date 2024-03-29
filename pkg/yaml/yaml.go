package yaml

import (
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/pkg/errors"
	yamlv2 "gopkg.in/yaml.v2"
	"gopkg.in/yaml.v3"
)

func Parse[T any](bs []byte) (*T, error) {
	var t T
	if err := yaml.Unmarshal(bs, &t); err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal yaml")
	}
	return &t, nil
}

func ParseStrict[T any](bs []byte) (*T, error) {
	var t T
	if err := yamlv2.UnmarshalStrict(bs, &t); err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal yaml")
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

func ParseFileStrict[T any](path string) (*T, error) {
	bytes, err := file.Read(path)
	if err != nil {
		return nil, err
	}
	return ParseStrict[T](bytes)
}

func Marshal(obj interface{}) ([]byte, error) {
	bytes, err := yaml.Marshal(obj)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to marshal yaml")
	}
	return bytes, nil
}

func MarshalString(obj interface{}) (string, error) {
	bytes, err := Marshal(obj)
	return string(bytes), err
}

func MarshalFile(obj interface{}, path string) error {
	bytes, err := Marshal(obj)
	if err != nil {
		return err
	}
	return file.Write(path, bytes, 0644)
}
