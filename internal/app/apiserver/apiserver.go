package apiserver

import (
	"fmt"
	"github.com/Kirnata/user_balance/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"path"
	"runtime"
)

// APIServer ...
type APIServer struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter() // router обрабатывает все входящие запросы и направляет их к нужным обработчикам
	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	//s.router.HandleFunc("/sinlePay", s.SinglePay())
	////s.router.HandleFunc("/p2pPay", s.handleHello())
	////s.router.HandleFunc("/balance", s.handleHello())
	////s.router.HandleFunc("/sinlePay", handlers.SinglePay).Methods("PUT")
	////s.router.HandleFunc("/sinlePay", handlers.P2PPay).Methods("PUT")
	////s.router.HandleFunc("/sinlePay", handlers.GetBalance).Methods("GET")
	s.router.HandleFunc("/singlePay", s.SinglePay).Methods("PUT")
	s.router.HandleFunc("/P2PPay", s.P2PPay).Methods("PUT")
	s.router.HandleFunc("/balance", s.GetBalance).Methods("GET")
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

//func (s *APIServer) SinglePay() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "SinglePay")
//	}
//}

func (s *APIServer) SinglePay(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "SinglePay")
}

func (s *APIServer) P2PPay(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "P2PPay")
}

func (s *APIServer) GetBalance(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Balance")
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	logrus.SetReportCaller(true)
	s.logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	s.logger.SetLevel(level)
	return nil
}
