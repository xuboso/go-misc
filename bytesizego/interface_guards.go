package main

import "io"

type A struct{}

var _ io.ReadCloser = (*A)(nil)

func New() A {
	return A{}
}

func (a A) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (a A) Close() error {
	panic("implement me")
}
