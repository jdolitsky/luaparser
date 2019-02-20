package luaparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LuaParserTestSuite struct {
	suite.Suite
}

func (suite *LuaParserTestSuite) SetupSuite() {

}

func (suite *LuaParserTestSuite) TearDownSuite() {

}

func (suite *LuaParserTestSuite) Test_0_Unmarshall() {
	type TestType struct {
		A string   `json:"a"`
		B int      `json:"b"`
		C []string `json:"c"`
	}

	globalVar := "myglobalvar"

	in := []byte(
		fmt.Sprintf(`
local a = "x"
local b = 23
local c = {"a", "b", "c", "d"}

%s = {a = a, b = b, c = c}
		`, globalVar),
	)

	out := TestType{}
	parser := LuaParser{globalVar, true}

	err := parser.Unmarhsall(in, &out)
	suite.Nil(err)
	suite.Equal("x", out.A)
	suite.Equal("23", out.B)
	suite.Equal("23", out.C)
}

func TestLuaParserTestSuite(t *testing.T) {
	suite.Run(t, new(LuaParserTestSuite))
}
