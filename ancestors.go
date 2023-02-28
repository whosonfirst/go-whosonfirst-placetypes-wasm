package wasm

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
)

func AncestorsFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		str_pt := args[0].String()
		str_roles := args[1].String()

		roles := strings.Split(str_roles, ",")

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				pt, err := placetypes.GetPlacetypeByName(str_pt)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to create placetype, %v", err))
					return
				}

				ancestors := placetypes.AncestorsForRoles(pt, roles)

				enc_ancestors, err := json.Marshal(ancestors)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to marshal ancestors, %v", err))
					return
				}

				resolve.Invoke(string(enc_ancestors))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
