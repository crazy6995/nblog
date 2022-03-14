package util

import (
	"testing"
)

func BenchmarkNUid_Generate(b *testing.B) {
	nUid := NewNUid()
	for i := 0; i < b.N; i++ {
		nUid.Generate()
		//fmt.Print(nUid.Generate())
		//fmt.Print("\t")
		//if (i+1)%line == 0 {
		//	fmt.Println()
		//}
	}
}
