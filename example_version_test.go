package gdk_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/researchlab/gdk"
)

// build binary by makefile, command
// make default -f version.Makefile
// ExampleVersion  version examples
func ExampleVersion() {
	v := gdk.Version()
	// version by console cmd
	args := os.Args
	if len(args) >= 2 && args[1] == "version" {
		fmt.Println(v.String())
		return
	}
	// version by http request
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, v)
	})
	http.ListenAndServe(":8082", nil)
}
