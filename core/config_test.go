package core

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"
)

func Test_GetLiveConfig(t *testing.T) {
	config, _ := getTargetMap(filepath.Join(ISCSI_Target))
	jsonStr, _ := json.Marshal(config)
	fmt.Printf("jsonStr: %v\n", string(jsonStr))
}
