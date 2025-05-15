package gobuildinfo_test

import (
	"fmt"

	gobuildinfo "github.com/bartventer/go-buildinfo"
)

const logo = `
_|                  _|  _|        _|  _|                _|_|
_|_|_|    _|    _|      _|    _|_|_|      _|_|_|      _|        _|_|
_|    _|  _|    _|  _|  _|  _|    _|  _|  _|    _|  _|_|_|_|  _|    _|
_|    _|  _|    _|  _|  _|  _|    _|  _|  _|    _|    _|      _|    _|
_|_|_|      _|_|_|  _|  _|    _|_|_|  _|  _|    _|    _|        _|_|
`

func ExampleInfo_String() {
	info := gobuildinfo.New(
		gobuildinfo.WithVersion("v1.0.0"),
		gobuildinfo.WithCommit("a696fbbcb8ae009e3f88df2d7b00c09bea903c9e"),
		gobuildinfo.WithDate("2023-01-01"),
		gobuildinfo.WithTreeState("clean"),
		gobuildinfo.WithDisableRuntime(), // Comment this line to include runtime info
		gobuildinfo.WithProject(gobuildinfo.Project{
			Name:        "MyApp",
			Description: "This is a sample app",
			URL:         "https://example.com",
			ASCIILogo:   logo,
		}),
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
	// Version:    v1.0.0
	// Commit:     a696fbbcb8ae009e3f88df2d7b00c09bea903c9e
	// Date:       2023-01-01
	// TreeState:  clean
}
