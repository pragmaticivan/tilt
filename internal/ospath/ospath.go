package ospath

import (
	"os"
	"path/filepath"
)

// Given absolute paths `dir` and `file`, returns
// the relative path of `file` relative to `dir`.
//
// Returns true if successful. If `file` is not under `dir`, returns false.
func Child(dir string, file string) (string, bool) {
	current := file
	child := ""
	for true {
		if dir == current {
			return child, true
		}

		if len(current) <= len(dir) || current == "." {
			return "", false
		}

		cDir := filepath.Dir(current)
		cBase := filepath.Base(current)
		child = filepath.Join(cBase, child)
		current = cDir
	}

	return "", false
}

func RealChild(dir string, file string) (string, bool, error) {
	realDir, err := RealAbs(dir)
	if err != nil {
		return "", false, err
	}
	realFile, err := RealAbs(file)
	if err != nil {
		return "", false, err
	}

	rel, isChild := Child(realDir, realFile)
	return rel, isChild, nil
}

// Returns the absolute version of this path, resolving all symlinks.
func RealAbs(path string) (string, error) {
	// Make the path absolute first, so that we find any symlink parents.
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// Resolve the symlinks.
	realPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		return "", err
	}

	// Double-check we're still absolute.
	return filepath.Abs(realPath)
}

// Like os.Getwd, but with all symlinks resolved.
func Realwd() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return RealAbs(path)
}

func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}

	return f.Mode().IsDir()
}