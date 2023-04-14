package cmd

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GetFileList(path string) (*[]string, error) {
	var list []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else if info.Mode().IsRegular() {
				list = append(list, path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func ShuffleFileList(list *[]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len((*list)), func(i, j int) {
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	})
}
