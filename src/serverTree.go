package main
import (
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
	"crypto/md5"
)

func TraverseDir(dir string) ServerFileTree {
	var sft ServerFileTree
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() != true {
			buf, _ := ioutil.ReadFile(path)
			h := md5.New()
			h.Write(buf)
			hashitself := h.Sum(nil)
			sft.Files = append(sft.Files, file{path, string(hashitself), info.ModTime()})
		}
		return nil
	})
	return sft
}

func Somethin () {
	fmt.Println("FUCK OFF GO!!!")
}
