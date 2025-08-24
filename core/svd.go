package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/CVN003/scstgateway/utils"
)

func (svd *SVD) Add() error {
	if err := svd.setMGMT(); err != nil {
		return err
	}
	if err := svd.addDevice(); err != nil {
		return err
	}
	if err := svd.SetVID(); err != nil {
		return err
	}
	if err := svd.SetPID(); err != nil {
		return err
	}
	if err := svd.SetUSN(); err != nil {
		return err
	}
	return nil
}

func (svd *SVD) setMGMT() error {
	if svd.Handler == "vdisk_blockio" || svd.Handler == "vdisk_fileio" {
		svd.mgmt = HandlerPath + svd.Handler + "/mgmt"
	} else {
		return fmt.Errorf("handler:%s not support!", svd.Handler)
	}
	return nil
}

func (svd *SVD) addDevice() error {
	cmd := "add_device " + svd.LunName + " filename=" + svd.LunDeviceFullPath
	if err := utils.WriteScstSysfs(svd.mgmt, cmd, 10); err != nil {
		l.Errorf("[AddDevice]:%s,path:%s failed:%v", svd.LunName, svd.LunDeviceFullPath, err)
		return fmt.Errorf("[AddDevice]:%s,path:%s failed:%v", svd.LunName, svd.LunDeviceFullPath, err)
	}
	l.Debugf("[AddDevice]:%s,path:%s success!", svd.LunName, svd.LunDeviceFullPath)
	return nil

}

func (svd *SVD) SetVID() error {
	if err := utils.WriteScstSysfs(filepath.Join(HandlerPath, svd.Handler, svd.LunName, "t10_vend_id"), svd.VID, 10); err != nil {
		l.Errorf("[SetVID]:%s,vid:%s failed:%v", svd.LunName, svd.VID, err)
		return fmt.Errorf("[SetVID]:%s,vid:%s failed:%v", svd.LunName, svd.VID, err)
	}
	l.Debugf("[SetVID]:%s,%s success!", svd.LunName, svd.VID)
	return nil
}

func (svd *SVD) SetPID() error {
	if err := utils.WriteScstSysfs(filepath.Join(HandlerPath, svd.Handler, svd.LunName, "prod_id"), svd.PID, 10); err != nil {
		l.Errorf("[SetPID]:%s,pid:%s failed:%v", svd.LunName, svd.PID, err)
		return fmt.Errorf("[SetPID]:%s,pid:%s failed:%v", svd.LunName, svd.PID, err)
	}
	l.Debugf("[SetPID]:%s,%s success!", svd.LunName, svd.PID)
	return nil
}

func (svd *SVD) SetUSN() error {
	if err := utils.WriteScstSysfs(filepath.Join(HandlerPath, svd.Handler, svd.LunName, "usn"), svd.USN, 10); err != nil {
		l.Errorf("[SetUSN]:%s,usn:%s failed:%v", svd.LunName, svd.USN, err)
		return fmt.Errorf("[SetUSN]:%s,usn:%s failed:%v", svd.LunName, svd.USN, err)
	}
	l.Debugf("[SetUSN]:%s,%s success!", svd.LunName, svd.USN)
	return nil
}

func (svd *SVD) ResyncDeviceSize() error {
	if err := utils.WriteScstSysfs(filepath.Join(HandlerPath, svd.Handler,
		svd.LunName, "resync_size"), "1", 10); err != nil {
		l.Errorf("[ResyncDeviceSize]:%s,size:%s failed:%v", svd.LunName, svd.Size, err)
		return fmt.Errorf("[ResyncDeviceSize]:%s,size:%v failed:%v", svd.LunName, svd.Size, err)
	}
	l.Debugf("[ResyncDeviceSize]:%s,size:%s success!", svd.LunName, svd.Size)
	return nil

}

func (svd *SVD) Remove() error {
	if err := svd.setMGMT(); err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(HandlerPath, svd.Handler, svd.LunName)); err != nil {
		l.Debugf("[RemoveDeivce]:%s not exist!", svd.LunName)
		return nil
	}
	cmd := "del_device " + svd.LunName
	if err := utils.WriteScstSysfs(svd.mgmt, cmd, 10); err != nil {
		l.Debugf("[RemoveDeivce]:%s,cmd:%s failed:%v,need recheck hold 1 second!", svd.LunName, cmd, err)
		time.Sleep(time.Second * 1)
		if _, err1 := os.Stat(filepath.Join(svd.mgmt, svd.LunName)); err1 != nil {
			l.Debugf("[RemoveDeivce]:%s success!", svd.LunName)
			return nil
		} else {
			return fmt.Errorf("[RemoveDeivce]:%s,cmd:%s failed:%v", svd.LunName, cmd, err)
		}
	}
	l.Debugf("[RemoveDeivce]:%s success!", svd.LunName)
	return nil
}
