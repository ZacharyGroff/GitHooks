// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"config"
	"processors"
)

// Injectors from wire.go:

func InitializeEndpoint() Endpoint {
	configConfig := config.NewConfig()
	deployer := processors.NewDeployer(configConfig)
	pushProcessor := processors.NewPushProcessor(deployer, configConfig)
	endpoint := NewEndpoint(configConfig, pushProcessor)
	return endpoint
}
