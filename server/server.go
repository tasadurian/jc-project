package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/thetommytwitch/jc-project/password"
)

// Server implements the http.Server
type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
	sigs   chan os.Signal
}

// New returns a new server ...
func New(sigs chan os.Signal) *Server {
	s := &Server{mux: http.NewServeMux(), sigs: sigs}

	if s.logger == nil {
		s.logger = log.New(os.Stdout, "", 0)
	}

	s.mux.HandleFunc("/", s.index)
	s.mux.HandleFunc("/hash/", s.shutdown)
	s.mux.HandleFunc("/shutdown/", s.shutdown)

	return s
}

// Logger ...
func Logger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func (s *Server) hash(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// write error
	}

	str := strings.Split(string(b), "=")
	if len(str) != 2 {
		// error
	}

	encoded := password.EncodeAndHash(str[1])

	time.Sleep(5 * time.Second)

	w.Write([]byte(encoded))
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

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "example Go server")

	s.mux.ServeHTTP(w, r)
}
