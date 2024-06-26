package upgradeinspectorv4

import (
	"encoding/json"
	"fmt"

	"github.com/chef/automate/components/automate-deployment/pkg/cli"
	"github.com/chef/automate/components/automate-deployment/pkg/inspector"
	"github.com/chef/automate/lib/io/fileutils"
	"github.com/chef/automate/lib/majorupgrade_utils"
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

type StoreESSettingsInspection struct {
	writer          *cli.Writer
	upgradeUtils    UpgradeV4Utils
	timeout         int64
	exitError       error
	exitedWithError bool
	fileutils       fileutils.FileUtils
}

func (ses *StoreESSettingsInspection) ShowInfo(index *int) error {
	return nil
}

func (ses *StoreESSettingsInspection) Skip() {
	return
}
func (ses *StoreESSettingsInspection) GetShortInfo() []string {
	return nil
}

func NewStoreESSettingsInspection(w *cli.Writer, utls UpgradeV4Utils, fileutils fileutils.FileUtils, timeout int64) *StoreESSettingsInspection {
	return &StoreESSettingsInspection{
		writer:       w,
		upgradeUtils: utls,
		timeout:      timeout,
		fileutils:    fileutils,
	}
}

func (ses *StoreESSettingsInspection) GetInstallationType() inspector.InstallationType {
	return inspector.EMBEDDED
}

func (ses *StoreESSettingsInspection) Inspect() error {
	esSetting, err := ses.getESSettings()
	if err != nil {
		ses.setExitError(err)
		return err
	}
	err = ses.storeESSettings(esSetting)
	if err != nil {
		ses.setExitError(err)
		return err
	}
	return nil
}

func (ses *StoreESSettingsInspection) setExitError(err error) {
	ses.exitError = err
	ses.exitedWithError = true
}

type ESSettings struct {
	TotalShardSettings int64 `json:"totalShardSettings" toml:"totalShardSettings"`
}

type IndicesShardTotal struct {
	Indices struct {
		Shards struct {
			Total int64 `json:"total"`
		} `json:"shards"`
	} `json:"indices"`
}

func (ses *StoreESSettingsInspection) getAllSearchEngineSettings() (*ESSettings, error) {
	esSettings := &ESSettings{}
	totalShard, err := ses.getTotalShards()
	if err != nil {
		esSettings.TotalShardSettings = majorupgrade_utils.INDICES_TOTAL_SHARD_DEFAULT
	} else {
		esSettings.TotalShardSettings = totalShard.Indices.Shards.Total
	}
	return esSettings, err
}

func (ses *StoreESSettingsInspection) getTotalShards() (*IndicesShardTotal, error) {
	indicesShardTotal := &IndicesShardTotal{}
	basePath := ses.upgradeUtils.GetESBasePath(ses.timeout)
	totalShard, err := ses.upgradeUtils.ExecRequest(basePath+"_cluster/stats?filter_path=indices.shards.total", "GET", nil)
	if err != nil {
		return indicesShardTotal, errors.Wrap(err, "Failed to fetch total shards")
	}
	err = json.Unmarshal(totalShard, indicesShardTotal)
	if err != nil {
		return indicesShardTotal, errors.Wrap(err, "failed to unmarshal total shards")
	}
	return indicesShardTotal, nil
}

func (ses *StoreESSettingsInspection) getESSettings() (*ESSettings, error) {
	esSettings, err := ses.getAllSearchEngineSettings()
	if err != nil {
		return esSettings, err
	}
	return esSettings, nil
}

func (ses *StoreESSettingsInspection) storeESSettings(esSettings *ESSettings) error {
	esSettingsJson, err := json.Marshal(esSettings)
	if err != nil {
		return errors.Wrap(err, "Failed to map elasticsearch settings to json")
	}
	habRoot := ses.fileutils.GetHabRootPath()
	esSettingFilePath := habRoot + majorupgrade_utils.V3_ES_SETTING_FILE
	err = ses.fileutils.WriteToFile(esSettingFilePath, esSettingsJson)
	if err != nil {
		return errors.Wrap(err, "Failed to write elasticsearch settings to file")
	}
	return nil
}

func (ses *StoreESSettingsInspection) ExitHandler() error {
	if ses.exitedWithError {
		ses.writer.Println(fmt.Errorf("["+color.New(color.FgRed).Sprint("Error")+"] %w", ses.exitError).Error())
	}
	return nil
}
