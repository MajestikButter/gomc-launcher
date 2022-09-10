package launcher

import (
	"encoding/json"
	"os"
)

type Settings struct {
	KeepOpen      bool   `json:"keepOpen"`
	LogsDirectory string `json:"logsDirectory"`
}

func (l *Settings) Load() error {
	data, err := os.ReadFile(SETTINGS_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, l)
}

func (l *Settings) Save() error {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(SETTINGS_FILE, data, FILE_MODE)
}
