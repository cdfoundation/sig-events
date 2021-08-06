package events

import (
	"encoding/json"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
	cetest "github.com/cloudevents/sdk-go/v2/test"
	"github.com/stretchr/testify/assert"
)

type expectedEventData struct {
	Type       string
	Data       map[string]string
	Extensions map[string]interface{}
}

func TestArtifactPackagedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.artifact.packaged.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"artifactid":      "",
			"artifactname":    "",
			"artifactversion": "",
		},
	}

	eventParams := ArtifactEventParams{ArtifactData: data}

	actual, _ := CreateArtifactEvent(ArtifactPackagedEventV1, eventParams)

	assertEvent(t, actual, expectedEvent)
}

func TestArtifactPublishedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.artifact.published.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"artifactid":      "",
			"artifactname":    "",
			"artifactversion": "",
		},
	}

	eventParams := ArtifactEventParams{ArtifactData: data}

	actual, _ := CreateArtifactEvent(ArtifactPublishedEventV1, eventParams)

	assertEvent(t, actual, expectedEvent)
}

func TestBuildStartedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.build.started.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"buildid":         "buildId",
			"buildname":       "buildName",
			"buildartifactid": "artifact",
		},
	}

	actual, _ := CreateBuildEvent(BuildStartedEventV1, "buildId", "buildName", "artifact", data)
	assertEvent(t, actual, expectedEvent)
}

func TestBuildQueuedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.build.queued.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"buildid":         "buildId",
			"buildname":       "buildName",
			"buildartifactid": "artifact",
		},
	}

	actual, _ := CreateBuildEvent(BuildQueuedEventV1, "buildId", "buildName", "artifact", data)
	assertEvent(t, actual, expectedEvent)
}

func TestBuildFinishedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.build.finished.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"buildid":         "buildId",
			"buildname":       "buildName",
			"buildartifactid": "artifact",
		},
	}

	actual, _ := CreateBuildEvent(BuildFinishedEventV1, "buildId", "buildName", "artifact", data)
	assertEvent(t, actual, expectedEvent)
}

func TestBranchCreatedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.repository.branch.created.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"branchid":           "branch id",
			"branchname":         "branch name",
			"branchrepositoryid": "repo id",
		},
	}

	actual, _ := CreateBranchEvent(BranchCreatedEventV1, "branch id", "branch name", "repo id", data)
	assertEvent(t, actual, expectedEvent)

}

func TestBranchDeletedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.repository.branch.deleted.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"branchid":           "branch id",
			"branchname":         "branch name",
			"branchrepositoryid": "repo id",
		},
	}

	actual, _ := CreateBranchEvent(BranchDeletedEventV1, "branch id", "branch name", "repo id", data)
	assertEvent(t, actual, expectedEvent)
}

func TestRepositoryCreatedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.repository.created.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"repositoryid":   "repository id",
			"repositoryname": "repository name",
			"repositoryurl":  "repository url",
		},
	}

	actual, _ := CreateRepositoryEvent(RepositoryCreatedEventV1, "repository id", "repository name", "repository url", data)
	assertEvent(t, actual, expectedEvent)
}

func TestRepositoryModifiedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.repository.modified.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"repositoryid":   "repository id",
			"repositoryname": "repository name",
			"repositoryurl":  "repository url",
		},
	}

	actual, _ := CreateRepositoryEvent(RepositoryModifiedEventV1, "repository id", "repository name", "repository url", data)
	assertEvent(t, actual, expectedEvent)
}

func TestRepositoryDeletedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.repository.deleted.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"repositoryid":   "repository id",
			"repositoryname": "repository name",
			"repositoryurl":  "repository url",
		},
	}

	actual, _ := CreateRepositoryEvent(RepositoryDeletedEventV1, "repository id", "repository name", "repository url", data)
	assertEvent(t, actual, expectedEvent)
}

func TestTaskRunStartedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.taskrun.started.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"taskrunid":         "taskrun id",
			"taskrunname":       "taskrun name",
			"taskrunpipelineid": "taskrunpipeline id",
		},
	}

	actual, _ := CreateTaskRunEvent(TaskRunStartedEventV1, "taskrun id", "taskrun name", "taskrunpipeline id", data)
	assertEvent(t, actual, expectedEvent)
}

func TestTaskRunFinishedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.taskrun.finished.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"taskrunid":         "taskrun id",
			"taskrunname":       "taskrun name",
			"taskrunpipelineid": "taskrunpipeline id",
		},
	}

	actual, _ := CreateTaskRunEvent(TaskRunFinishedEventV1, "taskrun id", "taskrun name", "taskrunpipeline id", data)
	assertEvent(t, actual, expectedEvent)
}

func TestPipelineRunStartedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.pipelinerun.started.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"pipelinerunid":     "pipeline run id",
			"pipelinerunname":   "pipeline run name",
			"pipelinerunstatus": "pipeline run status",
			"pipelinerunurl":    "pipeline run url",
			"pipelinerunerrors": "pipeline run errors",
		},
	}

	actual, _ := CreatePipelineRunEvent(
		PipelineRunStartedEventV1,
		"pipeline run id",
		"pipeline run name",
		"pipeline run status",
		"pipeline run url",
		"pipeline run errors",
		data)
	assertEvent(t, actual, expectedEvent)
}

func TestPipelineRunQueuedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.pipelinerun.queued.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"pipelinerunid":     "pipeline run id",
			"pipelinerunname":   "pipeline run name",
			"pipelinerunstatus": "pipeline run status",
			"pipelinerunurl":    "pipeline run url",
			"pipelinerunerrors": "pipeline run errors",
		},
	}

	actual, _ := CreatePipelineRunEvent(
		PipelineRunQueuedEventV1,
		"pipeline run id",
		"pipeline run name",
		"pipeline run status",
		"pipeline run url",
		"pipeline run errors",
		data)
	assertEvent(t, actual, expectedEvent)
}

func TestPipelineRunFinishedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.pipelinerun.finished.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"pipelinerunid":     "pipeline run id",
			"pipelinerunname":   "pipeline run name",
			"pipelinerunstatus": "pipeline run status",
			"pipelinerunurl":    "pipeline run url",
			"pipelinerunerrors": "pipeline run errors",
		},
	}

	actual, _ := CreatePipelineRunEvent(
		PipelineRunFinishedEventV1,
		"pipeline run id",
		"pipeline run name",
		"pipeline run status",
		"pipeline run url",
		"pipeline run errors",
		data)
	assertEvent(t, actual, expectedEvent)
}

func TestEnvironmentCreatedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.environment.created.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"envid":      "env Id",
			"envname":    "env name",
			"envrepourl": "env repo url",
		},
	}

	actual, _ := CreateEnvironmentEvent(EnvironmentCreatedEventV1, "env Id", "env name", "env repo url", data)
	assertEvent(t, actual, expectedEvent)
}

func TestEnvironmentModifiedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.environment.modified.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"envid":      "env Id",
			"envname":    "env name",
			"envrepourl": "env repo url",
		},
	}

	actual, _ := CreateEnvironmentEvent(EnvironmentModifiedEventV1, "env Id", "env name", "env repo url", data)
	assertEvent(t, actual, expectedEvent)
}

func TestEnvironmentDeletedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.environment.deleted.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"envid":      "env Id",
			"envname":    "env name",
			"envrepourl": "env repo url",
		},
	}

	actual, _ := CreateEnvironmentEvent(EnvironmentDeletedEventV1, "env Id", "env name", "env repo url", data)
	assertEvent(t, actual, expectedEvent)
}

func TestServiceDeployedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.service.deployed.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"serviceenvid":   "service env id",
			"servicename":    "service name",
			"serviceversion": "service version",
		},
	}

	actual, _ := CreateServiceEvent(ServiceDeployedEventV1, "service env id",	"service name",	"service version", data)
	assertEvent(t, actual, expectedEvent)
}

func TestServiceUpgradedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.service.upgraded.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"serviceenvid":   "service env id",
			"servicename":    "service name",
			"serviceversion": "service version",
		},
	}

	actual, _ := CreateServiceEvent(ServiceUpgradedEventV1, "service env id",	"service name",	"service version", data)
	assertEvent(t, actual, expectedEvent)
}

func TestServiceRemovedEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.service.removed.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"serviceenvid":   "service env id",
			"servicename":    "service name",
			"serviceversion": "service version",
		},
	}

	actual, _ := CreateServiceEvent(ServiceRemovedEventV1, "service env id",	"service name",	"service version", data)
	assertEvent(t, actual, expectedEvent)
}

func TestServiceRolledbackEvent(t *testing.T) {

	data := map[string]string{"test": "data"}

	expectedEvent := expectedEventData{
		Type: "cd.service.rolledback.v1",
		Data: data,
		Extensions: map[string]interface{}{
			"serviceenvid":   "service env id",
			"servicename":    "service name",
			"serviceversion": "service version",
		},
	}

	actual, _ := CreateServiceEvent(ServiceRolledbackEventV1, "service env id",	"service name",	"service version", data)
	assertEvent(t, actual, expectedEvent)
}

func assertEvent(t *testing.T, event event.Event, expected expectedEventData) {

	extensionNames := make([]string, 0, len(expected.Extensions))
	for key := range expected.Extensions {
		extensionNames = append(extensionNames, key)
	}

	assert.NotEmpty(t, event.ID())
	assert.NotEmpty(t, event.Time())
	assert.NoError(t, cetest.HasExtensions(expected.Extensions)(event))
	containsExtensions := cetest.ContainsExactlyExtensions(extensionNames...)
	assert.NoError(t, containsExtensions(event))
	assert.NoError(t, cetest.HasType(expected.Type)(event))
	bytes, _ := json.Marshal(expected.Data)
	assert.NoError(t, cetest.HasData(bytes)(event))
}
