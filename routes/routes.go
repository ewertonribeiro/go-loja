package routes

import (
	"net/http"

	"github.com/ewertonribeiro/go-loja/controllers"
)

func Router() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
