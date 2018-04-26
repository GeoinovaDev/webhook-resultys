package webhook

import (
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/motor/webhook/channel"
)

// Manager struct
type Manager struct {
	channel *channel.Channel
}

// New cria manager
func New(limit int) *Manager {
	manager := &Manager{}
	manager.channel = channel.New(limit)

	return manager
}

// Trigger dispara um webhook
func (manager *Manager) Trigger(url string, data interface{}) *Manager {
	if len(url) == 0 {
		return manager
	}

	manager.channel.Alloc(func() {
		request.New(url).PostJSON(data)
	})

	return manager
}
