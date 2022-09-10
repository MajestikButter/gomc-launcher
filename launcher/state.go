package launcher

import (
	"encoding/json"
	"os"
)

type State struct {
	ReleaseName string `json:"releaseName"`
	PreviewName string `json:"previewName"`
	WindowSize  *Vec2  `json:"windowSize"`
}

func (s *State) Load() error {
	data, err := os.ReadFile(STATE_FILE)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, s)
}

func (s *State) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(STATE_FILE, data, FILE_MODE)
}
