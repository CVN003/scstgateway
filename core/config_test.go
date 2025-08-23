package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test_GetLiveConfig(t *testing.T) {
	config, _ := getTargetMap(filepath.Join(ISCSI_Target))

	scstmap := &SCSTMapping{}
	scstmap.Targets = config
	jsonStr, _ := json.Marshal(scstmap)
	fmt.Printf("jsonStr: %v\n", string(jsonStr))
	os.WriteFile("/root/scst.json", jsonStr, 0644)
}
