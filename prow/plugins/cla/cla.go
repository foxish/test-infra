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

const (
	pluginName     = "cla"
	claContextName = "cla/linuxfoundation"
	claYesLabel    = "cncf-cla: yes"
	claNoLabel     = "cncf-cla: no"
)

func init() {
	plugins.RegisterStatusEventHandler(pluginName, handleStatusEvent)
}

type gitHubClient interface {
	CreateComment(owner, repo string, number int, comment string) error
	AddLabel(owner, repo string, number int, label string) error
	RemoveLabel(owner, repo string, number int, label string) error
	GetPullRequest(owner, repo string, number int) (*github.PullRequest, error)
	FindIssues(query string) ([]github.Issue, error)
}

func handleStatusEvent(pa *plugins.PluginAgent, se github.StatusEvent) error {
	return handle(pa.GitHubClient, se)
}

func handle(gc gitHubClient, se github.StatusEvent) error {
	if se.State == nil || se.Context == nil {
		return fmt.Errorf("Invalid status event delivered with empty state/context")
	}

	if *se.Context != claContextName {
		return nil
	}

	if *se.State == github.StatusPending {
		// do nothing and wait for state to be updated.
		return
	}


	if status == github.StatusSuccess {
		if obj.HasLabel(claYesLabel) {
			// status is success and we've already applied 'cncf-cla: yes'.
			return
		}
		if obj.HasLabel(cncfClaNoLabel) {
			obj.RemoveLabel(cncfClaNoLabel)
		}
		obj.AddLabel(cncfClaYesLabel)
		return
	}


	issues, err := gc.FindIssues(fmt.Sprintf("%v&repo=test-foxish/test&type=pr", *se.SHA))
	if err != nil {
		return err
	}

	for _, issue := range issues {
		if checkLabel()




	}

	pr, err := gc.GetPullRequest("test-foxish", "test", issues[0].Number)
	fmt.Printf("### Result: %v: %+v\n", err, pr)

	if pr.Head.SHA == *se.SHA {
		fmt.Println("### FOUND HEAD!!!")
	} else {
		fmt.Println("### FOUND unHEAD!!!")
	}

	// Go through each returned PR if the label state is likely to change

	// Is pr.HEAD.SHA = returnedSHA? Then set label.

	// If it goes from yes -> no, post comment.
	// If it is no, post comment.
	// If it is yes, don't say anything, set label if needed.


	//if status == github.StatusSuccess {
	//	if obj.HasLabel(claYesLabel) {
	//		// status is success and we've already applied 'cncf-cla: yes'.
	//		return
	//	}
	//	if obj.HasLabel(cncfClaNoLabel) {
	//		obj.RemoveLabel(cncfClaNoLabel)
	//	}
	//	obj.AddLabel(cncfClaYesLabel)
	//	return
	//}
	//
	//// If we are here, that means that the context is failure/error.
	//comments, err := obj.ListComments()
	//if err != nil {
	//	glog.Error(err)
	//	return
	//}
	//who := mungerutil.GetIssueUsers(obj.Issue).Author.Mention().Join()
	//
	//// Get a notification if it's time to ping.
	//notif := cla.pinger.PingNotification(
	//	comments,
	//	who,
	//	nil,
	//)
	//if notif != nil {
	//	obj.WriteComment(notif.String())
	//}
	//
	//if obj.HasLabel(cncfClaNoLabel) {
	//	// status reported error/failure and we've already applied 'cncf-cla: no' label.
	//	return
	//}
	//
	//if obj.HasLabel(cncfClaYesLabel) {
	//	obj.RemoveLabel(cncfClaYesLabel)
	//}
	//obj.AddLabel(cncfClaNoLabel)

	return nil
}

func checkLabel(toFind string, labels []github.Label) bool {
	for _, label := range labels {
		if toFind == label.Name {
			return true
		}
	}
	return false
}
