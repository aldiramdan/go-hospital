package configs

import (
	"log"
	"net/http"
	"os"

	"github.com/aldiramdan/hospital/databases/db"
	"github.com/aldiramdan/hospital/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start application",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	mux := http.NewServeMux()

	db := db.ConnectDB()
	routers.IndexRoute(mux, db)

	var address string = "0.0.0.0:8080"
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "0.0.0.0:" + PORT
	}

	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Println("Server running on", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

	return server.ListenAndServe()
}
