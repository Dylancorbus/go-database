package operations

import (
	"io/ioutil"
	"os"
	"strings"
)

func UpdateLine(path string, lineNumber int, update string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	info, _ := os.Stat(path)
	mode := info.Mode()
	array := strings.Split(string(file), "\n")
	array[lineNumber] = update
	ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
}
