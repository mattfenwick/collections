package main

import (
	"github.com/mattfenwick/collections/pkg/instances"
	"github.com/sirupsen/logrus"
)

func main() {
	DoOrDie(instances.ModelToText("pkg", "pkg"))
}

func DoOrDie(err error) {
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
}
