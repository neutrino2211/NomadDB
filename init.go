package main

import (
	"fmt"

	"github.com/neutrino2211/commander"
)

type InitCommand struct {
	commander.Command
	clusterAddresses []string
}

func (d *InitCommand) Init() {
	d.Logger.Init("init", 0)
	d.Optionals = map[string]*commander.Optional{
		"type": {
			Type:        "string",
			Description: "Type of cluster to initialise",
		},
	}

	d.Values = map[string]string{}

	d.Usage = "cluster init"
	d.Description = d.BuildHelp("Create a cluster type configuration")
}

func (d *InitCommand) Run() {
	typ := d.GetString("type", "database")

	d.Logger.LogString(fmt.Sprintf("Initialising a cluster of type %s", typ))
}
