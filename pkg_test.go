package testTool

import (
	"testing"
)

func Benchmark_GetGuid32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetGuid32()
	}
}

func Benchmark_RunFuncName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunFuncName()
	}
}

func Benchmark_PrintJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintJSON(map[string]string{"nihao": "你好", "shijie": "世界"})
	}
}

func Benchmark_GetMd5String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMd5String("9d9dce8ec1654ee28ad50ede7e04247b")
	}
}

func Benchmark_GetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString(i)
	}
}
