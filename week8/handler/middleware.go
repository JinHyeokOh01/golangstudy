package handler

import(
	"log"
	"net/http"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/auth"
)

func AuthMiddleware(j *auth.JWTer) func(next http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			req, err := j.FillContext(r)
			if err != nil{
				RespondJSON(r.Context(), w, ErrResponse{
					Message: "not find auth info",
					Details: []string{err.Error()},
				}, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func AdminMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Print("AdminMiddleware")
		if !auth.IsAdmin(r.Context()){
			RespondJSON(r.Context(), w, ErrResponse{
				Message: "not admin",
			}, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}