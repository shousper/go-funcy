package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shousper/go-funcy"
	"github.com/shousper/go-funcy/model"
	"github.com/sirupsen/logrus"
)

var (
	targetPath, targetType, keyField, groupFields string
	verbose                                       bool
)

func init() {
	flag.StringVar(&targetType, "type", "", "Name of type to generate against")
	flag.StringVar(&targetPath, "path", "", "Type import path")
	flag.StringVar(&keyField, "key-field", "ID", "Name of map key field")
	flag.StringVar(&groupFields, "group-fields", "", "Name of fields to group by (comma delimited)")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.Parse()

	if targetType == "" {
		panic(fmt.Errorf("missing type argument"))
	}

	if targetPath == "" {
		targetPath, _ = os.Getwd()
		if targetPath == "" {
			panic(fmt.Errorf("missing import path argument"))
		}
		targetPath, _ = filepath.Abs(targetPath)
		if targetPath == "" {
			panic(fmt.Errorf("missing import path argument"))
		}
	}

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	logrus.Info("Processing ", targetType, " from ", targetPath)

	groupFieldValues := strings.Split(groupFields, ",")
	if len(groupFieldValues) == 1 && groupFieldValues[0] == "" {
		groupFieldValues = nil
	}

	if err := funcy.Generate(targetPath, targetType, &model.Config{
		KeyField:    keyField,
		GroupFields: groupFieldValues,
	}); err != nil {
		logrus.Fatal(err)
	}
}
