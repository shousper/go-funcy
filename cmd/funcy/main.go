package main

import (
	"flag"
	"fmt"

	"github.com/shousper/go-funcy"
	"github.com/shousper/go-funcy/model"
	"github.com/sirupsen/logrus"
)

var (
	targetPath, targetType, keyField string
	verbose                          bool
)

func init() {
	flag.StringVar(&targetPath, "path", "", "Type import path")
	flag.StringVar(&targetType, "type", "", "Names of type to generate against")
	flag.StringVar(&keyField, "key-field", "ID", "Name of map key field")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.Parse()

	if targetPath == "" {
		panic(fmt.Errorf("missing import path argument"))
	}
	if targetType == "" {
		panic(fmt.Errorf("missing type argument"))
	}

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	logrus.Info("Processing ", targetType, " from ", targetPath)
	if err := funcy.Generate(targetPath, targetType, &model.Config{KeyField: keyField}); err != nil {
		panic(err)
	}
}
