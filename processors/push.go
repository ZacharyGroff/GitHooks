package processors

import (
	"log"
	"net/http"
	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/endpoint"
)

type PushProcessor struct {
	Deployer *Deployer
	Config *config.Config
}

func NewPushProcessor(deployer *Deployer, config *config.Config) *PushProcessor {
	return &PushProcessor{deployer, config}
}

func (pushProcessor PushProcessor) HandleRequest(message *main.Message) {
	output := pushProcessor.Deployer.Deploy()
	log.Printf("Successfully deployed script with output:\n%s\n", output)
}
