package operations

import (
	"bufio"
	"github.com/dylancorbus/go-database/internal/pkg/constants"
	"github.com/dylancorbus/go-database/internal/pkg/file/operations"
	constants2 "github.com/dylancorbus/go-database/internal/pkg/index/constants"

	"os"
	"strings"
)

func UpdateIndex(old string, new string) {
	file, _ := os.Open(constants.Index)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for i := 0; scanner.Scan(); i++ {
		if i != constants2.Indexes[old]-1 {
			continue
		}
		x := scanner.Text()
		if arr := strings.Split(x, ";"); arr[0] == old {
			arr[0] = new
			s := strings.Join(arr, ";")
			operations.UpdateLine(constants.Index, i, s)
			break
		}
	}
	x := constants2.Indexes[old]
	delete(constants2.Indexes, old)
	constants2.Indexes[new] = x
	file.Close()
}
