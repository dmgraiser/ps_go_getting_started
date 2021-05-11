package main

import (
	"net/http"

	"github.com/dmgraiser/ps_go_getting_started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
