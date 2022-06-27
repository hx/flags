package hids

import (
	"crypto/subtle"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hx/flags/actions"
	"github.com/hx/flags/states"
	"net/http"
	"strconv"
)

type HttpServer struct {
	state      states.State
	server     http.Server
	passphrase string
	router     *mux.Router
	perform    actions.Performer
}

func NewHttpServer(address, passphrase string) *HttpServer {
	server := &HttpServer{
		server:     http.Server{Addr: address},
		passphrase: passphrase,
		router:     mux.NewRouter(),
	}
	server.server.Handler = server

	server.router.HandleFunc("/flags", server.flags).Methods("GET")
	server.router.HandleFunc("/toggle/{id:\\d|[1-5]\\d|6[0-3]}", server.toggle).Methods("POST")

	server.router.PathPrefix("/").HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, "🥺", 404)
	})

	return server
}

func (h *HttpServer) Update(diff states.Diff) {
	h.state = h.state.Apply(diff)
}

func (h *HttpServer) Listen(perform actions.Performer) error {
	h.perform = perform
	return h.server.ListenAndServe()
}

func (h *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if h.passphrase != "" && subtle.ConstantTimeCompare(
		[]byte(request.Header.Get("Authorization")),
		[]byte(h.passphrase),
	) == 0 {
		http.Error(writer, "Invalid passphrase", 401)
		return
	}
	h.router.ServeHTTP(writer, request)
}

func (h *HttpServer) toggle(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	<-h.perform(actions.Toggle(id))
	h.flags(writer, request)
}

func (h *HttpServer) flags(writer http.ResponseWriter, _ *http.Request) {
	type flag struct {
		Index int  `json:"index"`
		State bool `json:"state"`
	}
	length := h.state.Len()
	flags := make([]flag, length)
	for i := range flags {
		flags[i] = flag{Index: i, State: h.state.Get(i)}
	}
	body, err := json.Marshal(map[string][]flag{"flags": flags})
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(append(body, '\n'))
}
