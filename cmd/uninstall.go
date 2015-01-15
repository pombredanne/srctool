// Copyright 2014-2015 The DevMine authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"errors"
	"os"

	"github.com/codegangsta/cli"

	"github.com/DevMine/srctool/config"
	"github.com/DevMine/srctool/log"
)

// Uninstall command uninstalls a language parser.
// It expectes only one command line argument: the parser name.
func Uninstall(c *cli.Context) {
	if !c.Args().Present() {
		log.Fatal("expected 1 argument, found 0")
	}

	parserName := "parser-" + c.Args().First()

	if err := uninstallParser(parserName, c.Bool("dry")); err != nil {
		log.Fatal(err)
	}
}

func uninstallParser(parserName string, dryMode bool) error {
	parserPath := config.ParserPath(parserName)

	if _, err := os.Stat(parserPath); os.IsNotExist(err) {
		log.Debug(err)
		return errors.New("the parser is not installed")
	}

	if dryMode {
		log.Info("parser path:", parserPath)
		return nil
	}

	log.Debug("removing ", parserPath)
	if err := os.RemoveAll(parserPath); err != nil {
		log.Debug(err)
		return errors.New("failed to remove the parser")
	}

	log.Success("parser successfully removed")
	return nil
}