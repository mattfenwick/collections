package yaml

import (
	"bytes"
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io"
)

func ParseManyFromFile[A any](path string) ([]A, error) {
	data, err := file.Read(path)
	if err != nil {
		return nil, err
	}
	return ParseMany[A](data)
}

func ParseMany[A any](data []byte) ([]A, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))

	var out []A
	for {
		var next A
		err := decoder.Decode(&next)
		if err == io.EOF {
			break
		} else if err != nil {
			return out, errors.Wrapf(err, "unable to decode")
		}
		out = append(out, next)
	}

	return out, nil
}
