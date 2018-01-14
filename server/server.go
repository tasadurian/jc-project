package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/thetommytwitch/jc-project/helpers"
	"github.com/thetommytwitch/jc-project/password"
	"github.com/thetommytwitch/jc-project/store"
)

// Server implements the http.Server
type Server struct {
	logger *log.Logger
	db     *store.Storer
	mux    *http.ServeMux
	sigs   chan os.Signal
}

// New returns a new server ...
func New(sigs chan os.Signal) *Server {
	s := &Server{mux: http.NewServeMux(), sigs: sigs}

	if s.logger == nil {
		s.logger = log.New(os.Stdout, "", 0)
	}

	s.db = store.NewStore()

	s.mux.HandleFunc("/", s.latencyMiddleware(s.index))
	s.mux.HandleFunc("/hash/", s.latencyMiddleware(s.hash))
	s.mux.HandleFunc("/shutdown/", s.latencyMiddleware(s.shutdown))
	s.mux.HandleFunc("/stats/", s.stats)

	return s
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func (s *Server) hash(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	case "GET":
		key, err := helpers.ParseURL(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		hash, err := s.db.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		w.Write([]byte(hash))

	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		str, err := helpers.GetPasswordString(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		key := rand.Intn(100)

		go func(key int, db *store.Storer) {
			time.Sleep(5 * time.Second)
			encoded := password.EncodeAndHash(str)
			db.Put(key, encoded)
		}(key, s.db)

		w.Write([]byte(fmt.Sprintf("%d", key)))

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (s *Server) shutdown(w http.ResponseWriter, r *http.Request) {
	id := os.Getpid()
	p, err := os.FindProcess(id)
	if err != nil {
		s.logger.Fatal(err)
	}

	p.Signal(os.Interrupt)
	w.Write([]byte("Shutdown process initiated..."))
}

func (s *Server) stats(w http.ResponseWriter, r *http.Request) {
	latency, err := s.db.GetAvgLatency()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseString := fmt.Sprintf("{\"total\": %d, \"average\": %f}",
		s.db.GetReqCount(), latency)

	w.Write([]byte(responseString))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) latencyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		since := time.Since(now)
		s.db.PutLatency(since.Seconds())
		s.logger.Printf("%s took %f seconds", r.URL.Path, since.Seconds())
	})
}
