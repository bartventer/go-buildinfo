package gobuildinfo

import (
	"fmt"
)

const logo = `
_|                  _|  _|        _|  _|                _|_|
_|_|_|    _|    _|      _|    _|_|_|      _|_|_|      _|        _|_|
_|    _|  _|    _|  _|  _|  _|    _|  _|  _|    _|  _|_|_|_|  _|    _|
_|    _|  _|    _|  _|  _|  _|    _|  _|  _|    _|    _|      _|    _|
_|_|_|      _|_|_|  _|  _|    _|_|_|  _|  _|    _|    _|        _|_|
`

func ExampleInfo_String() {
	info := New(
		WithVersion("v1.0.0"),
		WithCommit("a696fbbcb8ae009e3f88df2d7b00c09bea903c9e"),
		WithDate("2023-01-01"),
		WithTreeState("clean"),
		WithProject("MyApp", "This is a sample app", "https://example.com"),
		WithASCIILogo(logo),
	)
	info.runtime.GoVersion = "go1.24.2"

	fmt.Println(info.String())
	// Output:
	// _|                  _|  _|        _|  _|                _|_|
	// _|_|_|    _|    _|      _|    _|_|_|      _|_|_|      _|        _|_|
	// _|    _|  _|    _|  _|  _|  _|    _|  _|  _|    _|  _|_|_|_|  _|    _|
	// _|    _|  _|    _|  _|  _|  _|    _|  _|  _|    _|    _|      _|    _|
	// _|_|_|      _|_|_|  _|  _|    _|_|_|  _|  _|    _|    _|        _|_|
	//
	// MyApp: This is a sample app
	// https://example.com
	//
	// Version    v1.0.0
	// Commit     a696fbbcb8ae009e3f88df2d7b00c09bea903c9e
	// Date       2023-01-01
	// TreeState  clean
	// GoVersion  go1.24.2
	// Compiler   gc
	// Platform   linux/amd64
	// ModuleSum  none
}
