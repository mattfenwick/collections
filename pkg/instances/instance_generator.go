package instances

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/ioutil"
	"path"
	"strings"
)

var caser = cases.Title(language.AmericanEnglish)

func x() {
	//language.AmericanEnglish
	caser := cases.Title(language.AmericanEnglish)
	caser.String("abc")
}

const (
	wrapperTemplate = `type %s %s`

	eqTemplate = `func (a %s)Equal(b %s) bool {
    return a == b
}`

	ordTemplate = `func (a %s)Compare(b %s) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}`
)

func makeTypeWrapper(typeName string) string {
	return fmt.Sprintf(wrapperTemplate, caser.String(typeName), typeName)
}

func makeEqInstance(typeName string) string {
	name := caser.String(typeName)
	return fmt.Sprintf(eqTemplate, name, name)
}

func makeOrdInstance(typeName string) string {
	name := caser.String(typeName)
	return fmt.Sprintf(ordTemplate, name, name)
}

type Wrapper struct {
	Underlying  string
	EqInstance  bool
	OrdInstance bool
}

var (
	WrapperInstances = []*Wrapper{
		{"bool", true, false},
		{"uint", true, true},
		{"uint8", true, true},
		{"uint16", true, true},
		{"uint32", true, true},
		{"uint64", true, true},
		{"uintptr", true, true},
		{"int", true, true},
		{"int8", true, true},
		{"int16", true, true},
		{"int32", true, true},
		{"int64", true, true},
		{"float32", true, true},
		{"float64", true, true},
		{"complex64", true, false},
		{"complex128", true, false},
		{"string", true, true},
	}
)

func ModelToText(dir string, packageName string) error {
	var wrappers, eqs, ords []string
	for _, w := range WrapperInstances {
		wrappers = append(wrappers, makeTypeWrapper(w.Underlying))
		if w.EqInstance {
			eqs = append(eqs, makeEqInstance(w.Underlying))
		}
		if w.OrdInstance {
			ords = append(ords, makeOrdInstance(w.Underlying))
		}
	}

	err := WriteFile(path.Join(dir, "generated_wrappers.go"),
		strings.Join(append([]string{fmt.Sprintf("package %s", packageName)}, wrappers...), "\n\n"))
	if err != nil {
		return err
	}
	err = WriteFile(path.Join(dir, "generated_eq_instances.go"),
		strings.Join(append([]string{fmt.Sprintf("package %s", packageName)}, eqs...), "\n\n"))
	if err != nil {
		return err
	}
	return WriteFile(path.Join(dir, "generated_ord_instances.go"),
		strings.Join(append([]string{fmt.Sprintf("package %s", packageName)}, ords...), "\n\n"))
}

func WriteFile(filename string, contents string) error {
	return errors.Wrapf(ioutil.WriteFile(filename, []byte(contents), 0644), "unable to write file %s", filename)
}
