package gobuildinfo_test

import (
	_ "embed"
	"fmt"

	gobuildinfo "github.com/bartventer/go-buildinfo"
)

//go:embed testdata/logo.txt
var logo string

func ExampleInfo_String_withoutRuntime() {
	info := gobuildinfo.New(
		gobuildinfo.WithVersion("v1.0.0"),
		gobuildinfo.WithCommit("a696fbbcb8ae009e3f88df2d7b00c09bea903c9e"),
		gobuildinfo.WithDate("2023-01-01"),
		gobuildinfo.WithTreeState("clean"),
		gobuildinfo.WithProject("MyApp", "This is a sample app", "https://example.com"),
		gobuildinfo.WithASCIILogo(logo),
		gobuildinfo.WithDisableRuntime(), // Comment this line to include runtime information
	)

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
}
