package kitt

import (
	"net/http"
)

func Run(r *Router) {
	server := &http.ServeMux{}

	// static files
	fileServer := http.FileServer(http.Dir("./public"))
	server.Handle("/css/", fileServer)
	server.Handle("/img/", fileServer)
	server.Handle("/js/", fileServer)

	// router
	MuxIt(r, server)

	http.ListenAndServe(":3000", server)
}
