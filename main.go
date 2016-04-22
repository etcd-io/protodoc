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
//	  protodoc [flags]
//
//	Flags:
//	      --directories value                    comma separated map of target directory to parse options (e.g. 'dirA=message,dirB=message_service')
//	  -d, --directory string                     target directory where Protocol Buffer files are.
//	  -h, --help                                 help for protodoc
//	  -l, --languages value                      language options in field descriptions (Go, C++, Java, Python, Ruby, C#) (default [])
//	      --message-only-from-this-file string   if specified, it parses only the messages in this file within the directory
//	  -o, --output string                        output file path to save documentation
//	  -p, --parse value                          Protocol Buffer types to parse (message, service) (default [service,message])
//	  -t, --title string                         title of documentation
//
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/coreos/protodoc/parse"
	"github.com/spf13/cobra"
)

var (
	rootCommand = &cobra.Command{
		Use:   "protodoc",
		Short: "protodoc generates Protocol Buffer documentation.",
		RunE:  CommandFunc,
	}

	targetDirectory string
	parseOptions    []string
	languageOptions []string
	title           string
	outputPath      string

	targetDirectories       = newMapString()
	messageOnlyFromThisFile string
)

type mapString map[string][]parse.ParseOption

func newMapString() mapString {
	return mapString(make(map[string][]parse.ParseOption))
}
func (m mapString) String() string {
	return ""
}
func (m mapString) Set(s string) error {
	for _, elem := range strings.Split(s, ",") {
		pair := strings.Split(elem, "=")
		if len(pair) != 2 {
			return fmt.Errorf("invalid format %s", pair)
		}
		opts := []parse.ParseOption{}
		for _, v := range strings.Split(pair[1], "_") {
			switch v {
			case "message":
				opts = append(opts, parse.ParseMessage)
			case "service":
				opts = append(opts, parse.ParseService)
			}
		}
		m[pair[0]] = opts
	}
	return nil
}
func (m mapString) Type() string {
	return "map[string][]parse.ParseOption"
}

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	rootCommand.PersistentFlags().StringVarP(&targetDirectory, "directory", "d", "", "target directory where Protocol Buffer files are.")
	rootCommand.PersistentFlags().StringSliceVarP(&parseOptions, "parse", "p", []string{"service", "message"}, "Protocol Buffer types to parse (message, service)")
	rootCommand.PersistentFlags().StringSliceVarP(&languageOptions, "languages", "l", []string{}, "language options in field descriptions (Go, C++, Java, Python, Ruby, C#)")
	rootCommand.PersistentFlags().StringVarP(&title, "title", "t", "", "title of documentation")
	rootCommand.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "output file path to save documentation")

	rootCommand.PersistentFlags().Var(&targetDirectories, "directories", "comma separated map of target directory to parse options (e.g. 'dirA=message,dirB=message_service')")
	rootCommand.PersistentFlags().StringVar(&messageOnlyFromThisFile, "message-only-from-this-file", "", "if specified, it parses only the messages in this file within the directory")
}

func CommandFunc(cmd *cobra.Command, args []string) error {
	var rs string
	if len(targetDirectories) == 0 {
		log.Println("opening", targetDirectory)
		proto, err := parse.ReadDir(targetDirectory, "")
		if err != nil {
			return err
		}
		opts := []parse.ParseOption{}
		for _, v := range parseOptions {
			switch v {
			case "message":
				opts = append(opts, parse.ParseMessage)
			case "service":
				opts = append(opts, parse.ParseService)
			}
		}
		log.Println("converting to markdown", title)
		rs, err = proto.Markdown(title, opts, languageOptions...)
		if err != nil {
			return err
		}
	} else {
		for k, opts := range targetDirectories {
			log.Println("opening", k)
			c1 := filepath.Base(filepath.Dir(messageOnlyFromThisFile))
			c2 := filepath.Base(k)
			bs := ""
			if c1 == c2 {
				bs = messageOnlyFromThisFile
				log.Println("message only from this file:", messageOnlyFromThisFile)
			}
			proto, err := parse.ReadDir(k, bs)
			if err != nil {
				return err
			}
			ms, err := proto.Markdown("", opts, languageOptions...)
			if err != nil {
				return err
			}
			rs += ms
		}
		rs = fmt.Sprintf("### %s\n\n\n", title) + rs
	}
	err := toFile(rs, outputPath)
	if err != nil {
		return err
	}
	log.Printf("saved at %s", outputPath)
	return nil
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func toFile(txt, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	_, err = f.WriteString(txt)
	return err
}
