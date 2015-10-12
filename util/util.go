package util

import (
	"os"
	"path/filepath"
)

// 获取程序运行的目录
func GetDir() (string, error) {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(path), nil
}

// 递归获取一个目录的子目录
func RecursiveDir(dir string, l []string) ([]string, error) {
	dl, err := ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, d := range dl {
		if d.IsDir() {
			_dir := dir + "/" + d.Name()
			l = append(l, _dir)
			l, err = RecursiveDir(_dir, l)
			if err != nil {
				return l, err
			}
		}
	}
	return l, err
}

// 读取一个文件夹返回文件列表
func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}
