package operations

import (
	"io/ioutil"
	"os"
	"strings"
)

func DeleteLine(path string, lineNumber int) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	info, _ := os.Stat(path)
	mode := info.Mode()
	array := strings.Split(string(file), "\n")
	array = append(array[:lineNumber-1], array[lineNumber:]...)
	ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
	return array
}
