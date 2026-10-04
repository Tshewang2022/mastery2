package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	// importing it for the side effects
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nado/chirpy/internal/database"
)

type config struct {
	addr           string
	fileserverHits atomic.Int32
	queries        *database.Queries
	jwtSecret      string
	chirpyKey      string
}

// Register function is getting confined to the return type of handle func
func (cfg *config) handle() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.Delete)
	mux.HandleFunc("POST /api/users", cfg.Register)
	mux.HandleFunc("POST /api/chirps", cfg.handlerChirpsCreate)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handlerGetChirpsByID)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", cfg.handlerChirpsDelete)
	mux.HandleFunc("GET /api/chirps", cfg.handlerChirpsGet)
	mux.HandleFunc("POST /api/login", cfg.handleLogin)
	mux.HandleFunc("POST /api/refresh", cfg.Refresh)
	mux.HandleFunc("POST /api/revoke", cfg.RevokeRefreshTokens)
	mux.HandleFunc("PUT /api/users", cfg.handlerUsersUpdate)
	mux.HandleFunc("POST /api/polka/webhooks", cfg.handlerPolkaWebhooks)

	mux.Handle("/assets", http.FileServer(http.Dir(".")))
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "text/plain")
		r.Header.Set("charset", "utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	})
	return mux
}

func (cfg *config) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *config) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>

		</html>
			`,
		cfg.fileserverHits.Load())))
}
func (cfg *config) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env: %v", err)
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Must set db url")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error starting database: %s", err)
	}
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Must set jwt secret")
	}

	chirpyKey := os.Getenv("POLKA_KEY")
	if chirpyKey == "" {
		log.Fatal("Must set chirpy key")
	}
	dbQueries := database.New(db)

	cfg := &config{
		addr:      ":8080",
		queries:   dbQueries,
		jwtSecret: secret,
		chirpyKey: chirpyKey,
	}

	// final code that runs is this one;
	s := &http.Server{
		Addr:           cfg.addr,
		Handler:        cfg.handle(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())

}
