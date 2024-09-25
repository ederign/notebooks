/*
Copyright 2024.

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

package api

import (
	"github.com/kubeflow/notebooks/workspaces/backend/internal/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *App) GetWorkspacesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	namespace := ps.ByName(NamespacePathParam)

	var workspaces []models.WorkspaceModel
	var err error
	if namespace == "" {
		workspaces, err = a.repositories.Workspace.GetAllWorkspaces(r.Context())
	} else {
		workspaces, err = a.repositories.Workspace.GetWorkspaces(r.Context(), namespace)
	}

	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	modelRegistryRes := Envelope{
		"workspaces": workspaces,
	}

	err = a.WriteJSON(w, http.StatusOK, modelRegistryRes, nil)

	if err != nil {
		a.serverErrorResponse(w, r, err)
	}

}
