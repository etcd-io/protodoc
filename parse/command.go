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

package parse

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use:   "parse",
		Short: "parses Protocol Buffer files to generate documentation.",
	}
)

var (
	title           string
	targetDir       string
	targetPath      string
	languageOptions []string
)

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	Command.PersistentFlags().StringVarP(&title, "title", "t", "", "title of documentation")
	Command.PersistentFlags().StringVarP(&targetDir, "target-directory", "d", "", "target directory with .proto files")
	Command.PersistentFlags().StringVarP(&targetPath, "target-path", "p", "", "file path to save the documentation")
	Command.PersistentFlags().StringSliceVarP(&languageOptions, "language-options", "o", []string{"Go", "C++", "Java", "Python"}, "language options in field descriptions")
}

func CommandFunc(cmd *cobra.Command, args []string) error {
	log.Println("opening", targetDir)
	proto, err := ReadDir(targetDir)
	if err != nil {
		return err
	}
	return proto.Markdown(title, targetPath, languageOptions...)
}
