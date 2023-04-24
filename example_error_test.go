package gdk_test

import (
	"encoding/json"
	"fmt"

	"github.com/researchlab/gdk"
)

func jsonMarshal(size int) error {
	_, err := json.Marshal(make(chan struct{}, size))
	if err != nil {
		return gdk.ErrorCause(err)
	}
	return nil
}

// deal with the given error
func ExampleErrorCause() {
	const ERR_MARSHAL_FAILED = 1000
	var size = 10
	var e gdk.Err
	err := jsonMarshal(size)

	if err != nil {
		e = gdk.ErrorCause(err).WithCode(ERR_MARSHAL_FAILED).WithTag("ParseError").
			WithFields(map[string]interface{}{
				"inputs": map[string]int{
					"size": 10,
				},
			})
	}
	if e != nil {
		// e not equal nil
	}
	errTxt := e.DetailText()
	fmt.Println(errTxt)
	// Output:
	// CallChains=ExampleErrorCause.jsonMarshal, Tag=ParseError, Fields={"inputs":{"size":10}}, Code=1000, Error=json: unsupported type: chan struct {}
}

// new error with format
func ExampleErrorf() {
	const MIN_VALUE = 5
	const ERR_PARAMS_INVALID = "ERR PARAMS INVALID"
	err := gdk.Errorf("params invalid, size need > %d", MIN_VALUE).WithCode(ERR_PARAMS_INVALID)
	fmt.Println(err.DetailText())
	//Output:
	// CallChains=ExampleErrorf, Code=ERR PARAMS INVALID, Error=params invalid, size need > 5
}

// use error Templates
func ExampleErrorT() {
	const (
		ERR_PARAMS_INVALID = "PARAMS INVALID"
	)
	var errorTemplates = map[any]string{
		ERR_PARAMS_INVALID: "params invalid, include(%v)",
	}
	gdk.SetGlobalErrorTemplates(errorTemplates)
	gdk.SetGlobalTag("ip:192.168.190.70")
	gdk.SetGlobalFields(map[string]interface{}{
		"service": "timer-job-3",
	})

	err := gdk.ErrorT(ERR_PARAMS_INVALID, "size need > 5, code need > 0")
	fmt.Println(err.DetailText())
	// Output:
	// CallChains=ExampleErrorT, GlobalFields={"service":"timer-job-3"}, Code=PARAMS INVALID, Error=params invalid, include(size need > 5, code need > 0)
}
