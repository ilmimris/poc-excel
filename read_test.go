package read_test

import (
	"testing"

	read "github.com/ilmimris/poc-excel"
)

func BenchmarkMethodA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		read.Method1()
	}
}

func BenchmarkMethodB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		read.Method2()
	}
}

func BenchmarkMethodC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		read.XlsxModule()
	}
}
