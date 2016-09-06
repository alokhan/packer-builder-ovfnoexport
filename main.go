package main

import (
	"github.com/alois/packer-builder-ovfnoexport/builder/ovfnoexport"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	// Choose the appropriate type of plugin. You should only use one of these
	// at a time, which means you will have a separate plugin for each builder,
	// provisioner, or post-processor.
	server.RegisterBuilder(new(ovfnoexport.Builder))
	server.Serve()
}
