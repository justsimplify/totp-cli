package modules

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Modules to store all modules
var Modules = make(map[string]func() interface{})

// GetModule is an interface for get related operations
type GetModule interface {
	Get(cmd *cobra.Command, args []string) (interface{}, error)
}

// FlagModule is an interface to resource that implement disable command
type FlagModule interface {
	Flags(cmd *cobra.Command)
}

// ModuleAdd is a function to add new modules
func ModuleAdd(moduleName string, creator func() interface{}) {
	Modules[moduleName] = creator
}

// Get returns the resource of the required type from the modules map
func Get(moduleName string) (interface{}, error) {
	creator, ok := Modules[moduleName]
	if !ok {
		return nil, fmt.Errorf("module '%s' not supported", moduleName)
	}

	return creator(), nil
}

// GetAllModules get all modules
func GetAllModules() []string {
	var res []string
	for k := range Modules {
		res = append(res, k)
	}
	return res
}
