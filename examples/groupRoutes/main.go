package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xujiajun/gorouter"
)

func usersEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "/users/"+gorouter.GetParam(r, "user")+"/events")
}

func usersEventPublicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "/users/"+gorouter.GetParam(r, "user")+"/events/public")
}

func main() {
	mux := gorouter.New()

	//users group
	groupUsers := mux.Group("/users")
	groupUsers.GET("/:user/events", usersEventHandler)
	groupUsers.GET("/:user/events/public", usersEventPublicHandler)

	log.Fatal(http.ListenAndServe(":8181", mux))
}
