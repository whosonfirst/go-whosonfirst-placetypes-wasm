package main

import (
	"log"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"github.com/whosonfirst/go-whosonfirst-placetypes-wasm"
)

func main() {

	spec, err := placetypes.DefaultWOFPlacetypeSpecification()

	if err != nil {
		log.Fatalf("Failed to load WOF placetype specification, %v", err)
	}

	descendants_func := wasm.DescendantsFunc(spec)
	defer descendants_func.Release()

	ancestors_func := wasm.AncestorsFunc(spec)
	defer ancestors_func.Release()

	placetypes_func := wasm.PlacetypesFunc(spec)
	defer placetypes_func.Release()

	parents_func := wasm.ParentsFunc(spec)
	defer parents_func.Release()

	children_func := wasm.ChildrenFunc(spec)
	defer children_func.Release()

	isvalid_func := wasm.IsValidFunc(spec)
	defer isvalid_func.Release()

	iscore_func := wasm.IsCoreFunc(spec)
	defer iscore_func.Release()

	js.Global().Set("whosonfirst_placetypes_descendants", descendants_func)
	js.Global().Set("whosonfirst_placetypes_ancestors", ancestors_func)
	js.Global().Set("whosonfirst_placetypes_children", children_func)
	js.Global().Set("whosonfirst_placetypes_parents", parents_func)

	js.Global().Set("whosonfirst_placetypes_is_valid", isvalid_func)
	js.Global().Set("whosonfirst_placetypes_is_core", iscore_func)
	js.Global().Set("whosonfirst_placetypes", placetypes_func)

	c := make(chan struct{}, 0)

	log.Println("Who's On First placetypes WASM binary initialized")
	<-c
}
