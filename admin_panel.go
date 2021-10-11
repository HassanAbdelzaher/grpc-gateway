package admin_panel

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed admin_panel/build/*
var EmbdedContent embed.FS

type AdminPanelHandler struct {
	handler http.Handler
}

func NewAdminPanelHandler() *AdminPanelHandler {
	fsys := fs.FS(EmbdedContent)
	contentStatic, _ := fs.Sub(fsys, "admin_panel/build")
	staticFileServer := http.FileServer(http.FS(contentStatic))
	return &AdminPanelHandler{handler: staticFileServer}
}

func (p *AdminPanelHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			wr.WriteHeader(500)
			fmt.Fprint(wr, r)
		}
	}()
	if p.handler == nil {
		wr.WriteHeader(500)
		fmt.Fprint(wr, "invalied static http handler")
		return
	}
	p.handler.ServeHTTP(wr, r)
}
