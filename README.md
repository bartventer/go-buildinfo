[![Go Reference](https://pkg.go.dev/badge/github.com/bartventer/go-buildinfo.svg)](https://pkg.go.dev/github.com/bartventer/go-buildinfo)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartventer/go-buildinfo)](https://goreportcard.com/report/github.com/bartventer/go-buildinfo)
[![Test](https://github.com/bartventer/go-buildinfo/actions/workflows/default.yml/badge.svg)](https://github.com/bartventer/go-buildinfo/actions/workflows/default.yml)
[![codecov](https://codecov.io/github/bartventer/go-buildinfo/graph/badge.svg?token=xkncFVoRsX)](https://codecov.io/github/bartventer/go-buildinfo)

# go-buildinfo

A zero-dependency Go module for capturing and displaying build-time and runtime metadata.

## Installation

```bash
go get github.com/bartventer/go-buildinfo
```

## Usage
```go

package main

import (
    "fmt"

    gobuildinfo "github.com/bartventer/go-buildinfo"
)

const logo = `
    __          _ __    ___       ____    
   / /_  __  __(_) /___/ (_)___  / __/___ 
  / __ \/ / / / / / __  / / __ \/ /_/ __ \
 / /_/ / /_/ / / / /_/ / / / / / __/ /_/ /
/_.___/\__,_/_/_/\__,_/_/_/ /_/_/  \____/
`

// Typically set by the linker during the build process
var (
    version   = "v1.0.0"
    commit    = "d4c3db2e5f8a4b1e9f7c2a1b3d4e5f6a7b8c9d0e"
    date      = "2025-10-01"
    treeState = "clean"
)

func main() {
    info := gobuildinfo.New(
        gobuildinfo.WithVersion(version),
        gobuildinfo.WithCommit(commit),
        gobuildinfo.WithDate(date),
        gobuildinfo.WithTreeState(treeState),
        gobuildinfo.WithProject(gobuildinfo.Project{
            Name:        "BuildInfo Example",
            Description: "A simple example of using go-buildinfo",
            URL:         "https://example.com",
            ASCIILogo:   logo,
        }),
    )

    fmt.Println(info.String())
}
```

## Example Output

```plaintext
    __          _ __    ___       ____    
   / /_  __  __(_) /___/ (_)___  / __/___ 
  / __ \/ / / / / / __  / / __ \/ /_/ __ \
 / /_/ / /_/ / / / /_/ / / / / / __/ /_/ /
/_.___/\__,_/_/_/\__,_/_/_/ /_/_/  \____/ 
BuildInfo Example: A simple example of using go-buildinfo
https://example.com

Version:    v1.0.0
Commit:     d4c3db2e5f8a4b1e9f7c2a1b3d4e5f6a7b8c9d0e
Date:       2023-10-01
TreeState:  clean
GoVersion:  go1.20.3
Compiler:   gc
Platform:   linux/amd64
ModuleSum:  h1:abc1234567890defghijklmnopqrstuvwxyz1234567890==
```

## License
This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Inspired by the [go-version](https://github.com/caarlos0/go-version) module.