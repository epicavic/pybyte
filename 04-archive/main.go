package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func makeArchive(target string, sources []string) error {
	for _, source := range sources {
		_, err := os.Stat(source)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(target)
	if err != nil {
		return err
	}
	defer file.Close()

	archive := zip.NewWriter(file)
	defer archive.Close()

	var baseDir string
	for _, source := range sources {
		info, err := os.Stat(source)
		if err != nil {
			return err
		}
		if info.IsDir() {
			baseDir = filepath.Base(source)
		}
		err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	source_dir := []string{"./source"}
	target_dir := "./target"
	arcname := filepath.Join(target_dir, time.Now().Format("2006-01-02_15:04:05")+".zip")
	if err := makeArchive(arcname, source_dir); err != nil {
		fmt.Println("Backup FAILED")
		os.Exit(1)
	}
	fmt.Println("Successful backup to", arcname)
}
