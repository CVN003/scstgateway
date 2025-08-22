package core

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/CVN003/scstgateway/utils"
)

func GetLiveConfig() (SCSTMapping, error) {
	config := SCSTMapping{}

	if iscsi_tgt, err := getTargetMap(ISCSI_Target); err != nil {
		return config, err
	} else {
		config.Targets = append(config.Targets, iscsi_tgt...)
	}
	if qla2x00t_tgt, err := getTargetMap(QLA2x00T_Target); err != nil {
		// ignore qla2x00t target not found error
		return config, nil
	} else {
		config.Targets = append(config.Targets, qla2x00t_tgt...)
	}

	return config, nil
}

func getTargetMap(path string) ([]Target, error) {
	tgtGrp := []Target{}
	tg, err := utils.ListCurDirAbsPath(path)
	if err != nil {
		return []Target{}, err
	}
	for _, t := range tg {
		tgt := Target{}
		iniGrp := []IniGroup{}
		dir, _ := os.ReadDir(filepath.Join(t, "ini_groups"))
		for _, grp := range dir {
			if grp.Name() == "mgmt" {
				continue
			}
			ig := IniGroup{}
			ig.GrpName = grp.Name()
			ini_groups, _ := os.ReadDir(filepath.Join(t, "ini_groups", grp.Name(), "initiators"))
			for _, acl := range ini_groups {
				if acl.Name() == "mgmt" {
					continue
				}
				ig.Ini = append(ig.Ini, acl.Name())
			}
			dir, _ = os.ReadDir(filepath.Join(t, "ini_groups", grp.Name(), "luns"))
			for _, lun := range dir {
				tmpLUN := Lun{}
				if lun.Name() == "mgmt" {
					continue
				}
				if _, err := os.Stat(filepath.Join(t, "ini_groups", grp.Name(), "luns", lun.Name(), "device/dummy")); err != nil {
					tmpLUN.LunName, _ = utils.ReadFirstLine(filepath.Join(t, "ini_groups", grp.Name(), "luns", lun.Name(), "device/pr_file_name"))
					tmpLUN.LunName = filepath.Base(tmpLUN.LunName)
					tmpLUN.LunDeviceFullPath, _ = utils.ReadFirstLine(filepath.Join(t, "ini_groups", grp.Name(), "luns", lun.Name(), "device/filename"))

				} else {
					tmpLUN.LunName = "dummy"
					tmpLUN.LunDeviceFullPath = ""
				}
				tmpLUN.LunID, _ = strconv.Atoi(lun.Name())
				ig.Lun = append(ig.Lun, tmpLUN)
			}
			iniGrp = append(iniGrp, ig)
		}
		tgt.IniGroup = iniGrp
		tgt.Name = filepath.Base(t)
		tgt.TargetType = filepath.Base(path)
		tgtGrp = append(tgtGrp, tgt)

	}

	return tgtGrp, nil
}
