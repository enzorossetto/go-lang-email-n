package endpoints

import (
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]interface{}{"error": "request does not contain an Authorization header"})
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/emailn_provider")

		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]interface{}{"error": "error connecting to the identity provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "emailn"})
		_, err = verifier.Verify(r.Context(), token)

		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]interface{}{"error": "invalid token"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
