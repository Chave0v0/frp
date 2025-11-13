package frpc

import (
	"embed"

	"github.com/Chave0v0/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
