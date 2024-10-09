/*
 *    Copyright (C) 2014-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Raphael Mutschler <info@raphaelmutschler.de>
 */

// Package Filewriterbee is a Bee that can send Filewriter notifications.
package filewriterbee

import (
	"os"

	"github.com/muesli/beehive/bees"
)

// FilewriterBee is a Bee that sends Filewriter notifications
type FilewriterBee struct {
	bees.Bee

	filename     string
}

// Run executes the Bee's event loop.
func (mod *FilewriterBee) Run(cin chan bees.Event) {
	select {
	case <-mod.SigChan:
		return
	}
}

// Action triggers the action passed to it.
func (mod *FilewriterBee) Action(action bees.Action) []bees.Placeholder {
	outs := []bees.Placeholder{}

	switch action.Name {
	case "write":
		var message string
		action.Options.Bind("message", &message)

	    mod.LogDebugf("Filewriter message: %s", message)

		//write message to file, create file if it does not exist, append the message at the end of the file
		file, err := os.OpenFile(mod.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			mod.LogErrorf("Failed to open file: %v", err)
			return outs
		}
		defer file.Close()

		if _, err := file.WriteString(message + "\n"); err != nil {
			mod.LogErrorf("Failed to write to file: %v", err)
			return outs
		}

	default:
		panic("Unknown action triggered in " + mod.Name() + ": " + action.Name)
	}

	return outs
}

// ReloadOptions parses the config options and initializes the Bee.
func (mod *FilewriterBee) ReloadOptions(options bees.BeeOptions) {
	mod.SetOptions(options)
	options.Bind("filename", &mod.filename)
}
