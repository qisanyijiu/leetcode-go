package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cnt := 0
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil{
			return err
		}
		if strings.Contains(path,".go")&& !strings.Contains(path, "_test"){
			cnt ++
		}
		return nil
	})
	if err != nil{
		log.Println(err)
	}
	log.Println(cnt)
}
