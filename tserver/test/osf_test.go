package client

import "testing"
import (
	"go-grpc-test/tserver/client"
)

func BenchmarkHello(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.GetClient("127.0.0.1:10005")
		}
	})
}

func BenchmarkHello2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client.GetClient("127.0.0.1:10006")
		}
	})
}
