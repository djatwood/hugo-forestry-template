package second

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fmt.Fprintf(w, "Hello, %s!", queryParams.Get("name"))
}
