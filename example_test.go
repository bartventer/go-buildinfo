package gobuildinfo

import (
	"fmt"
)

func ExampleInfo_String() {
	info := New(
		WithVersion("v1.0.0"),
		WithCommit("abc123"),
		WithDate("2023-01-01"),
		WithTreeState("clean"),
		WithProject("MyApp", "This is a sample app", "https://example.com"),
		WithASCIILogo("<<<ASCII Art>>>"),
	)
	info.runtime.GoVersion = "go1.24.2"

	fmt.Println(info.String())
	// Output:
	// <<<ASCII Art>>>
	// MyApp: This is a sample app
	// https://example.com
	//
	// Version    v1.0.0
	// Commit     abc123
	// Date       2023-01-01
	// TreeState  clean
	// GoVersion  go1.24.2
	// Compiler   gc
	// Platform   linux/amd64
	// ModuleSum  none
}
