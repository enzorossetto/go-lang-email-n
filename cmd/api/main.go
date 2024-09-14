package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type product struct {
	ID   int
	Name string
}

type myHandler struct{}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler"))
}

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(myMiddleware)

	m := myHandler{}
	r.Handle("/myHandler", m)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")

		println("root GET endpoint")

		w.Write([]byte(param))
	})

	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("middleware before")
		next.ServeHTTP(w, r)
		println("middleware after next")
	})
}
