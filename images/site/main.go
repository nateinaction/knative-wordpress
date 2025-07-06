package main

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	err := copyDir("/tmp/wordpress", "/var/www/html")
	if err != nil {
		panic(err)
	}

	time.Sleep(9223372036 * time.Second)
}

// dst is the destination directory path where the contents of the source directory (src)
// will be copied. If the directory does not exist, it will be created along with any
// necessary parent directories. Existing files in the destination may be overwritten.
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		dstFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer dstFile.Close()
		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}
