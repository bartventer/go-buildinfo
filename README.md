# go-buildinfo

A lightweight Go module for capturing and displaying build-time and runtime metadata, including version, commit, build date, and custom application details.

## Installation

```bash
go get github.com/bartventer/go-buildinfo
```

## Usage
```go

package main

import (
    _ "embed"
    "fmt"
    "os"

    gobuildinfo "github.com/bartventer/go-buildinfo"
)

//go:embed logo.txt
var logo string

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
        gobuildinfo.WithProject(
            "BuildInfo Example", 
            "A simple example of using go-buildinfo", 
            "https://example.com",
        ),
        gobuildinfo.WithASCIILogo(logo),
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

Version: v1.0.0
Commit: d4c3db2e5f8a4b1e9f7c2a1b3d4e5f6a7b8c9d0e
Date: 2023-10-01
TreeState: clean
GoVersion: go1.20.3
Compiler: gc
Platform: linux/amd64
ModuleSum: h1:abc1234567890defghijklmnopqrstuvwxyz1234567890==
```

## License
This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Inspired by the [go-version](https://github.com/caarlos0/go-version) module.