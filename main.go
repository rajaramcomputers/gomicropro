package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", myHandler("Customer New Handler"))

	s := http.Server{
		Addr: ":3000",
	}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started and listening on :3000 <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

type myHandler string

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Gomicropro")
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "1234567890",
		MaxAge: 3600,
	})
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, string(mh))
	fmt.Fprintln(w, r.Header)

}
