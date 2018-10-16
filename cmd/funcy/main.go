package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shousper/go-funcy"
	"github.com/shousper/go-funcy/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	targetPath  string
	generators  []string
	keyField    string
	groupFields []string
	verbose     bool

	rootCmd = &cobra.Command{
		Use:   "funcy [flags] TYPE",
		Short: "funcy is a \"generic\" code generation tool",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Info("Processing ", args[0], " from ", targetPath)

			if err := funcy.Generate(targetPath, args[0], &model.Config{
				Generators:  generators,
				KeyField:    keyField,
				GroupFields: groupFields,
			}); err != nil {
				logrus.Fatal(err)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&targetPath, "path", "p", "", "Type import path, can be relative to GOPATH")
	rootCmd.PersistentFlags().StringSliceVarP(&generators, "generators", "g", nil, "list of generators to run")
	rootCmd.PersistentFlags().StringVarP(&keyField, "key-field", "k", "ID", "name of field on type to populate map key")
	rootCmd.PersistentFlags().StringSliceVarP(&groupFields, "group-fields", "f", nil, "name of fields on type to group by")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")
}

func initConfig() {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if targetPath == "" {
		var err error
		targetPath, err = os.Getwd()
		if err != nil {
			fmt.Println("invalid current directory: ", err)
			os.Exit(1)
		}
		targetPath, err = filepath.Abs(targetPath)
		if err != nil {
			fmt.Println("unable to resolve target path: ", err)
			os.Exit(1)
		}
		if targetPath == "" {
			fmt.Println("missing import path argument")
			os.Exit(1)
		}
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
