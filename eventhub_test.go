package eventhub

import (
	"fmt"
	"testing"
)

func ExampleEventHub() {
	m := New[int, int]()

	fn1 := func(data int) { fmt.Println("1 call=", data) }
	fn2 := func(data int) { fmt.Println("2 call=", data) }
	fn3 := func(data int) { fmt.Println("3 call=", data) }
	fn4 := func(data int) { fmt.Println("4 call=", data) }

	m.On(1, fn1).On(2, fn2).On(3, fn3).On(4, fn4)

	m.Fire(1, 1)
	m.Fire(2, 2)
	m.Fire(3, 3)
	m.Fire(4, 4)

	m.Off(2, fn2).Off(3, fn3)

	m.Fire(1, 100)
	m.Fire(2, 200)
	m.Fire(3, 300)
	m.Fire(4, 400)

	// Output:
	// 1 call= 1
	// 2 call= 2
	// 3 call= 3
	// 4 call= 4
	// 1 call= 100
	// 4 call= 400
}

func BenchmarkOn(b *testing.B) {
	m := New[int, int]()
	fn := func(data int) {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.On(i, fn)
	}
}

func BenchmarkOff100000(b *testing.B) {
	m := New[int, int]()
	fn := func(data int) {}
	for i := 0; i < 100000; i++ {
		m.On(i, fn)
	}
	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		m.Off(i, fn)
	}
}

func BenchmarkFire100000(b *testing.B) {
	m := New[int, int]()
	fn := func(data int) {}
	for i := 0; i < 100000; i++ {
		m.On(i, fn)
	}
	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		m.Fire(i, i)
	}
}
