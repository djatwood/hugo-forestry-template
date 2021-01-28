package currentTime

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintln(w, "Hello, World!")
	fmt.Fprintln(w, currentTime)
}
