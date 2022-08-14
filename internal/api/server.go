package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"os"
	"wb_l0/internal"
	"wb_l0/internal/repo"
)

const imagePath = "/home/vlad/GolandProjects/wb_l0/assets/frank.JPG"

type Server struct {
	*mux.Router
}

func NewServer(cacheData *repo.Repository) *Server {
	s := &Server{Router: mux.NewRouter()}
	s.routes(cacheData)
	return s
}

func (s *Server) routes(cacheData *repo.Repository) {
	s.HandleFunc("/display-data", s.displayHandler(cacheData)).Methods("GET")
	s.HandleFunc("/", s.indexHandler()).Methods("GET")
}

func (s *Server) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if buf, err := os.ReadFile(imagePath); err != nil {
			panic(err)
		} else {
			w.Header().Set("Content-Type", "image/png")
			w.Write(buf)
		}
	}
}

func (s *Server) displayHandler(cacheData *repo.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uri, err := url.ParseRequestURI(r.RequestURI)
		if err != nil {
			panic(err)
		}

		requestedOrderUid := uri.Query()["order-uid"][0]
		if data, err := (*cacheData).GetById(requestedOrderUid); err != nil {
			errorMsg := fmt.Errorf("getting by id: %w", err).Error()
			w.Write([]byte(errorMsg))
		} else {
			displayedData := internal.MapStoredToDisplayed(data)
			dataByte, err := json.MarshalIndent(displayedData, "", "    ")
			if err != nil {
				panic(err)
			}
			w.Write(dataByte)
		}
	}
}
