package logic

import (
	"context"
	"fmt"

	"github.com/mattn/anko/env"
)

type Anko struct {
	ctx    context.Context
	env    *env.Env
	params map[string]any
}

func NewAnko() *Anko {
	e := env.NewEnv()
	e.Define("println", fmt.Println)
	return &Anko{
		ctx:    context.Background(),
		env:    e,
		params: make(map[string]any),
	}
}

// GetEnv
func (a *Anko) GetEnv() *env.Env {
	return a.env
}

// SetParams
func (a *Anko) SetParams(name string, elem any) {
	a.params[name] = elem
	a.env.Define(name, elem)
}

// GetParams
func (a *Anko) GetParams(name string) any {
	return a.params[name]
}
