package webhook

import (
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

	manager.resource.Alloc(func() {
		request.New(url).PostJSON(data)
	})

	return manager
}
