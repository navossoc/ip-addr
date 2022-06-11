package main

import (
	"fmt"
	"net/http"
	"sort"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "RemoteAddr: %s\n\n", r.RemoteAddr)

		keys := make([]string, 0, len(r.Header))
		for k := range r.Header {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := r.Header.Get(k)
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	})

	http.ListenAndServe(":80", nil)
}
