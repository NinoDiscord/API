package controllers

type Controller struct {
	Automod AutomodController
}

func NewController() *Controller {
	return &Controller{
		AutomodController{},
	}
}
