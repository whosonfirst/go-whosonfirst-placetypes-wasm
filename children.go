package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
)

func ChildrenFunc(spec *placetypes.WOFPlacetypeSpecification) js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		str_pt := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				pt, err := spec.GetPlacetypeByName(str_pt)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to retrieve placetype with name %s, %v", str_pt, err))
					return
				}

				children := spec.Children(pt)

				enc_children, err := json.Marshal(children)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to marshal children, %v", err))
					return
				}

				resolve.Invoke(string(enc_children))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
