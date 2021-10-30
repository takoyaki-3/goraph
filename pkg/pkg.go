package pkg

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func FindFiles(path string,extension string)[]string{
	files := []string{}
	paths,_ := Dirwalk(path)
	for _,v:=range paths{
		if strings.HasSuffix(v,extension) {
			files = append(files, v)
		}
	}
	return files
}

func Dirwalk(dir string) ([]string, []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths, file_names []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
		file_names = append(file_names, file.Name())
	}
	return paths, file_names
}
