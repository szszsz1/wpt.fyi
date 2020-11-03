// Copyright 2018 The WPT Dashboard Project. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package webapp

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/web-platform-tests/wpt.fyi/shared"
	"google.golang.org/appengine"
)

// NOTE(lukebjerring): If tweaking service worker locally, change to
// sevenCharSHA, _ = regexp.Compile("^[0-9a-f]{7}|None$")
var sevenCharSHA, _ = regexp.Compile("^[0-9a-f]{7}$")

func serviceWorkerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	aeAPI := shared.NewAppEngineAPI(ctx)
	if !aeAPI.IsFeatureEnabled("serviceWorker") {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-Type", "application/javascript")
	version := strings.Split(appengine.VersionID(ctx), ".")[0]
	if !sevenCharSHA.MatchString(version) {
		http.Error(w, fmt.Sprintf("Service worker not implemented for version '%s'", version), http.StatusNotImplemented)
		return
	}

	data := struct {
		Version string
	}{
		Version: version,
	}
	// We don't need user info in this template.
	RenderTemplate(w, nil, "service-worker.js", data)
}
