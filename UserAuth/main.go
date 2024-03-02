package main

import (
	c "UserAuth/color"
	db "UserAuth/database"
	h "UserAuth/handlers"
	sl "UserAuth/serverLog"
	"fmt"
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

	port := ""
	if err := godotenv.Load(".env"); err != nil {
		sl.ErrorLog(err)
	} else {
		port = os.Getenv("PORT")
	}

	db.InitDB()

	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()

	r.Handle("/api/get-token", h.GetTokensHandler).Methods("POST")
	r.Handle("/api/refresh-token", h.RefreshTokenHandler).Methods("PUT")
	r.Handle("/", h.IndexHandler)

	http.ListenAndServe(":"+port, r)
}
