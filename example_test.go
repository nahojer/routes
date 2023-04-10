package routes_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/nahojer/routes"
)

func Example() {
	type ctxKey string

	rt := routes.NewTrie[http.Handler]()

	handlePing := func(w http.ResponseWriter, r *http.Request) {
		pong := r.Context().Value(ctxKey("pong")).(string)
		w.WriteHeader(200)
		fmt.Fprint(w, pong)
	}
	rt.Add(http.MethodGet, "/ping/:pong", http.HandlerFunc(handlePing))

	req := httptest.NewRequest(http.MethodGet, "http://localhost/ping/"+url.PathEscape("It's-a me, Mario!"), nil)
	w := httptest.NewRecorder()

	h, params, found := rt.Lookup(req)
	if !found {
		panic("never reached")
	}

	req = req.WithContext(context.WithValue(req.Context(), ctxKey("pong"), params["pong"]))
	h.ServeHTTP(w, req)

	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Body: %q\n", w.Body.String())
	// Output:
	// Status: 200
	// Body: "It's-a me, Mario!"
}
