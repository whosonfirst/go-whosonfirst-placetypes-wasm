package main

import (
	"log"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes-wasm"
)

func main() {

	descendants_func := wasm.DescendantsFunc()
	defer descendants_func.Release()

	ancestors_func := wasm.AncestorsFunc()
	defer ancestors_func.Release()

	placetypes_func := wasm.PlacetypesFunc()
	defer placetypes_func.Release()

	isvalid_func := wasm.IsValidFunc()
	defer isvalid_func.Release()

	js.Global().Set("whosonfirst_placetypes_descendants", descendants_func)
	js.Global().Set("whosonfirst_placetypes_ancestors", ancestors_func)
	js.Global().Set("whosonfirst_placetypes_is_valid", isvalid_func)
	js.Global().Set("whosonfirst_placetypes", placetypes_func)

	c := make(chan struct{}, 0)

	log.Println("Who's On First placetypes WASM binary initialized")
	<-c
}
