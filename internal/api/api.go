// Copyright (c) 2022 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

// Package api provides HTTP API of the Bridge.
//
// API endpoints:
//  * /focus, see focusHandler
package api

import (
	"fmt"
	"net/http"

	"github.com/ProtonMail/proton-bridge/internal/bridge"
	"github.com/ProtonMail/proton-bridge/internal/config/settings"
	"github.com/ProtonMail/proton-bridge/internal/events"
	"github.com/ProtonMail/proton-bridge/pkg/listener"
	"github.com/ProtonMail/proton-bridge/pkg/ports"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("pkg", "api") //nolint[gochecknoglobals]
)

type apiServer struct {
	host          string
	settings      *settings.Settings
	eventListener listener.Listener
}

// NewAPIServer returns prepared API server struct.
func NewAPIServer(settings *settings.Settings, eventListener listener.Listener) *apiServer { //nolint[golint]
	return &apiServer{
		host:          bridge.Host,
		settings:      settings,
		eventListener: eventListener,
	}
}

// Starts the server.
func (api *apiServer) ListenAndServe() {
	mux := http.NewServeMux()
	mux.HandleFunc("/focus", wrapper(api, focusHandler))

	addr := api.getAddress()
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Info("API listening at ", addr)
	if err := server.ListenAndServe(); err != nil {
		api.eventListener.Emit(events.ErrorEvent, "API failed: "+err.Error())
		log.Error("API failed: ", err)
	}
	defer server.Close() //nolint[errcheck]
}

func (api *apiServer) getAddress() string {
	port := api.settings.GetInt(settings.APIPortKey)
	newPort := ports.FindFreePortFrom(port)
	if newPort != port {
		api.settings.SetInt(settings.APIPortKey, newPort)
	}
	return getAPIAddress(api.host, newPort)
}

func getAPIAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
