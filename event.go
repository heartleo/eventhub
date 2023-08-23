package eventhub

import "reflect"

type HandleFunc[D any] func(data D)

type Event[K comparable, D any] struct {
	k  K
	fv reflect.Value
	f  HandleFunc[D]
}

func newEvent[K comparable, D any](key K, fn HandleFunc[D]) *Event[K, D] {
	return &Event[K, D]{k: key, fv: reflect.ValueOf(fn), f: fn}
}
