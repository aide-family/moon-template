// Package api is the api package for the sovereign service.
package api

import (
	nethttp "net/http"
	"strings"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/aide-family/sovereign/pkg/middler"
)

type BasicAuthConfig interface {
	GetEnabled() string
	GetUsername() string
	GetPassword() string
}

type HandlerBinding struct {
	Name      string
	Enabled   bool
	BasicAuth BasicAuthConfig
	Handler   nethttp.Handler
	Path      string
	UsePrefix bool
	FullPath  string
}

func BindHandlerWithAuth(httpSrv *http.Server, binding HandlerBinding) {
	if !binding.Enabled {
		klog.Debugf("%s is not enabled", binding.Name)
		return
	}

	endpoint, err := httpSrv.Endpoint()
	if err != nil {
		klog.Errorw("msg", "get http server endpoint failed", "error", err)
		return
	}

	handler := binding.Handler
	basicAuth := binding.BasicAuth
	if strings.EqualFold(basicAuth.GetEnabled(), "true") {
		handler = middler.BasicAuthMiddleware(basicAuth.GetUsername(), basicAuth.GetPassword())(handler)
		klog.Debugf("[%s] endpoint: %s%s (Basic Auth: %s:%s)", binding.Name, endpoint, binding.FullPath, basicAuth.GetUsername(), basicAuth.GetPassword())
	} else {
		klog.Debugf("[%s] endpoint: %s%s (No Basic Auth)", binding.Name, endpoint, binding.FullPath)
	}

	if binding.UsePrefix {
		httpSrv.HandlePrefix(binding.Path, handler)
	} else {
		httpSrv.Handle(binding.Path, handler)
	}
}
