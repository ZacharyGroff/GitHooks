package router

import (
	"log"
	"net/http"
	"io/ioutil"
	"github.com/ZacharyGroff/GitDeployer/processors"
	"github.com/ZacharyGroff/GitDeployer/models"
	"github.com/ZacharyGroff/GitDeployer/validation"
)

type Router struct {
	pushProcessor *processors.PushProcessor
	validator *validation.Validator
}

func (router Router) Route(request *http.Request) {
	message, err := models.NewMessage(request)
	if err != nil {
		log.Fatalf("Failed to create message with error: %s\n", err)
	}

	err = router.validate(message, request)
	if err != nil {
		log.Fatalf("Message validation failed with error: %s\n", err)
	}

	err = router.routeEvent(message)
	if err != nil {
		log.Fatalf("Unable to route event with error %s\n", err)
	}
}

func (router Router) validate(message *models.Message, request *http.Request) error {
	hmac, err := message.GetHeaderField("X-Hub-Signature")
	trimmedHmac := []byte(hmac)[5:]
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return err
	}

	return router.validator.ValidateHmac(trimmedHmac, body)
}

func (router Router) routeEvent(message *models.Message) error {
	event, err := message.GetHeaderField("X-Github-Event")

	if err != nil {
		return err
	}

	switch event {
	case "push":
		router.pushProcessor.HandleMessage(message)
	default:
		log.Printf("Unhandled event %s detected.\n", event)
	}

	return nil
}

func NewRouter(p *processors.PushProcessor, v *validation.Validator) *Router {
	return &Router{p, v}
}
