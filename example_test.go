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
	type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)

	rt := routes.NewTrie[Handler]()

	rt.Add(http.MethodGet, "/ping/:pong", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		pong := ctx.Value("pong").(string)
		w.WriteHeader(200)
		fmt.Fprint(w, pong)
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost/ping/"+url.PathEscape("It's-a me, Mario!"), nil)
	w := httptest.NewRecorder()

	h, params, found := rt.Lookup(req)
	if !found {
		panic("never reached")
	}

	ctx := context.WithValue(req.Context(), "pong", params["pong"])
	h(ctx, w, req)

	fmt.Printf("Status: %d\n", w.Code)
	fmt.Printf("Body: %q\n", w.Body.String())
	// Output:
	// Status: 200
	// Body: "It's-a me, Mario!"
}
