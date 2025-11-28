package events

import (
	"sync"
	"time"
)

type Event struct {
	Name      string
	Payload   any
	TimeStamp time.Time
}

type Handlerfunc func(Event)

type EventBus struct {
	Subscribers map[string][]Handlerfunc
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		Subscribers: make(map[string][]Handlerfunc),
	}
}

func (b *EventBus) Subscribe(topic string, handler Handlerfunc) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Subscribers[topic] = append(b.Subscribers[topic], handler)
}

// wildcard support: "order.*"
func match(topic, pattern string) bool {
	if pattern == topic {
		return true
	}
	if len(pattern) > 2 && pattern[len(pattern)-1] == '*' {
		prefix := pattern[:len(pattern)-1]
		return len(topic) >= len(prefix) && topic[:len(prefix)] == prefix
	}
	return false
}

func (b *EventBus) Publish(topic string, event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for pattern, handlers := range b.Subscribers {
		if match(topic, pattern) {
			for _, handler := range handlers {
				go handler(event) // async
			}
		}
	}
}
