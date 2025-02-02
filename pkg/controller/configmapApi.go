/*
Copyright 2016 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/fission/fission/pkg/utils"
)

func (a *API) ConfigMapExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["configmap"]
	ns := a.extractQueryParamFromRequest(r, "namespace")
	if len(ns) == 0 {
		ns = metav1.NamespaceDefault
	}

	_, err := a.kubernetesClient.CoreV1().ConfigMaps(ns).Get(r.Context(), name, metav1.GetOptions{})
	if err != nil {
		name = utils.EscapeQuotes(name)
		ns = utils.EscapeQuotes(ns)
		a.logger.Error("error getting config map", zap.Error(err), zap.String("config_map_name", name), zap.String("namespace", ns))
		a.respondWithError(w, err)
		return
	}
	a.respondWithSuccess(w, nil)
}
