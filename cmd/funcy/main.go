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
	targetType, targetPath string
	generators             string
	keyField, groupFields  string
	verbose                bool
)

func init() {
	flag.StringVar(&targetType, "type", "", "Name of type to generate against")
	flag.StringVar(&targetPath, "path", "", "Type import path")
	flag.StringVar(&generators, "generators", "", "Name of generators to run")
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

	groupFieldValues := stringSliceOrNil(groupFields)
	generatorValues := stringSliceOrNil(generators)

	if err := funcy.Generate(targetPath, targetType, &model.Config{
		Generators: generatorValues,
		KeyField:    keyField,
		GroupFields: groupFieldValues,
	}); err != nil {
		logrus.Fatal(err)
	}
}

func stringSliceOrNil(value string) []string {
	values := strings.Split(value, ",")
	if len(values) == 1 && values[0] == "" {
		return nil
	}
	return values
}