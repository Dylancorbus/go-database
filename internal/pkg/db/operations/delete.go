package operations

import (
	"github.com/dylancorbus/go-database/internal/pkg/constants"
	"github.com/dylancorbus/go-database/internal/pkg/file/operations"
	constants2 "github.com/dylancorbus/go-database/internal/pkg/index/constants"
	operations2 "github.com/dylancorbus/go-database/internal/pkg/index/operations"
	"strconv"
	"strings"
)

func DeleteItem(id string) {
	arr := operations2.Delete(id)
	operations.DeleteLine(constants.Transaction, constants2.Indexes[id])
	for i := 0; i < len(arr); i++ {
		textArr := strings.Split(arr[i], ";")
		line := i + 1
		if len(textArr) == 1 {
			break
		}
		if str := strconv.Itoa(line); str == textArr[1] {
			continue
		} else {
			textArr[1] = str
		}
		constants2.Indexes[textArr[0]] = line
		operations.UpdateLine(constants.Index, i, strings.Join(textArr, ";"))
	}
}