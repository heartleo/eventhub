package eventhub

import (
	"reflect"
)

type EventHuber[K comparable, D any] interface {
	On(key K, fn HandleFunc[D]) EventHuber[K, D]
	Off(key K, fn HandleFunc[D]) EventHuber[K, D]
	Fire(key K, data D) EventHuber[K, D]
}

type EventHub[K comparable, D any] struct {
	events map[K][]*Event[K, D]
}

func New[K comparable, D any]() EventHuber[K, D] {
	return &EventHub[K, D]{
		events: make(map[K][]*Event[K, D]),
	}
}

func (h *EventHub[K, D]) On(key K, fn HandleFunc[D]) EventHuber[K, D] {
	h.addEvent(key, fn)
	return h
}

func (h *EventHub[K, D]) Off(key K, fn HandleFunc[D]) EventHuber[K, D] {
	h.removeEvent(key, fn)
	return h
}

func (h *EventHub[K, D]) Fire(key K, data D) EventHuber[K, D] {
	for _, event := range h.events[key] {
		event.f(data)
	}
	return h
}

func (h *EventHub[K, D]) addEvent(key K, fn HandleFunc[D]) {
	h.events[key] = append(h.events[key], newEvent(key, fn))
}

func (h *EventHub[K, D]) removeEvent(key K, fn HandleFunc[D]) {
	fv := reflect.ValueOf(fn)
	es := h.events[key]
	for i, e := range es {
		if e.fv == fv {
			es[i] = nil // help gc
			es = append(es[:i], es[i+1:]...)
			break
		}
	}

	if len(es) > 0 {
		h.events[key] = es
	} else {
		delete(h.events, key)
	}
}
