package main

import (
	"testing"
)

func Test(t *testing.T) {
	m := New(httpGetBody)
	Sequential(m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	Concurrent(m)
}
