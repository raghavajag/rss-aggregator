package main

import "net/http"

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong!")
}
