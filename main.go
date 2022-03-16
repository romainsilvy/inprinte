package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"

	utils "inprinte/backend/utils"
	CRUD "inprinte/backend/CRUD"
	// databaseTools "inprinte/backend/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}











func manageRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/users/", CRUD.GetUsers).Methods("GET")

	// router.HandleFunc("/user/{userid}", CreateUser).Methods("POST")
	// router.HandleFunc("/user/{userid}", CRUD.GetUsers).Methods("GET")
	// router.HandleFunc("/user/{userid}", UpdateUser).Methods("UPDATE")
	// router.HandleFunc("/user/{userid}", DeleteUser).Methods("DELETE")


	fmt.Println("Server at 8080")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ZEUB")
}

func main() {
	utils.InprinteAscii()

	r := mux.NewRouter()

    r.HandleFunc("/", CRUD.GetUsers)

    http.ListenAndServe(":8080", r)

}
