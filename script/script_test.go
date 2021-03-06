package script_test

import (
	"errors"
	"testing"

	"github.com/d5/tengo/assert"
	"github.com/d5/tengo/script"
)

func TestScript_Add(t *testing.T) {
	s := script.New([]byte(`a := b`))
	assert.NoError(t, s.Add("b", 5))     // b = 5
	assert.NoError(t, s.Add("b", "foo")) // b = "foo"  (re-define before compilation)
	c, err := s.Compile()
	assert.NoError(t, err)
	assert.NoError(t, c.Run())
	assert.Equal(t, "foo", c.Get("a").Value())
	assert.Equal(t, "foo", c.Get("b").Value())
}

func TestScript_Remove(t *testing.T) {
	s := script.New([]byte(`a := b`))
	err := s.Add("b", 5)
	assert.NoError(t, err)
	assert.True(t, s.Remove("b")) // b is removed
	_, err = s.Compile()          // should not compile because b is undefined
	assert.Error(t, err)
}

func TestScript_Run(t *testing.T) {
	s := script.New([]byte(`a := b`))
	err := s.Add("b", 5)
	assert.NoError(t, err)
	c, err := s.Run()
	assert.NoError(t, err)
	assert.NotNil(t, c)
	compiledGet(t, c, "a", int64(5))
}

func TestScript_DisableBuiltinFunction(t *testing.T) {
	s := script.New([]byte(`a := len([1, 2, 3])`))
	c, err := s.Run()
	assert.NoError(t, err)
	assert.NotNil(t, c)
	compiledGet(t, c, "a", int64(3))
	s.DisableBuiltinFunction("len")
	_, err = s.Run()
	assert.Error(t, err)
}

func TestScript_DisableStdModule(t *testing.T) {
	s := script.New([]byte(`math := import("math"); a := math.abs(-19.84)`))
	c, err := s.Run()
	assert.NoError(t, err)
	assert.NotNil(t, c)
	compiledGet(t, c, "a", 19.84)
	s.DisableStdModule("math")
	_, err = s.Run()
	assert.Error(t, err)
}

func TestScript_SetUserModuleLoader(t *testing.T) {
	s := script.New([]byte(`math := import("mod1"); a := math.foo()`))
	_, err := s.Run()
	assert.Error(t, err)
	s.SetUserModuleLoader(func(moduleName string) (res []byte, err error) {
		if moduleName == "mod1" {
			res = []byte(`foo := func() { return 5 }`)
			return
		}

		err = errors.New("module not found")
		return
	})
	c, err := s.Run()
	assert.NoError(t, err)
	assert.NotNil(t, c)
	compiledGet(t, c, "a", int64(5))
}
