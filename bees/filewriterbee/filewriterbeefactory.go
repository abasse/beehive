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
	"github.com/muesli/beehive/bees"
)

// FilewriterBeeFactory is a factory for FilewriterBees.
type FilewriterBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *FilewriterBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := FilewriterBee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *FilewriterBeeFactory) ID() string {
	return "Filewriterbee"
}

// Name returns the name of this Bee.
func (factory *FilewriterBeeFactory) Name() string {
	return "Filewriter"
}

// Description returns the description of this Bee.
func (factory *FilewriterBeeFactory) Description() string {
	return "Write text to file"
}

// Image returns the filename of an image for this Bee.
func (factory *FilewriterBeeFactory) Image() string {
	return factory.ID() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *FilewriterBeeFactory) LogoColor() string {
	return "#cccccc"
}

// Options returns the options available to configure this Bee.
func (factory *FilewriterBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "filename",
			Description: "Filename",
			Type:        "string",
			Mandatory:   true,
		},
	}
	return opts
}

// Events describes the available events provided by this Bee.
func (factory *FilewriterBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{}
	return events
}

// Actions describes the available actions provided by this Bee.
func (factory *FilewriterBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "write",
			Description: "Writes a text message to a file",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "message",
					Description: "Text to write",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
	}
	return actions
}

func init() {
	f := FilewriterBeeFactory{}
	bees.RegisterFactory(&f)
}
