package data

import "testing"

const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}

// 执行测试用例
// go test -bench=. -cpuprofile=cpu.prof

// 启动pprof可视化界面
// go tool pprof -http=:8080 cpu.prof
// go tool pprof cpu.prof
