package main

import (
	"github.com/upalchowdhury/dist-service/internal/small_test"
	//"github.com/upalchowdhury/dist-service/internal/server"
)

func main() {
	// srv := server.NewHTTPServer(":8080")
	// log.Fatal(srv.ListenAndServe())
	small_test.TestSegment()
}
