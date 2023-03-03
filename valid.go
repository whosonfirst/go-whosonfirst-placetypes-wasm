package wasm

import (
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
)

func IsValidFunc(spec *placetypes.WOFPlacetypeSpecification) js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		str_pt := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				is_valid := spec.IsValidPlacetype(str_pt)

				if !is_valid {
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
