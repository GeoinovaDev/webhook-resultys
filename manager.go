package webhook

import (
	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/motor/resource"
)

// Manager struct
type Manager struct {
	resource *resource.Resource
}

// New cria manager
func New(limit int) *Manager {
	manager := &Manager{}
	manager.resource = resource.New(limit)

	return manager
}

// Trigger dispara um webhook
func (manager *Manager) Trigger(url string, data interface{}) *Manager {
	if len(url) == 0 {
		return manager
	}

	// manager.resource.Alloc(func() {
	exec.Tryx(3, func() {
		msg, err := request.New(url).PostJSON(data)
		if err != nil {
			panic(msg)
		}
	}).Catch(func(err string) {
		exception.Raise(err, exception.CRITICAL)
	})
	// })

	return manager
}
