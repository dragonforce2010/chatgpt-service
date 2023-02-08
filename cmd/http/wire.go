//go:build wireinject
// +build wireinject

package main

import (
	"github.com/dragonforce2010/chatgpt-service/client"
	"github.com/dragonforce2010/chatgpt-service/domain/chat"
	"github.com/dragonforce2010/chatgpt-service/router"
	"github.com/google/wire"
)

func InitializeRouter() *router.Router {
	wire.Build(client.NewClient, router.NewRouter, chat.NewChatHandler, chat.NewChatRouter, chat.NewChatService)
	return &router.Router{}
}
