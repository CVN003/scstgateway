package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// func isSysfsPathExist(path string, l *zap.SugaredLogger) bool {
// 	_, err := os.Stat(path)
// 	if err != nil {
// 		l.Debugf("path:%v not exist!,err:%v", path, err)
// 		return false
// 	} else {
// 		return true
// 	}
// }

func ReadFirstLine(fp string) (string, error) {
	f, err := os.Open(fp)
	if err == nil {
		fr := bufio.NewReader(f)
		linedata, _, err := fr.ReadLine()
		if err == nil {
			return string(linedata), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func WriteScstSysfs(path, cmd string, tmout int) error {
	if tmout > 10 {
		tmout = 10
	}
	ret := make(chan error)
	go func() {
		if err := os.WriteFile(path, []byte(cmd), 0600); err != nil {
			ret <- err
		} else {
			ret <- nil
		}
	}()
	select {
	case r := <-ret:
		if r != nil {
			return r
		} else {
			return nil
		}
	case <-time.After(time.Duration(tmout) * time.Second):
		return fmt.Errorf("WriteSCSTSysfs timeout")
	}
}

func ListCurDirAbsPath(path string) ([]string, error) {
	var RetFiles []string
	fl, err := os.Open(path)
	if err != nil {
		return nil, err
	} else {
		allcurdir, _ := fl.Readdirnames(0)
		for _, dirname := range allcurdir {
			ret, _ := IsDir(filepath.Join(path, dirname))
			if ret {
				RetFiles = append(RetFiles, path+"/"+dirname)
			}
		}

	}
	return RetFiles, nil
}

func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}
