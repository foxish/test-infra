/*
Copyright 2016 The Kubernetes Authors.

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

package cla

import (
	"fmt"

	"k8s.io/test-infra/prow/github"
	"k8s.io/test-infra/prow/plugins"
)

const pluginName = "cla"

func init() {
	plugins.RegisterStatusEventHandler(pluginName, handleStatusEvent)
}

type gitHubClient interface {
	CreateComment(owner, repo string, number int, comment string) error
	AddLabel(owner, repo string, number int, label string) error
	RemoveLabel(owner, repo string, number int, label string) error
}

func handleStatusEvent(pa *plugins.PluginAgent, se github.StatusEvent) error {
	return handle(pa.GitHubClient, se)
}

func handle(gc gitHubClient, se github.StatusEvent) error {
	fmt.Println("HANDLING STATUSEVENT %+v", se)
	return nil
}
