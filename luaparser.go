package luaparser

import (
	"bytes"
	"fmt"

	"github.com/Azure/golua/lua"
	"github.com/Azure/golua/std"
)

type (

	// LuaParser parses a simple Lua script into a Go object
	LuaParser struct {
		GlobalVar string
		Debug     bool
	}
)

// Unmarshal takes Lua script raw data, executes it,
// then stores the result in the value pointed to by out
func (p *LuaParser) Unmarhsall(in []byte, out interface{}) error {
	var opts = []lua.Option{lua.WithTrace(p.Debug), lua.WithVerbose(p.Debug)}
	state := lua.NewState(opts...)
	defer state.Close()
	std.Open(state)

	r := bytes.NewReader(in)
	err := state.ExecFrom(r)
	if err != nil {
		return err
	}

	state.GetGlobal(p.GlobalVar)
	value := state.Pop()

	// TODO: figure out how to unmarshall value into out
	fmt.Println(lua.IsNumber(value))

	return nil
}
