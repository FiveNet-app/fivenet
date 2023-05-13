// Copied from Reddit user "maxmcd" here:
// https://www.reddit.com/r/golang/comments/twucb0/comment/j4x7xbx/?utm_source=reddit&utm_medium=web2x&context=3

package syncx

import "sync"

type Map[K comparable, V any] struct {
	m sync.Map
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	return v.(V), loaded
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool { return f(key.(K), value.(V)) })
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *Map[K, V]) Swap(key K, value V) (previous any, loaded bool) {
	prev, loaded := m.m.Swap(key, value)
	return prev.(V), loaded
}

func (m *Map[K, V]) Keys() []K {
	keys := []K{}

	m.m.Range(func(key any, value any) bool {
		keys = append(keys, key.(K))
		return true
	})

	return keys
}
