package wasm

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
)

func DescendantsFunc(spec *placetypes.WOFPlacetypeSpecification) js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		var roles []string

		str_pt := args[0].String()

		if len(args) > 1 {
			str_roles := args[1].String()
			roles = strings.Split(str_roles, ",")
		}

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				pt, err := spec.GetPlacetypeByName(str_pt)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to create placetype, %v", err))
					return
				}

				var descendants []*placetypes.WOFPlacetype

				if len(roles) == 0 {
					descendants = spec.Descendants(pt)
				} else {
					descendants = spec.DescendantsForRoles(pt, roles)
				}

				enc_descendants, err := json.Marshal(descendants)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to marshal descendants, %v", err))
					return
				}

				resolve.Invoke(string(enc_descendants))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
