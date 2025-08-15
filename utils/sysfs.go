package utils

import (
	"bufio"
	"fmt"
	"os"
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
