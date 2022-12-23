package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vansh0100/E-Library/controller"
	"github.com/Vansh0100/E-Library/router"
)

//
const mongoUrl="mongodb+srv://vanshj:vansh238@cluster0.n789mnp.mongodb.net/?retryWrites=true&w=majority"

func main() {
	controller.Connection(mongoUrl)

	r:=router.Router()

	fmt.Println("Server is up and running!")
	log.Fatal(http.ListenAndServe(":4000",r))
}