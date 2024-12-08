package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Abhinav7903/split/db/postgres"
	"github.com/Abhinav7903/split/db/redis"
	"github.com/Abhinav7903/split/pkg/groups"
	"github.com/Abhinav7903/split/pkg/mail"
	"github.com/Abhinav7903/split/pkg/sessmanager"
	"github.com/Abhinav7903/split/pkg/users"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

type Server struct {
	router      *mux.Router
	redis       *redis.Redis
	logger      *slog.Logger
	user        users.Repository
	sessmanager sessmanager.Repository
	mail        mail.Repository
	group       groups.Repository
}

type ResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func Run(envType *string) {
	viper.SetConfigName("json")

	var level slog.Level
	if *envType == "dev" {
		viper.SetConfigName("dev-split")
		level = slog.LevelDebug
	} else {
		viper.SetConfigName("prod-split")
		level = slog.LevelInfo
	}

	viper.AddConfigPath("$HOME/.split")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("Error reading config file", err)
		return
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// db
	postgres := postgres.NewPostgres()
	redis := redis.NewRedis(envType)
	server := &Server{
		router:      mux.NewRouter(),
		redis:       redis,
		logger:      logger,
		user:        postgres,
		sessmanager: redis,
		mail: mail.NewMail(
			viper.GetString("mail_id"),
			viper.GetString("mail_pass"),
			viper.GetString("app-pass"),
		),
		group: postgres,
	}

	server.RegisterRoutes()
	port := ":8080"
	if *envType != "dev" {
		port = ":8194"
	}
	server.logger.Info("Starting server", "mode", *envType, "port", port)

	if err := http.ListenAndServe(port, server); err != nil {
		server.logger.Error("Server failed to start", "error", err)
	}

}

func (s *Server) respond(
	w http.ResponseWriter,
	data interface{},
	status int,
	err error,
) {
	// Set content type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err == nil {
		resp := &ResponseMsg{
			Message: "success",
			Data:    data,
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			s.logger.Error("Error in encoding the response", "error", err)
			return
		}
		return
	}
	resp := &ResponseMsg{
		Message: err.Error(),
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("Error in encoding the response", "error", err)
		return
	}
}
