package luaparser

import (
	"bytes"
	"encoding/json"
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
	table := state.Pop().(lua.Table)

	m := map[string]interface{}{}
	table.ForEach(parseTableFunc(m))

	// TODO: figure out how to unmarshall value into out
	fmt.Println(m)

	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}

func parseTableFunc(m map[string]interface{}) func(key lua.Value, value lua.Value) {
	return func(k lua.Value, v lua.Value) {
		ks := k.String()
		if v.Type() == lua.TableType {
			tmp := map[string]interface{}{}
			v.(lua.Table).ForEach(parseTableFunc(tmp))
			m[ks] = &tmp
		} else {
			m[ks] = v
		}
	}
}