package rand

import "testing"

const n = 32

func BenchmarkBytesMaskImprSrcSB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(n)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(n)
	}
}
