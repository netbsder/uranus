package uranus

import (
	"sort"
	"sync"
)

type Module interface {
	RegisterRouter()
}

var (
	modulesMu sync.RWMutex
	modules   = make(map[string]Module)
)

// Register makes a uranus module available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, module Module) {
	modulesMu.Lock()
	defer modulesMu.Unlock()

	if module == nil {
		panic("uranus: Register module is nil")
	}

	if _, dup := modules[name]; dup {
		panic("uranus: Register called twice for module " + name)
	}

	modules[name] = module
	module.RegisterRouter()
}

// Modules returns a sorted list of the names of the registered modules.
func Modules() []string {
	modulesMu.RLock()
	defer modulesMu.RUnlock()

	var list []string
	for name := range modules {
		list = append(list, name)
	}

	sort.Strings(list)
	return list
}
