// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// protodoc generates Protocol Buffer documentation.
//
//	Usage:
//	protodoc [flags]
//
//	Flags:
//	-h, --help                 help for protodoc
//	-o, --languages value      language options in field descriptions (default [Go,C++,Java,Python])
//	-p, --target-path string   file path to save the documentation
//	-t, --title string         title of documentation
//
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coreos/prodoc/parse"
	"github.com/spf13/cobra"
)

var (
	rootCommand = &cobra.Command{
		Use:   "protodoc",
		Short: "protodoc generates Protocol Buffer documentation.",
		RunE:  CommandFunc,
	}

	title           string
	targetDir       string
	targetPath      string
	languageOptions []string
)

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	rootCommand.PersistentFlags().StringVarP(&title, "title", "t", "", "title of documentation")
	rootCommand.PersistentFlags().StringVarP(&targetPath, "target-path", "p", "", "file path to save the documentation")
	rootCommand.PersistentFlags().StringSliceVarP(&languageOptions, "languages", "o", []string{"Go", "C++", "Java", "Python"}, "language options in field descriptions")
}

func CommandFunc(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("need 1 argument of target directory, got %q", args)
	}
	targetDir := args[0]
	log.Println("opening", targetDir)
	proto, err := parse.ReadDir(targetDir)
	if err != nil {
		return err
	}
	err = proto.Markdown(title, targetPath, languageOptions...)
	if err != nil {
		return err
	}
	log.Printf("saved at %s", targetPath)
	return nil
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
