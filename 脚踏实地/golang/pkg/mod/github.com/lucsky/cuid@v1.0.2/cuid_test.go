package cuid

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func Test_CUIDType(t *testing.T) {
	c := New()
	if reflect.TypeOf(c).Name() != "string" {
		t.Error("Incorrect CUID type")
	}
}

func Test_CUIDFormat(t *testing.T) {
	c := New()
	format := regexp.MustCompile(fmt.Sprintf("c[0-9a-z]{%d}", 6*blockSize))
	if !format.MatchString(c) {
		t.Error("Incorrect format")
	}
}

func Test_CUIDCollisions(t *testing.T) {
	ids := map[string]bool{}
	for i := 0; i < 600000; i++ {
		id := New()
		if ids[id] == true {
			t.Errorf("Collision detected, at iteration %d", i)
		}
		ids[id] = true
	}
}

func Test_CUIDSlugFormat(t *testing.T) {
	c := Slug()
	if len(c) < 6 || len(c) > 12 {
		t.Errorf("Slug incorrect format. Len: %d", len(c))
	}
}

func Test_CUIDSlugCollisions(t *testing.T) {
	ids := map[string]bool{}
	for i := 0; i < 60000; i++ {
		id := Slug()
		if ids[id] == true {
			t.Errorf("Collision detected, at iteration %d", i)
		}
		ids[id] = true
	}
}

func newCUID(chn chan error) {
	New()
	chn <- nil
}

func Test_DataRaces(t *testing.T) {
	chn := make(chan error)

	go newCUID(chn)
	go newCUID(chn)

	<-chn
	<-chn
}

func Benchmark_CUIDGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func Benchmark_CUIDSlugGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slug()
	}
}
