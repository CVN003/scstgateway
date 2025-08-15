package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CVN003/scstgateway/utils"
)

func (t *Target) AddTarget() error {

	if t.TargetType == "iscsi" {
		//todo
		return nil
	}
	if t.TargetType == "qla2x00t" {
		l.Warnf("[AddTarget]:qlax00t:%s ignore!", t.Name)
		return nil
	}
	return fmt.Errorf("[AddTarget]:%s not support", t.TargetType)

}

func (t *Target) AddGroup(grpname string) error {
	if _, err := os.Stat(filepath.Join(TargetPath, t.TargetType, t.Name, "ini_groups", grpname)); err == nil {
		l.Warnf("[AddGroup]:%s:%s already exist!", t.Name, grpname)
		return nil
	}
	cmd := "create " + grpname
	if err := utils.WriteScstSysfs(filepath.Join(TargetPath, t.TargetType, t.Name, "ini_groups/mgmt"), cmd, 10); err != nil {
		l.Errorf("[AddGroup]:%s:%s create failed:%v", t.Name, grpname, err)
		return fmt.Errorf("[AddGroup]:%s:%s create failed:%v", t.Name, grpname, err)
	}
	l.Debugf("[AddGroup]:%s/%s create success!", t.Name, grpname)
	return nil

}
