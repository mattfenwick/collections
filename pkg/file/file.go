package file

import (
	"github.com/pkg/errors"
	"io/fs"
	"os"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(errors.Wrapf(err, "unable to determine if file %s exists", path))
	}
}

// WriteString wraps calls to os.WriteFile, ensuring that errors are wrapped in a stack trace
func WriteString(filename string, contents string, perm fs.FileMode) error {
	return Write(filename, []byte(contents), perm)
}

// Write wraps calls to os.WriteFile, ensuring that errors are wrapped in a stack trace
func Write(filename string, bytes []byte, perm fs.FileMode) error {
	return errors.Wrapf(os.WriteFile(filename, bytes, perm), "unable to write file %s", filename)
}

// ReadString wraps calls to os.ReadFile, ensuring that errors are wrapped in a stack trace
func ReadString(filename string) (string, error) {
	bytes, err := Read(filename)
	return string(bytes), err
}

// Read wraps calls to os.ReadFile, ensuring that errors are wrapped in a stack trace
func Read(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	return bytes, errors.Wrapf(err, "unable to read file %s", filename)
}
