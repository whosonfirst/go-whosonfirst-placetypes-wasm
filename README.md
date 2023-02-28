# go-whosonfirst-placetypes-wasm

Go package for compiling methods from the [go-whosonfirst-placetypes](https://github.com/whosonfirst/go-whosonfirst-placetypes) package to a JavaScript-compatible WebAssembly (wasm) binary. It also provides a net/http middleware packages for appending the necessary static assets and HTML resources to use the wasm binary in web applications.

## Build

To build the `validate_feature` WebAssembly binary for use in your applications run the following command:

```
GOOS=js GOARCH=wasm go build -mod vendor -ldflags="-s -w" -o static/wasm/whosonfirst_placetypes.wasm cmd/placetypes/main.go
```

## Use

In order to load the `whosonfirst_placetypes` functions you will need to include the `wasm_exec.js` and `whosonfirst.placetypes.wasm.js` JavaScript files, or functional equivalents. Both scripts are bundled with this package in the [static/javascript](static/javascript) folder.

## Functions

### whosonfirst_placetypes

Return a JSON-encoded list of all valid Who's On First placetypes.

```
	whosonfirst_placetypes()
	    .then((data) => { ... });
```

### whosonfirst_placetypes_ancestors

Return a JSON-encoded list of all the ancestors for a given placetype.

```
	whosonfirst_placetypes_ancestors("region", "common,optional,common_optional")
	    .then((data) => { ... });
```

### whosonfirst_placetypes_descendants

Return a JSON-encoded list of all the descendants for a given placetype.

```
	whosonfirst_placetypes_ancestors("descendants", "common,optional,common_optional")
	    .then((data) => { ... });
```

## Middleware

The `go-whosonfirst-validate-wasm/http` package provides methods for appending static assets and HTML resources to existing web applications to facilitate the use of the `validate_feature` WebAssembly binary. For example:

```
package main

import (
	"embed"
	"log"
	"net/http"

	wasm "github.com/whosonfirst/go-whosonfirst-placetypes-wasm/http"
)

//go:embed index.html example.*
var FS embed.FS

func main() {

	mux := http.NewServeMux()

	wasm.AppendAssetHandlers(mux)

	http_fs := http.FS(FS)
	example_handler := http.FileServer(http_fs)

	wasm_opts := wasm.DefaultWASMOptions()
	wasm_opts.EnableWASMExec()

	example_handler = wasm.AppendResourcesHandler(example_handler, wasm_opts)

	mux.Handle("/", example_handler)

	addr := "localhost:8080"
	log.Printf("Listening for requests on %s\n", addr)

	http.ListenAndServe(addr, mux)
}
```

_Error handling omitted for brevity._

## Example

There is a full working example of this application in the `cmd/example` folder. To run this application type the following command:

```
$> make example
go run -mod vendor cmd/example/main.go
2023/01/31 15:11:48 Listening for requests on localhost:8080
```

Then open `http://localhost:8080` in a  web browser. You should see something like this:

![](docs/images/wof-placetypes-wasm.png)

For example:

![](docs/images/wof-placetypes-wasm.gif)

## See also

* https://github.com/whosonfirst/go-whosonfirst-placetypes