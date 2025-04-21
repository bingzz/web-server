package routes

import (
	"fmt"
	"io"
	"net/http"
)

func InitializeAPI() {
	fmt.Println("Intializing Endpoints...")
	// Set endpoints

	http.HandleFunc("GET /hello", getHello)       // GET
	http.HandleFunc("POST /hello", writeHello)    // POST
	http.HandleFunc("PUT /hello", putHello)       // PUT
	http.HandleFunc("DELETE /hello", removeHello) // DELETE

}

func getHello(w http.ResponseWriter, r *http.Request) {

	// Send response back to client
	io.WriteString(w, "GET Hello")
}

func writeHello(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Failed to read request body", http.StatusBadRequest)
	// 	return
	// }
	// defer r.Body.Close()

	// var payload schema.RequestPayload
	// if err := json.Unmarshal(body, &payload); err != nil {
	// 	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Printf("Received payload: %+v\n", payload)

	// Send response back to client
	io.WriteString(w, "POST Hello")
}

func putHello(w http.ResponseWriter, r *http.Request) {
	// Send response back to client
	io.WriteString(w, "PUT Hello")
}

func removeHello(w http.ResponseWriter, r *http.Request) {
	// Send response back to client
	io.WriteString(w, "REMOVE Hello")
}
