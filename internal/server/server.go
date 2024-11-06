package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
)

type Server struct {
	wg sync.WaitGroup
}

func (s *Server) StartServer(port int, htmlContent string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := w.Write([]byte(htmlContent))
		if err != nil {
			http.Error(w, "Error sending HTML", http.StatusInternalServerError)
			return
		}
	})

	address := fmt.Sprintf(":%d", port)
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := http.ListenAndServe(address, nil); err != nil {
			fmt.Fprintf(os.Stderr, "Server startup error: %v\n", err)
		}
	}()

	fmt.Printf("Server started on port %d\n", port)
	fmt.Printf("http://localhost:%d\n", port)
	return nil
}

func FindFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, fmt.Errorf("failed to find a free port: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	return port, nil
}

func (s *Server) Wait() {
	s.wg.Wait()
}
