package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func indexResponse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintf(w, "<html>\n<body>\n<h1>fakeservice</h1>\n</body>\n</html>\n")
}

func main() {
	fmt.Printf("Starting fakeservice on localhost:9000...\n")
	router := httprouter.New()
	router.GET("/", indexResponse)

	// Start me up!
	http.ListenAndServe(":9000", router)
}
