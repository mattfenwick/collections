package file

import (
	"github.com/pkg/errors"
	"io/fs"
	"io/ioutil"
	"os"
)

func DoesFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(errors.Wrapf(err, "unable to determine if file %s exists", path))
	}
}

// WriteFile wraps calls to ioutil.WriteFile, ensuring that errors are wrapped in a stack trace
func WriteFile(filename string, contents string, perm fs.FileMode) error {
	return errors.Wrapf(ioutil.WriteFile(filename, []byte(contents), perm), "unable to write file %s", filename)
}

// WriteFileBytes wraps calls to ioutil.WriteFile, ensuring that errors are wrapped in a stack trace
func WriteFileBytes(filename string, bytes []byte, perm fs.FileMode) error {
	return errors.Wrapf(ioutil.WriteFile(filename, bytes, perm), "unable to write file %s", filename)
}

// ReadFile wraps calls to ioutil.ReadFile, ensuring that errors are wrapped in a stack trace
func ReadFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	return string(bytes), errors.Wrapf(err, "unable to read file %s", filename)
}

// ReadFileBytes wraps calls to ioutil.ReadFile, ensuring that errors are wrapped in a stack trace
func ReadFileBytes(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	return bytes, errors.Wrapf(err, "unable to read file %s", filename)
}
