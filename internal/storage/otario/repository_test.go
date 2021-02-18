package otario

import "testing"

func BenchmarkGetBerrs(b *testing.B) {
	repository := NewRepository()

	for n := 0; n < b.N; n++ {
		repository.GetBeers()
	}
}
