package channel

import (
	"sync"

	"git.resultys.com.br/lib/lower/collection/queue"
)

// Channel struct
type Channel struct {
	Limit int

	used      int
	mutex     *sync.Mutex
	onRelease func()
	fila      *queue.Queue
}

// New cria o ralloc
func New(limit int) *Channel {
	channel := &Channel{Limit: limit}

	channel.mutex = &sync.Mutex{}
	channel.fila = queue.New()

	return channel
}

// Alloc aloca canal de comunicação
func (channel *Channel) Alloc(callback func()) *Channel {
	channel.mutex.Lock()

	if channel.used == channel.Limit {
		channel.fila.Push(item{cb: callback})
		channel.mutex.Unlock()
		return channel
	}

	channel.used++
	channel.mutex.Unlock()

	callback()

	channel.Release()

	return channel
}

// Release libera o canal
func (channel *Channel) Release() *Channel {
	channel.mutex.Lock()
	channel.used--
	channel.mutex.Unlock()

	if channel.fila.IsEmpty() {
		return channel
	}

	item := channel.fila.Pop().(item)
	channel.Alloc(item.cb)

	return channel
}
