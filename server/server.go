package server


import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/redis/go-redis/v9"
    "github.com/spf13/viper"
    "golang.org/x/exp/slog"
)

type Server struct {
	rotuer *mux.Router
	redis *redis.Redis
	logger *slog.Logger
	
}


type ResponseMsg struct {
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}

func (s *Server ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 	s.router.ServeHTTP(w, r)
}

func Run (envType *string){
	viper.SetConfigName("json")

	var level slog.Level
	if *envType == "dev" {
		viper.SetConfigName("dev-split")
		level = slog.LevelDebug
	} else {
		viper.SetConfigName("prod-split")
		level = slog.LevelInfo
	}

	viper.AddConfigPath("$HOME/.config")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("Error reading config file", err)
		return
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	logger := slog.New(handler)
	slog.SetDefault(logger)


	server := &Server{
		router: mux.NewRouter(),
		redis: redis.NewRedis(),
		logger: logger,
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
