package main

func main(){
	envType := flag.String("env", "dev", "Environment type production or development")
	flag.Parse()
	slog.Info("Environment type: ", *envType)
	slog.Info("Starting server...")
	server.Run(*envType)
}