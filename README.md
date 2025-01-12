# `svgpath`

As the name implies, `svgpath` is a tool that specialilzes in creating SVG path data, enabling the integration of smooth geopetric shapes into web-based graphics or SVG files.

## Usage

`svgpath` is currently only offered as a Go package, but a command-line interface and web API are under development.

#### Adding `svgpath` to Your Go Project

To use `svgpath` in your Go project, simply import the package:

```bash
go get github.com/bsalbilla06/svgpath
```

**Example usage**

```go
package main

import (
    "fmt"
    "log"

    "github.com/bsalbilla06/svgpath"
)

func main() {
    polyline := "_p~iF~ps|U_ulLnnqC_mqNvxq`@"

    svgPath, err := svgpath.GeneratePath(polyline)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Generated SVG Path:", svgPath)
}
```