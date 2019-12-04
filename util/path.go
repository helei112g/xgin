package util

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取绝对路径
func AbsPath(path string) string {
	if strings.HasPrefix(path, "$HOME") {
		path = userHomeDir() + path[5:]
	}

	if strings.HasPrefix(path, "$") {
		end := strings.Index(path, string(os.PathSeparator))
		if end == -1 {
			return ""
		}

		path = os.Getenv(path[1:end]) + path[end:]
	}

	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}

	if apath, err := filepath.Abs(path); err == nil {
		return filepath.Clean(apath)
	}

	return ""
}

// 获取用户家目录
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
