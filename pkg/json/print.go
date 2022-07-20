package json

import "fmt"

func Print(obj interface{}) {
	PrintOptions(obj, DefaultMarshalOptions)
}

func PrintOptions(obj interface{}, options *MarshalOptions) {
	bytes, err := MarshalWithOptions(obj, options)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(bytes))
}
