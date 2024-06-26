package backup

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	api "github.com/chef/automate/api/interservice/deployment"
)

// TestBackup tests the backup execution
func TestBackup(t *testing.T) {
	t.Run("with failing operations", func(t *testing.T) {
		exec, bctx, _ := testBackupExecutor(t, testFailSyncSpec)

		err := exec.Backup(bctx)

		require.Error(t, err, "returns an error of a sync operation fails")
		assert.Contains(t, "test operation failed", err.Error())
	})

	t.Run("with successful operations", func(t *testing.T) {
		exec, bctx, _ := testBackupExecutor(t, testSuccessSpec)

		require.NoError(t, exec.Backup(bctx))
	})

	t.Run("timeout doesn't hang", func(t *testing.T) {
		exec, bctx, _ := testBackupExecutorWithTimeout(t, testSuccessSpec, 0)

		err := exec.Backup(bctx)
		require.Error(t, err, "a timeout has happened")
		assert.Contains(t, err.Error(), "context deadline exceeded")
	})
}

func TestProgressCalculator(t *testing.T) {
	syncCal := NewProgressCalculator()
	svcs := []string{"service-a", "service-b", "service-c"}

	for _, s := range svcs {
		syncCal.Update(OperationProgress{Name: s, Progress: 0})
	}

	syncCal.Update(OperationProgress{Name: "service-a", Progress: 80})
	syncCal.Update(OperationProgress{Name: "service-b", Progress: 40})
	assert.Equal(t, float64(40), syncCal.Percent())
}

func testLocationSpec(tmpdir string) LocationSpecification {
	return FilesystemLocationSpecification{
		Path: tmpdir,
	}
}

func testBackupExecutorWithTimeout(t testing.TB, spec Spec, timeout time.Duration) (*Executor, Context, chan api.DeployEvent_Backup_Operation) {
	// buffer the event channel so we don't block on events
	eventChan := make(chan api.DeployEvent_Backup_Operation, 100)
	errChan := make(chan error)
	tmpDir := t.TempDir()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	parsed, _ := time.Parse(api.BackupTaskFormat, "20060102150405")
	ts, _ := ptypes.TimestampProto(parsed)
	bctx := Context{
		ctx:          ctx,
		backupTask:   &api.BackupTask{Id: ts},
		locationSpec: testLocationSpec(tmpDir),
	}

	return NewExecutor(
		WithEventChan(eventChan),
		WithErrorChan(errChan),
		WithSpec(spec),
		WithCancel(cancel),
	), bctx, eventChan
}

func testBackupExecutor(t testing.TB, spec Spec) (*Executor, Context, chan api.DeployEvent_Backup_Operation) {
	return testBackupExecutorWithTimeout(t, spec, 2*time.Second)
}

var testSuccessSpec = Spec{
	Name: "service-a",
	testSyncOps: []testOperation{
		{name: "sync-op-1"},
		{name: "sync-op-2"},
		{name: "sync-op-3"},
	},
}

var testFailSyncSpec = Spec{
	Name: "service-a",
	testSyncOps: []testOperation{
		{name: "sync-op-1"},
		{
			name: "sync-op-2-fail",
			fail: true,
		},
		{name: "sync-op-3"},
	},
}
