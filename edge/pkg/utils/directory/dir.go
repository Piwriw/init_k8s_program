package directory

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Directory struct {
}

func NewDirC() *Directory {
	return &Directory{}
}

func (d *Directory) GetDirEntryNum(DirPath string) (int, error) {
	file, err := os.Open(DirPath)
	if err != nil {
		return 0, err
	}
	fileStat, err := file.Stat()
	if err != nil {
		if os.IsNotExist(err) {
			return 0, errors.New(fmt.Sprintf("Dir does not exist, Error: %v", err))
		}
		return 0, err
	}
	if !fileStat.IsDir() {
		return 0, errors.New(fmt.Sprintf("%s is not dir", DirPath))
	}
	dirEntries, err := file.ReadDir(-1)
	if err != nil {
		return 0, err
	}
	//for k, v := range dirEntries {
	//	klog.Infoln(k, " dir Entry: ", v.Name())
	//}
	return len(dirEntries), nil
}

func (d *Directory) GetDirEntryPath(ParentPath string) (absArr []string, err error) {
	file, err := os.Open(ParentPath)
	if err != nil {
		return nil, err
	}
	fileStat, err := file.Stat()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New(fmt.Sprintf("Dir does not exist, Error: %v", err))
		}
		return nil, err
	}
	if !fileStat.IsDir() {
		return nil, errors.New(fmt.Sprintf("%s is not dir", ParentPath))
	}
	dirEntries, err := file.ReadDir(-1)
	if err != nil {
		return nil, err
	}
	for _, v := range dirEntries {
		abs, err := filepath.Abs(filepath.Join(file.Name(), v.Name()))
		if err != nil {
			return nil, err
		}
		absArr = append(absArr, abs)
	}

	return absArr, nil
}

// Create creates a new file at 'path' and sets permissions defined
// by 'mode'.
func Create(path string, mode os.FileMode) (*os.File, error) {

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	err = os.Chmod(path, mode)
	if err != nil {
		os.Remove(path)
		return nil, err
	}

	return f, nil
}

// Mkdir create a new directory at 'path' and sets permissions defined
// by 'mode'.
func Mkdir(path string, mode os.FileMode) error {

	err := os.Mkdir(path, mode)
	if err != nil {
		return err
	}

	// Go1.10 seems to disregard the 'mode' argument...
	err = os.Chmod(path, mode)
	if err != nil {
		// os.Remove(path)
		return err
	}

	return nil
}

// MkdirAll creates all missing directories in a provided paths, and chmods them.
func MkdirAll(path string, mode os.FileMode) error {

	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	abs = filepath.ToSlash(abs)
	elements := strings.Split(abs, "/")

	cumulativePath := "/"

	for _, e := range elements {

		cumulativePath = filepath.Join(cumulativePath, e)

		if cumulativePath != "" {
			if _, err := os.Stat(cumulativePath); os.IsNotExist(err) {
				err := Mkdir(cumulativePath, mode)
				if err != nil {
					return err
				}

				err = os.Chmod(cumulativePath, mode)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}
