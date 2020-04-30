package webhook

import (
	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/logfile"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/motor/resource"
)

// Manager struct
type Manager struct {
	resource   *resource.Resource
	logSuccess *logfile.Log
}

// New cria m
func New(limit int) *Manager {
	return &Manager{
		resource:   resource.New(limit),
		logSuccess: logfile.New("webhook_success.log").Limit(1000),
	}
}

// Trigger dispara um webhook
func (m *Manager) Trigger(url string, data interface{}) *Manager {
	if len(url) == 0 {
		return m
	}

	m.resource.Alloc(func() {
		exec.Tryx(3, func() {
			msg, err := request.New(url).PostJSON(data)
			if err != nil {
				panic(msg)
			}
			// m.logSuccess.Add(url)
		}).Catch(func(err string) {
			exception.Raise(err, exception.CRITICAL)
		})
	})

	return m
}
