package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://zanders.cloud"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://go.dev",
		"https://kubernetes.io",
		"https://zanders.cloud",
	}

	want := map[string]bool{
		"https://go.dev":        true,
		"https://kubernetes.io": true,
		"https://zanders.cloud": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
