package json

import (
	"bytes"
	"encoding/json"
	"github.com/mattfenwick/collections/pkg/file"
	"github.com/pkg/errors"
)

type MarshalOptions struct {
	EscapeHTML bool
	Indent     bool
	Sort       bool
}

var DefaultMarshalOptions = &MarshalOptions{EscapeHTML: true, Indent: true, Sort: false}

func MarshalHelper(obj interface{}, escapeHtml bool, indent bool) ([]byte, error) {
	encodeBuffer := &bytes.Buffer{}
	encoder := json.NewEncoder(encodeBuffer)
	encoder.SetEscapeHTML(escapeHtml)
	err := encoder.Encode(obj)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to encode json")
	}

	if indent {
		var indentBuffer bytes.Buffer
		err = json.Indent(&indentBuffer, encodeBuffer.Bytes(), "", "  ")
		if err != nil {
			return nil, errors.Wrapf(err, "unable to indent json")
		}
		return indentBuffer.Bytes(), nil
	}
	return encodeBuffer.Bytes(), nil
}

func MarshalWithOptions(obj interface{}, options *MarshalOptions) ([]byte, error) {
	if options.Sort {
		unsortedBytes, err := MarshalHelper(obj, options.EscapeHTML, false)
		if err != nil {
			return nil, err
		}
		return SortOptions(unsortedBytes, false, options.Indent)
	} else {
		return MarshalHelper(obj, options.EscapeHTML, options.Indent)
	}
}

func MarshalToFileOptions(obj interface{}, path string, options *MarshalOptions) error {
	content, err := MarshalWithOptions(obj, options)
	if err != nil {
		return err
	}
	return file.WriteFileBytes(path, content, 0644)
}

func MarshalToFile(obj interface{}, path string) error {
	return MarshalToFileOptions(obj, path, DefaultMarshalOptions)
}

func MarshalToString(obj interface{}) (string, error) {
	bytes, err := MarshalWithOptions(obj, DefaultMarshalOptions)
	return string(bytes), err
}

func Marshal(obj interface{}) ([]byte, error) {
	return MarshalWithOptions(obj, DefaultMarshalOptions)
}
