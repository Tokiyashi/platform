package api

import (
	"log"
	"net/http"
	"platform/internal/db"
	"platform/internal/handlers"
	jwt_auth "platform/internal/jwt"
	"strings"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type API struct {
	db     *db.DB
	router *mux.Router
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		log.Printf("Raw Authorization header: '%s'", authHeader)

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authorization header is empty"))
			return
		}

		// Ensure Bearer prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = "Bearer " + authHeader
		}

		jwt := &jwt_auth.JWTAuth{}
		_, err := jwt.ValidateToken(strings.TrimPrefix(authHeader, "Bearer "))

		if err != nil {
			log.Printf("Token validation error: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Ошибка авторизации: " + err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (api *API) endpoints() {
	uh := handlers.NewUserHandler(api.db)
	ch := handlers.NewCourseHandler(api.db)

	// Public routes (no auth required)
	api.router.HandleFunc("/auth", uh.Auth).Methods("POST")
	api.router.HandleFunc("/users", uh.AddOne).Methods("POST") // registration
	api.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Protected routes (auth required)
	protected := api.router.PathPrefix("").Subrouter()
	protected.Use(middleware)

	// Course routes
	protected.HandleFunc("/courses", ch.GetCurrentCourses).Methods("GET")
	protected.HandleFunc("/courses", ch.AddOneCourse).Methods("POST")
	protected.HandleFunc("/courses/{id}", ch.UpdateCourse).Methods("PUT")
	protected.HandleFunc("/courses/{id}", ch.DeleteOneCourse).Methods("DELETE")

	// User routes
	protected.HandleFunc("/users/{id}", uh.GetUser).Methods("GET")
	protected.HandleFunc("/users/{id}/courses", uh.GetUserCourses).Methods("GET")
}

func (api *API) Start() error {
	return http.ListenAndServe(":8080", api.router)
}

func New(db *db.DB) *API {
	api := &API{
		db:     db,
		router: mux.NewRouter(),
	}
	api.endpoints()
	return api
}
