package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
)

func ParentsFunc(spec *placetypes.WOFPlacetypeSpecification) js.Func {

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

				parents := make([]*placetypes.WOFPlacetype, len(pt.Parent))

				for idx, pid := range pt.Parent {

					p, err := spec.GetPlacetypeById(pid)

					if err != nil {
						reject.Invoke(fmt.Sprintf("Failed to retrieve placetype with ID %d, %v", pid, err))
						return
					}

					parents[idx] = p
				}

				enc_pt, err := json.Marshal(parents)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to marshal parent placetypes, %v", err))
					return
				}

				resolve.Invoke(string(enc_pt))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
