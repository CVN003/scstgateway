package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/CVN003/scstgateway/utils"
)

func AddDummy(tgt, tgt_type, ini_grp string) error {
	cmd := "add dummy 0"
	if err := utils.WriteScstSysfs(filepath.Join(TargetPath, tgt_type, tgt, "ini_groups", ini_grp, "luns/mgmt"), cmd, 10); err != nil {
		l.Errorf("[AddDummy]:%s/%s add failed:%v", tgt, ini_grp, err)
		return fmt.Errorf("[AddDummy]:%s/%s add failed:%v", tgt, ini_grp, err)
	}
	l.Debugf("[AddDummy]:%s/%s success!", tgt, ini_grp)
	return nil
}

func AddLun2Grp(lun, tgt, tgt_type, ini_grp string) error {
	lunList, err := os.ReadDir(filepath.Join(TargetPath, tgt_type, tgt, "ini_groups", ini_grp, "luns"))
	if err != nil {
		l.Errorf("[AddLun2Grp]:%s/%s read ini_group failed:%v", tgt, ini_grp, err)
		return fmt.Errorf("[AddLun2Grp]:%s/%s read ini_group failed:%v", tgt, ini_grp, err)
	}
	l.Debugf("[AddLun2Grp]:%s/%s Listing all attacted luns", tgt, ini_grp)
	for i, lun := range lunList {
		l.Debugf("[AddLun2Grp]:%d/%s", i, lun.Name())
	}

	// 0号lun(目录)等于mgmt时，意味这ini_grp没有任何lun attached
	if lunList[0].Name() == "mgmt" {
		if err := AddDummy(tgt, tgt_type, ini_grp); err != nil {
			return err
		}
		//首个lunid 从1开始
		CurLunID := "1"
		cmd := "add " + lun + " " + CurLunID
		if err := utils.WriteScstSysfs(filepath.Join(TargetPath, tgt_type, tgt, "ini_groups", ini_grp, "luns/mgmt"), cmd, 10); err != nil {
			l.Errorf("[AddLun2Grp]:%s/%s/%s add failed:%v", tgt, ini_grp, lun, err)
			return fmt.Errorf("[AddLun2Grp]:%s/%s/%s add failed:%v", tgt, ini_grp, lun, err)
		}
		l.Errorf("[AddLun2Grp]:%s/%s/%s add success!", tgt, ini_grp, lun)
		return nil
	}
	var LunIDGrp []int
	for _, n := range lunList {
		if n.Name() != "mgmt" {
			i, _ := strconv.ParseInt(n.Name(), 10, 64)
			LunIDGrp = append(LunIDGrp, int(i))
		}
	}
	sort.Ints(LunIDGrp)
	l.Debugf("[AddLun2Grp]:%s/%s Current LunIDGrp:%v", tgt, ini_grp, LunIDGrp)
	// 按大小排序后，取最后一个的值再+1
	NewLunID := LunIDGrp[len(LunIDGrp)-1] + 1
	cmd := "add " + lun + " " + strconv.Itoa(NewLunID)
	if err := utils.WriteScstSysfs(filepath.Join(TargetPath, tgt_type, tgt, "ini_groups", ini_grp, "luns/mgmt"), cmd, 10); err != nil {
		l.Errorf("[AddLun2Grp]:%s/%s/%s add failed:%v", tgt, ini_grp, lun, err)
		return fmt.Errorf("[AddLun2Grp]:%s/%s/%s add failed:%v", tgt, ini_grp, lun, err)
	}
	l.Errorf("[AddLun2Grp]:%s/%s/%s add success!", tgt, ini_grp, lun)
	return nil

}

func (t *Target) AddIni2Group(ini, ini_group string) error {
	if _, err := os.Stat(filepath.Join(TargetPath, t.TargetType, t.Name, "ini_groups", ini_group, "initiators", ini)); err != nil {
		cmd := "add " + ini
		if err := utils.WriteScstSysfs(filepath.Join(TargetPath, t.TargetType, t.Name, "ini_groups", ini_group, "initiators/mgmt"), cmd, 10); err != nil {
			l.Errorf("[AddIni2Group]:%s/%s/%s add failed:%v", t.Name, ini_group, ini, err)
			return fmt.Errorf("[AddIni2Group]:%s/%s/%s add failed:%v", t.Name, ini_group, ini, err)
		}
		l.Debugf("[AddIni2Group]:%s/%s/%s add success!", t.Name, ini_group, ini)
		return nil
	}
	l.Debugf("[AddIni2Group]:%s/%s/%s existed,skip!", t.Name, ini_group, ini)
	return nil
}
