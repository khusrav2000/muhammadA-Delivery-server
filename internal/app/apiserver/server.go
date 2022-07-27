package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/model"
	"github.com/khusrav2000/muhammadA-Delivery-server/internal/app/store"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "MuhammadA"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectLoginOrPassword = errors.New("incorect login or password")
	errNotAuthenticated         = errors.New("not authenticated")
	errNoPermission             = errors.New("no permission")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/login", s.handleSessionsCreate()).Methods("POST")

	// private need authenticate
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")

	// users
	private.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	private.HandleFunc("/myprofile", s.handleMyProfile()).Methods("GET")

	// Point (Pharmacies)
	private.HandleFunc("/pharmacies", s.handlePharmacyCreate()).Methods("POST")      // add new point (pharmacy)
	private.HandleFunc("/pharmacies", s.handlePharmacies()).Methods("GET")           // get all points by user access
	private.HandleFunc("/pharmacies/{id}", s.handleUpdatePharmacy()).Methods("POST") // update pharmacy :id

	/*
		// users
		private.HandleFunc("/users/id", s.handleUsersGetByID()).Methods("GET") // get user info
		private.HandleFunc("/users/id", s.handleUsersUpdate()).Methods("POST") // update user

		// Points (Pharmacies)
		private.HandleFunc("/points/id", s.handlePointsGetByID()).Methods("GET")       // get point info by ID
		private.HandleFunc("/points/id", s.handlePointsDeleteByID()).Methods("DELETE") // delete point
		private.HandleFunc("/points/search", s.handlePointsSearch()).Methods("GET")    // search points and get list point

		// Products
		private.HandleFunc("/products", s.handleProductsCreate()).Methods("POST")          // add new product
		private.HandleFunc("/products", s.handleProducts()).Methods("GET")                 // get all products in which there is access
		private.HandleFunc("/products/id", s.handleProductsGetByID()).Methods("GET")       // get product info by id
		private.HandleFunc("/products/id", s.handleProductsDeleteByID()).Methods("DELETE") // delete product
		private.HandleFunc("/products/search", s.handleProductsSearch()).Methods("GET")    // search product by ...

		// Orders
		private.HandleFunc("/orders", s.handleOrdersCreate()).Methods("POST")              // create new order
		private.HandleFunc("/orders", s.handleOrders()).Methods("GET")                     // get all orders in which there is access
		private.HandleFunc("/orders/id", s.handleOrdersGetByID()).Methods("GET")           // get order info with ID
		private.HandleFunc("/orders/id/delivered", s.handleOrdersDelivered()).Methods("?") // make order delivered
		private.HandleFunc("/orders/filter", s.handleOrdersFilter()).Methods("GET")        // get orders by filter
	*/

}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-REQUEST-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"compleated with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		id, ok := session.Values["user_id"]
		log.Println(ok)
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Role     string `json:"role"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		log.Println("START CREATE USER!")
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		user := r.Context().Value(ctxKeyUser).(*model.User)
		log.Println(user)
		hasAccess, err := s.store.User().CheckAccessFor(user, "create_users")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		if !hasAccess {
			s.error(w, r, http.StatusBadRequest, errNoPermission)
			return
		}
		u := &model.User{
			Login:    req.Login,
			Password: req.Password,
			Role:     req.Role,
			Name:     req.Name,
			Surname:  req.Surname,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().FindByLogin(req.Login)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectLoginOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleMyProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxKeyUser).(*model.User)

		if err := s.store.User().GetProfile(user); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, user)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
