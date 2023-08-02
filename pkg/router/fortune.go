package router

import (
	"fortune-back-go/pkg/controller"
	"net/http"
)

type FortuneRouter struct {
	FortuneController *controller.FortuneController
}

func NewFortuneRouter() *FortuneRouter {
	return &FortuneRouter{
		FortuneController: controller.NewFortuneController(),
	}
}

func (r *FortuneRouter) Initialize() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/get", r.FortuneController.GetFortune)

	return router
}
