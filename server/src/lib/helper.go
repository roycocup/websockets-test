package lib

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func Check(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

func DD(o interface{}) {
	fmt.Println("############")
	spew.Dump(o)
	fmt.Println("############")
	os.Exit(100)
}
