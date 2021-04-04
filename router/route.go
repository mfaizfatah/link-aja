package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/link-aja/controller"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

// route struct with value Controllers Interface
type route struct {
	ctrl controller.Controllers
}

// Router represent the Router contract
type Router interface {
	Router(port string)
}

/*NewRouter will create an object that represent the Router interface (Router)
 * @parameter
 * c - controllers Interface
 *
 * @represent
 * interface Router
 *
 * @return
 * struct route with value Controllers Interface
 */
func NewRouter(c controller.Controllers) Router {
	return &route{ctrl: c}
}

func (c *route) Router(port string) {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Get("/account/{accNumber}", c.ctrl.CheckSaldo)
		r.Post("/account/{accNumber}/transfer", c.ctrl.Transfer)
	})

	logrus.Infof("Server running on port : %s", port)
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), cors.AllowAll().Handler(router)))
}
