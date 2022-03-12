package main

import (
	"database/sql"
	"forum/model"
	"forum/pkg"
	"forum/routes"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", "sqlite/forum.db?_foreign_keys=on")
	if err != nil {
		log.Println(err)
		return
	}

	model.Db = db

	//creating tables in sqlite database
	if err = model.InitSQL(); err != nil {
		log.Println(err)
		return
	}
	// Creating a folder
	err = os.MkdirAll("./static/img_posts", os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
}

func init() {
	//loading config file
	pkg.LoadConfig()
	//opening file for logging
	file, err := os.OpenFile("forum.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	pkg.Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", routes.Home)

	mux.HandleFunc("/login", routes.Login)
	mux.HandleFunc("/logedin", routes.Logedin)
	mux.HandleFunc("/signup", routes.Signup)
	mux.HandleFunc("/signedup", routes.Signedup)
	mux.HandleFunc("/logout", routes.Logout)

	mux.HandleFunc("/title/", routes.Title)
	mux.Handle("/createpost", pkg.Middleware(http.HandlerFunc(routes.Createpost)))
	mux.Handle("/savepost", pkg.Middleware(http.HandlerFunc(routes.Savepost)))

	mux.Handle("/savecomment", pkg.Middleware(http.HandlerFunc(routes.Savecomment)))
	mux.Handle("/like", pkg.Middleware(http.HandlerFunc(routes.Like)))
	mux.Handle("/comlike", pkg.Middleware(http.HandlerFunc(routes.ComLike)))

	mux.HandleFunc("/cats/", routes.Cats)
	mux.HandleFunc("/liked", routes.Liked)
	mux.HandleFunc("/mine", routes.Mine)

	mux.HandleFunc("/deletePost", routes.DeletePost)
	mux.HandleFunc("/deleteComm", routes.DeleteComm)
	mux.HandleFunc("/editPost/", routes.EditPost)
	mux.HandleFunc("/editComm/", routes.EditComm)
	mux.HandleFunc("/editedPost/", routes.EditedPost)
	mux.HandleFunc("/editedComm/", routes.EditedComm)

	mux.HandleFunc("/notification", routes.Notification)

	mux.HandleFunc("/activity", routes.Activity)
	server := &http.Server{
		Addr:    pkg.Config.Address,
		Handler: mux,
	}

	log.Printf("Listening on %s port ...\n", server.Addr)
	log.Println(server.ListenAndServe())
}
