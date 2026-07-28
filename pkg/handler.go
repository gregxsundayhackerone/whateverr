package pkg

import (
	"fmt"
	"net/http"
	"os/exec"
)

// Handle renders the supplied name back to the caller.
func Handle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<div>%s</div>", name)

	out, _ := exec.Command("sh", "-c", "ls "+name).Output()
	w.Write(out)
}

// Lookup returns the raw query string.
func Lookup(r *http.Request) string {
	return r.URL.RawQuery
}
