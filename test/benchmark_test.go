package test

import (
	"testing"

	"github.com/bychannel/event"
)

func BenchmarkManager_Publish_no_listener(b *testing.B) {
	em := event.NewManager("test")
	em.Listen("aa.bb", event.ListenerFunc(func(e event.IEvent) error {
		return nil
	}))

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = em.Publish("aa.bb", nil)
	}
}

func BenchmarkManager_Publish_normal(b *testing.B) {
	em := event.NewManager("test")
	em.Listen("aa.bb", event.ListenerFunc(func(e event.IEvent) error {
		return nil
	}))

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = em.Publish("aa.bb", nil)
	}
}

func BenchmarkManager_Publish_wildcard(b *testing.B) {
	em := event.NewManager("test")
	em.Listen("aa.*", event.ListenerFunc(func(e event.IEvent) error {
		return nil
	}))

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = em.Publish("aa.bb", nil)
	}
}
