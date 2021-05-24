package events

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CDEventType string

const (

	// PipelineRun events
	PipelineRunStartedEventV1 CDEventType = "cd.pipelinerun.started.v1"
	PipelineRunFinishedEventV1 CDEventType = "cd.pipelinerun.finished.v1"
	PipelineRunQueuedEventV1 CDEventType = "cd.pipelinerun.queued.v1"

	// TaskRun events
	TaskRunStartedEventV1 CDEventType = "cd.taskrun.started.v1"
	TaskRunFinishedEventV1 CDEventType = "cd.taskrun.finished.v1"

	// Repository events
	RepositoryCreatedEventV1 CDEventType = "cd.repository.created.v1"
	RepositoryModifiedEventV1 CDEventType = "cd.repository.modified.v1"
	RepositoryDeletedEventV1 CDEventType = "cd.repository.deleted.v1"

	BranchCreatedEventV1 CDEventType = "cd.repository.branch.created.v1"
	BranchDeletedEventV1 CDEventType = "cd.repository.branch.deleted.v1"

	// Change Events
	ChangeCreatedEventV1 CDEventType = "cd.repository.change.created.v1"
	ChangeUpdatedEventV1 CDEventType = "cd.repository.change.updated.v1"
	ChangeReviewedEventV1 CDEventType = "cd.repository.change.reviewed.v1"
	ChangeMergedEventV1 CDEventType = "cd.repository.change.merged.v1"
	ChangeAbandonedEventV1 CDEventType = "cd.repository.change.abandoned.v1"

	// Build Events
	BuildStartedEventV1 CDEventType = "cd.build.started.v1"
	BuildQueuedEventV1 CDEventType = "cd.build.queued.v1"
	BuildFinishedEventV1 CDEventType = "cd.build.finished.v1"

	// Test Events
	TestCaseStartedEventV1 CDEventType = "cd.test.case.started.v1"
	TestCaseQueuedEventV1 CDEventType = "cd.test.case.queued.v1"
	TestCaseFinishedEventV1 CDEventType = "cd.test.case.finished.v1"

	TestSuiteStartedEventV1 CDEventType = "cd.test.suite.started.v1"
	TestSuiteQueuedEventV1 CDEventType = "cd.test.suite.queued.v1"
	TestSuiteFinishedEventV1 CDEventType = "cd.test.suite.finished.v1"

	// Artifact Events
	ArtifactPackagedEventV1 CDEventType = "cd.artifact.packaged.v1"
	ArtifactPublishedEventV1 CDEventType = "cd.artifact.published.v1"

	// Environment Events
	EnvironmentCreatedEventV1 CDEventType = "cd.environment.created.v1"
	EnvironmentModifiedEventV1 CDEventType = "cd.environment.modified.v1"
	EnvironmentDeletedEventV1 CDEventType = "cd.environment.deleted.v1"

	// Service Events
	ServiceDeployedEventV1 CDEventType = "cd.service.deployed.v1"
	ServiceUpgradedEventV1 CDEventType = "cd.service.upgraded.v1"
	ServiceRolledbackEventV1 CDEventType = "cd.service.rolledback.v1"
	ServiceRemovedEventV1 CDEventType = "cd.service.removed.v1"

)


func (t CDEventType) String() string {
	return string(t)
}


func CreateRepositoryEvent(eventType CDEventType,
	repositoryId string,
	repositoryName string,
	repositoryURL string,
	repositoryData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(string(eventType))
	event.SetTime(time.Now())

	setExtensionForRepositoryEvents(event, repositoryId, repositoryName, repositoryURL )

	event.SetData(cloudevents.ApplicationJSON, repositoryData)

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
	envData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(string(eventType))
	event.SetTime(time.Now())

	setExtensionForTaskRunEvents(event, taskRunId, taskRunName, taskRunPipelineId )

	event.SetData(cloudevents.ApplicationJSON, envData)

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
	envData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(string(eventType))
	event.SetTime(time.Now())

	setExtensionForPipelineRunEvents(event, pipelineRunId, pipelineRunName, pipelineRunStatus,pipelineRunURL, pipelineRunErrors )

	event.SetData(cloudevents.ApplicationJSON, envData)

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
	event.SetType(string(eventType))
	event.SetTime(time.Now())

	setExtensionForEnvEvents(event, envId, envName, envRepoUrl)

	event.SetData(cloudevents.ApplicationJSON, envData)

	return event, nil
}

func setExtensionForEnvEvents(event cloudevents.Event, envId string, envName string, envRepoUrl string) {
	event.SetExtension("envId", envId)
	event.SetExtension("envname", envName)
	event.SetExtension("envrepourl", envRepoUrl)

}


func CreateServiceEvent(eventType CDEventType,
	serviceEnvId string,
	serviceName string,
	serviceVersion string,
	serviceData map[string]string) (cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(uuid.NewV4().String())
	event.SetType(string(eventType))
	event.SetTime(time.Now())

	setExtensionForServiceEvents(event, serviceEnvId, serviceName, serviceVersion)

	event.SetData(cloudevents.ApplicationJSON, serviceData)

	return event, nil
}

func setExtensionForServiceEvents(event cloudevents.Event, serviceEnvId string, serviceName string, serviceVersion string) {
	event.SetExtension("serviceenvid", serviceEnvId)
	event.SetExtension("servicename", serviceName)
	event.SetExtension("serviceversion", serviceVersion)
}