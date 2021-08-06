package cmd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/event"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
	cetest "github.com/cloudevents/sdk-go/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestArtifactEvent(t *testing.T) {
	executeAndVerifyEvent(t, []string{"artifact", "packaged"}, "cd.artifact.packaged.v1")
	executeAndVerifyEvent(t, []string{"artifact", "finished"}, "cd.artifact.published.v1")
}

func TestBranchEvent(t *testing.T) {
	executeAndVerifyEvent(t, []string{"branch", "created"}, "cd.repository.branch.created.v1")
	executeAndVerifyEvent(t, []string{"branch", "deleted"}, "cd.repository.branch.deleted.v1")
}

func TestBuildEvent(t *testing.T) {
	executeAndVerifyEvent(t, []string{"build", "started"}, "cd.build.started.v1")
	executeAndVerifyEvent(t, []string{"build", "finished"}, "cd.build.finished.v1")
	executeAndVerifyEvent(t, []string{"build", "queued"}, "cd.build.queued.v1")
}

func TestEnvironmentEvent(t *testing.T) {
		executeAndVerifyEvent(t, []string{"env", "created"}, "cd.environment.created.v1")
		executeAndVerifyEvent(t, []string{"env", "deleted"}, "cd.environment.deleted.v1")
		executeAndVerifyEvent(t, []string{"env", "modified"}, "cd.environment.modified.v1")
}

func TestPipelineEvent(t *testing.T) {
	executeAndVerifyEvent(t, []string{"pipelinerun", "started"}, "cd.pipelinerun.started.v1")
	executeAndVerifyEvent(t, []string{"pipelinerun", "finished"}, "cd.pipelinerun.finished.v1")
	executeAndVerifyEvent(t, []string{"pipelinerun", "queued"}, "cd.pipelinerun.queued.v1")
}

func TestRepositoryEvent(t *testing.T) {
		executeAndVerifyEvent(t, []string{"repository", "created"}, "cd.repository.created.v1")
		executeAndVerifyEvent(t, []string{"repository", "deleted"}, "cd.repository.deleted.v1")
		executeAndVerifyEvent(t, []string{"repository", "modified"}, "cd.repository.modified.v1")
}

func TestServiceEvent(t *testing.T) {
		executeAndVerifyEvent(t, []string{"service", "deployed"}, "cd.service.deployed.v1")
		executeAndVerifyEvent(t, []string{"service", "upgraded"}, "cd.service.upgraded.v1")
		executeAndVerifyEvent(t, []string{"service", "removed"}, "cd.service.removed.v1")
		executeAndVerifyEvent(t, []string{"service", "rolledback"}, "cd.service.rolledback.v1")
}

func TestTaskRunEvent(t *testing.T) {
	executeAndVerifyEvent(t, []string{"taskrun", "started"}, "cd.taskrun.started.v1")
	executeAndVerifyEvent(t, []string{"taskrun", "finished"}, "cd.taskrun.finished.v1")
}


func executeAndVerifyEvent(t *testing.T, args []string, eventType string) {
	var receivedEvent *event.Event
	okResponse := func(res http.ResponseWriter, req *http.Request) {
		message := cehttp.NewMessageFromHttpRequest(req)
		receivedEvent, _ = binding.ToEvent(context.Background(), message)
	}

	testServer := httptest.NewServer(http.HandlerFunc(okResponse))
	defer func() { testServer.Close() }()

	CDE_SINK = testServer.URL
	rootCmd.SetArgs(args)
	_, err := rootCmd.ExecuteC()
	assert.NoError(t, err)
	t.Log(receivedEvent)
	assert.NoError(t, cetest.IsValid()(*receivedEvent))
	assert.NoError(t, cetest.HasType(eventType)(*receivedEvent))
}
