package internal

import (
	"os"
	"path"
	"strconv"
)

//ReadSymlink check if symlink does exist
func readSymlink(link filepath) (bool, error) {
	// FIXME: currently is not used, implement part where
	// the number of opened fd-s only will be checking
	if _, err := os.Stat(link); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

//readDir list files in directory
func readDir(dir filepath) ([]os.DirEntry, error) {
	dirFiles, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return dirFiles, nil
}

//mkPidFdPath construct filepath to fd dir from pid
func mkPidFdPath(id pid) filepath {
	return path.Join("/proc/", strconv.Itoa(id), "/fd/")
}

//GetFdNum get number of all fd-s associated with given pid
func GetFdNum(id pid) (int, error) {
	dirFiles, err := readDir(mkPidFdPath(id))
	if err != nil {
		//TODO: check if process exists
		return 0, err
	}
	return len(dirFiles), nil
}
