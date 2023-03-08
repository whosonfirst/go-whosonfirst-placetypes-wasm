package wasm

import (
	"fmt"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
)

func IsCoreFunc(spec *placetypes.WOFPlacetypeSpecification) js.Func {

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

				if !pt.IsCorePlacetype() {
					reject.Invoke()
					return
				}

				resolve.Invoke()
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
