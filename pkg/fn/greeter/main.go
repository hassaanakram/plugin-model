package main

import (
	"fmt"
	"github.com/hassaanakram/plugin-model/pkg/plugins"
	"github.com/hassaanakram/plugin-model/pkg/plugins/registry"
)

/*
IMPLEMENTATION NOTE: The program execution follows the following steps:
- fmt is imported
- registry is imported
- github.com/hassaanakram/plugin-model/pkg/plugins is imported
  -- When this package is imported, the init() function defined in each plugin file is called which registers the plugins
     with the registry.PluginMap
- main() is executued
*/

/* ADDING NEW PLUGINS:
- Create a new plugin type in pkg/plugins and define the Greeter interace's methods. Please see the existing
  Greeters for understanding.
*/

func main() {
	fmt.Println(registry.PluginMap)
	// This statement is redundant but required as to register the plugins, the plugins package must be imported
	// and the compiler does not allow unused imports. This will be improved.
	_ = plugins.EnglishGreeter{}

	for _, greeter := range registry.PluginMap {
		greeter.Greet()
	}
}
