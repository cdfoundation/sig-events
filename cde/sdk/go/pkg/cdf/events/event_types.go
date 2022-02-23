/*
Copyright 2021 The CD Events SDK Authors

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
package events

import (
	"time"

	"github.com/cdfoundation/sig-events/cde/sdk/go/pkg/cdf/events/extensions"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	uuid "github.com/satori/go.uuid"
)

type CDEventType string

const (

	// PipelineRun events
	PipelineRunStartedEventV1  CDEventType = "dev.cdevents.pipelinerun.started.v1"
	PipelineRunFinishedEventV1 CDEventType = "dev.cdevents.pipelinerun.finished.v1"
	PipelineRunQueuedEventV1   CDEventType = "dev.cdevents.pipelinerun.queued.v1"

	// TaskRun events
	TaskRunStartedEventV1  CDEventType = "dev.cdevents.taskrun.started.v1"
	TaskRunFinishedEventV1 CDEventType = "dev.cdevents.taskrun.finished.v1"

	// Repository events
	RepositoryCreatedEventV1  CDEventType = "dev.cdevents.repository.created.v1"
	RepositoryModifiedEventV1 CDEventType = "dev.cdevents.repository.modified.v1"
	RepositoryDeletedEventV1  CDEventType = "dev.cdevents.repository.deleted.v1"

	BranchCreatedEventV1 CDEventType = "dev.cdevents.repository.branch.created.v1"
	BranchDeletedEventV1 CDEventType = "dev.cdevents.repository.branch.deleted.v1"

	// Change Events
	ChangeCreatedEventV1   CDEventType = "dev.cdevents.repository.change.created.v1"
	ChangeUpdatedEventV1   CDEventType = "dev.cdevents.repository.change.updated.v1"
	ChangeReviewedEventV1  CDEventType = "dev.cdevents.repository.change.reviewed.v1"
	ChangeMergedEventV1    CDEventType = "dev.cdevents.repository.change.merged.v1"
	ChangeAbandonedEventV1 CDEventType = "dev.cdevents.repository.change.abandoned.v1"

	// Build Events
	BuildStartedEventV1  CDEventType = "dev.cdevents.build.started.v1"
	BuildQueuedEventV1   CDEventType = "dev.cdevents.build.queued.v1"
	BuildFinishedEventV1 CDEventType = "dev.cdevents.build.finished.v1"

	// Test Events
	TestCaseStartedEventV1  CDEventType = "dev.cdevents.test.case.started.v1"
	TestCaseQueuedEventV1   CDEventType = "dev.cdevents.test.case.queued.v1"
	TestCaseFinishedEventV1 CDEventType = "dev.cdevents.test.case.finished.v1"

	TestSuiteStartedEventV1  CDEventType = "dev.cdevents.test.suite.started.v1"
	TestSuiteQueuedEventV1   CDEventType = "dev.cdevents.test.suite.queued.v1"
	TestSuiteFinishedEventV1 CDEventType = "dev.cdevents.test.suite.finished.v1"

	// Artifact Events
	ArtifactPackagedEventV1  CDEventType = "dev.cdevents.artifact.packaged.v1"
	ArtifactPublishedEventV1 CDEventType = "dev.cdevents.artifact.published.v1"

	// Environment Events
	EnvironmentCreatedEventV1  CDEventType = "dev.cdevents.environment.created.v1"
	EnvironmentModifiedEventV1 CDEventType = "dev.cdevents.environment.modified.v1"
	EnvironmentDeletedEventV1  CDEventType = "dev.cdevents.environment.deleted.v1"

	// Service Events
	ServiceDeployedEventV1   CDEventType = "dev.cdevents.service.deployed.v1"
	ServiceUpgradedEventV1   CDEventType = "dev.cdevents.service.upgraded.v1"
	ServiceRolledbackEventV1 CDEventType = "dev.cdevents.service.rolledback.v1"
	ServiceRemovedEventV1    CDEventType = "dev.cdevents.service.removed.v1"
	ServicePublishedV1       CDEventType = "dev.cdevents.service.published,v1"
)

func (t CDEventType) String() string {
	return string(t)
}

type ServiceEventParams struct {
	ServiceEnvID              string
	ServiceVersion            string
	ServiceName               string
	ServiceNamespace          string
	ServiceData               map[string]string
	ServiceActiveRevisionName string
}

type ArtifactEventParams struct {
	ArtifactId      string
	ArtifactName    string
	ArtifactVersion string
	ArtifactData    map[string]string
}

func CreateArtifactEvent(eventType CDEventType, params ArtifactEventParams) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForArtifactEvents(event, params.ArtifactId, params.ArtifactName, params.ArtifactVersion)

	err := event.SetData(cloudevents.ApplicationJSON, params.ArtifactData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForArtifactEvents(event cloudevents.Event, artifactId string, artifactName string, artifactVersion string) {
	event.SetExtension(extensions.ArtifactIdExtension, artifactId)
	event.SetExtension(extensions.ArtifactNameExtension, artifactName)
	event.SetExtension(extensions.ArtifactVersionExtension, artifactVersion)

}

func CreateBuildEvent(eventType CDEventType,
	buildId string,
	buildName string,
	buildArtifactId string,
	buildData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForBuildEvents(event, buildId, buildName, buildArtifactId)

	err := event.SetData(cloudevents.ApplicationJSON, buildData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForBuildEvents(event cloudevents.Event, buildId string, buildName string, buildArtifactId string) {
	event.SetExtension("buildid", buildId)
	event.SetExtension("buildname", buildName)
	event.SetExtension("buildartifactid", buildArtifactId)

}

func CreateBranchEvent(eventType CDEventType,
	branchId string,
	branchName string,
	branchRepositoryId string,
	branchData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForBranchEvents(event, branchId, branchName, branchRepositoryId)

	err := event.SetData(cloudevents.ApplicationJSON, branchData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForBranchEvents(event cloudevents.Event, branchId string, branchName string, branchRepositoryId string) {
	event.SetExtension("branchid", branchId)
	event.SetExtension("branchname", branchName)
	event.SetExtension("branchrepositoryid", branchRepositoryId)

}

func CreateRepositoryEvent(eventType CDEventType,
	repositoryId string,
	repositoryName string,
	repositoryURL string,
	repositoryData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForRepositoryEvents(event, repositoryId, repositoryName, repositoryURL)

	err := event.SetData(cloudevents.ApplicationJSON, repositoryData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForRepositoryEvents(event cloudevents.Event, repositoryId string, repositoryName string, repositoryURL string) {
	event.SetExtension("repositoryid", repositoryId)
	event.SetExtension("repositoryname", repositoryName)
	event.SetExtension("repositoryurl", repositoryURL)

}

func CreateTaskRunEvent(eventType CDEventType,
	taskRunId string,
	taskRunName string,
	taskRunPipelineId string,
	taskRunData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForTaskRunEvents(event, taskRunId, taskRunName, taskRunPipelineId)

	err := event.SetData(cloudevents.ApplicationJSON, taskRunData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForTaskRunEvents(event cloudevents.Event, taskRunId string, taskRunName string, taskRunPipelineId string) {
	event.SetExtension("taskrunid", taskRunId)
	event.SetExtension("taskrunname", taskRunName)
	event.SetExtension("taskrunpipelineid", taskRunPipelineId)

}

func CreatePipelineRunEvent(eventType CDEventType,
	pipelineRunId string,
	pipelineRunName string,
	pipelineRunStatus string,
	pipelineRunURL string,
	pipelineRunErrors string,
	pipelineRunData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForPipelineRunEvents(event, pipelineRunId, pipelineRunName, pipelineRunStatus, pipelineRunURL, pipelineRunErrors)

	err := event.SetData(cloudevents.ApplicationJSON, pipelineRunData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForPipelineRunEvents(event cloudevents.Event, pipelineRunId string, pipelineRunName string,
	pipelineRunStatus string, pipelineRunURL string, pipelineRunErrors string) {
	event.SetExtension("pipelinerunid", pipelineRunId)
	event.SetExtension("pipelinerunname", pipelineRunName)
	event.SetExtension("pipelinerunstatus", pipelineRunStatus)
	event.SetExtension("pipelinerunurl", pipelineRunURL)
	event.SetExtension("pipelinerunerrors", pipelineRunErrors)
}

func CreateEnvironmentEvent(eventType CDEventType,
	envId string,
	envName string,
	envRepoUrl string,
	envData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForEnvEvents(event, envId, envName, envRepoUrl)

	err := event.SetData(cloudevents.ApplicationJSON, envData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForEnvEvents(event cloudevents.Event, envId string, envName string, envRepoUrl string) {
	event.SetExtension("envId", envId)
	event.SetExtension("envname", envName)
	event.SetExtension("envrepourl", envRepoUrl)

}

func CreateServiceEvent(eventType CDEventType,
	params ServiceEventParams) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(eventType.String())
	event.SetTime(time.Now())

	setExtensionForServiceEvents(event, params)

	err := event.SetData(cloudevents.ApplicationJSON, params.ServiceData)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return event, nil
}

func setExtensionForServiceEvents(event cloudevents.Event, params ServiceEventParams) {
	emptyOmitter := NewEmptyOmittedEvent(&event)
	emptyOmitter.Set(extensions.ServiceEnvIdExtension, params.ServiceEnvID).
		Set(extensions.ServiceNameExtension, params.ServiceName).
		Set(extensions.ServiceVersionExtension, params.ServiceVersion).
		Set(extensions.ServiceNamespaceExtension, params.ServiceNamespace).
		Set(extensions.ServiceActiveRevisionNameExtension, params.ServiceActiveRevisionName)
}

type emptyOmitted struct {
	event *cloudevents.Event
}

func NewEmptyOmittedEvent(event *cloudevents.Event) *emptyOmitted {
	return &emptyOmitted{
		event: event,
	}
}

// Set checks to see if the value being set is not an empty string
// if so, it will skip the field in the cloud event. This behavior
// is very similar to the omitempty struct tag of the json package
func (e *emptyOmitted) Set(ext, val string) *emptyOmitted {
	if val != "" {
		e.event.SetExtension(ext, val)
	}

	return e
}

func (e *emptyOmitted) GetEvent() *cloudevents.Event {
	return e.event
}
