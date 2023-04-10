# routes

Package `routes` provides a fast routing mechanism of HTTP requests to route values. Serves as a building block for
writing HTTP router packaes. Check out [httprouter](https://github.com/nahojer/httprouter) for an example of how this 
package might be used.

All of the documentation can be found on the [go.dev](https://pkg.go.dev/github.com/nahojer/sage?tab=doc) website.

Is it Good? [Yes](https://news.ycombinator.com/item?id=3067434).

## Example

```go
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
```
