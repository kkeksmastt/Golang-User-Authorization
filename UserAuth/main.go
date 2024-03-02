package main

import (
	c "UserAuth/color"
	db "UserAuth/database"
	h "UserAuth/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(c.White + "██╗░░██╗███████╗██╗░░░░░██╗░░░░░░█████╗░  ███╗░░░███╗███████╗██████╗░░█████╗░██████╗░░██████╗" + c.Reset)
	fmt.Println(c.Red + "██║░░██║██╔════╝██║░░░░░██║░░░░░██╔══██╗  ████╗░████║██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔════╝" + c.Reset)
	fmt.Println(c.Green + "███████║█████╗░░██║░░░░░██║░░░░░██║░░██║  ██╔████╔██║█████╗░░██║░░██║██║░░██║██║░░██║╚█████╗░" + c.Reset)
	fmt.Println(c.Yellow + "██╔══██║██╔══╝░░██║░░░░░██║░░░░░██║░░██║  ██║╚██╔╝██║██╔══╝░░██║░░██║██║░░██║██║░░██║░╚═══██╗" + c.Reset)
	fmt.Println(c.Blue + "██║░░██║███████╗███████╗███████╗╚█████╔╝  ██║░╚═╝░██║███████╗██████╔╝╚█████╔╝██████╔╝██████╔╝" + c.Reset)
	fmt.Println(c.Purple + "╚═╝░░╚═╝╚══════╝╚══════╝╚══════╝░╚════╝░  ╚═╝░░░░░╚═╝╚══════╝╚═════╝░░╚════╝░╚═════╝░╚═════╝░" + c.Reset)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()

	r.Handle("/api/get-token", h.GetTokensHandler).Methods("POST")
	r.Handle("/api/refresh-token", h.RefreshTokenHandler).Methods("PUT")
	r.Handle("/", h.IndexHandler)

	http.ListenAndServe(":"+port, r)
}
