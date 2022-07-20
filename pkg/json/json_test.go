package json

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

type testType struct {
	Q int `json:"q"`
	C int `json:"c"`
	M int `json:"m"`
	B int `json:"b"`
	Z int `json:"z"`
	L int `json:"l"`
	O int `json:"o"`
}

func RunJsonTests() {
	Describe("Json", func() {
		jsonString := `{
  "b": 3,
  "c": 17,
  "l": 14,
  "m": 16,
  "o": 17,
  "q": 15,
  "z": 2
}
`
		expected := &testType{
			B: 3,
			C: 17,
			L: 14,
			M: 16,
			O: 17,
			Q: 15,
			Z: 2,
		}
		expectedMap := map[string]int{
			"b": 3,
			"c": 17,
			"l": 14,
			"m": 16,
			"o": 17,
			"q": 15,
			"z": 2,
		}

		It("Read into type", func() {
			parsed, err := ParseString[testType](jsonString)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(parsed).To(gomega.Equal(expected))
		})

		It("Read into map", func() {
			parsed, err := ParseString[map[string]int](jsonString)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(*parsed).To(gomega.Equal(expectedMap))
		})

		It("Write map", func() {
			s, err := MarshalToString(expectedMap)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(s).To(gomega.Equal(jsonString))

			bytes, err := MarshalWithOptions(expectedMap, DefaultMarshalOptions)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(bytes).To(gomega.Equal([]byte(jsonString)))
		})

		It("Write type: sort order IS NOT predictable", func() {
			s, err := MarshalToString(expected)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("comparing strings: \n\n%s\n\n%s\n\n\n", s, jsonString)
			gomega.Expect(s).ToNot(gomega.Equal(jsonString))

			bytes, err := MarshalWithOptions(expected, DefaultMarshalOptions)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("comparing bytes: \n\n%s\n\n%s\n\n\n", string(bytes), jsonString)
			gomega.Expect(bytes).ToNot(gomega.Equal([]byte(jsonString)))
		})

		It("Marshal/sort type: sort order IS NOW predictable", func() {
			sortOptions := &MarshalOptions{EscapeHTML: true, Indent: true, Sort: true}
			typeBytes, err := MarshalWithOptions(expected, sortOptions)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("comparing from type: \n\n%s\n\n%s\n\n\n", string(typeBytes), jsonString)
			gomega.Expect(string(typeBytes)).To(gomega.Equal(jsonString))

			mapBytes, err := MarshalWithOptions(expectedMap, sortOptions)
			gomega.Expect(err).To(gomega.BeNil())
			fmt.Printf("comparing from map: \n\n%s\n\n%s\n\n\n", string(mapBytes), jsonString)
			gomega.Expect(string(mapBytes)).To(gomega.Equal(jsonString))
		})
	})
}
