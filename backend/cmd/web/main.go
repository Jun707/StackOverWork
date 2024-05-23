package main

import "database/sql"

// "stack/web/backend/internal/users"

// "github.com/gin-gonic/gin"

func main() {
	// router := gin.Default()
	// router.GET("/users", users.GetUsers)
	db, err := sql.Open("mysql", "username:password@tcp(job-database1.cipitaoficoy.us-east-1.rds.amazonaws.com)")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()


	// router.Run(":3000")
	// mux := http.NewServeMux()

	// fileServer := http.FileServer(http.Dir("frontend/build"))
    // mux.Handle("/static/", fileServer)
	// mux.HandleFunc("GET /{$}", server.Server)

	// log.Print("starting server on :3000")

	// err := http.ListenAndServe(":3000", mux)
	// log.Fatal(err)
}