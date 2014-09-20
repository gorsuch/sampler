package sampler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSampler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Howdy")
	}))
	defer ts.Close()

	s := New(10 * time.Second)
	sample, err := s.Sample(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if sample.StatusCode != 200 {
		t.Fatalf("%d != 200", sample.StatusCode)
	}

	if sample.T2.Sub(sample.T1) <= 0 {
		t.Fatal("T2 should be less than T1")
	}
}
