package operations

import (
	"github.com/dylancorbus/go-database/internal/pkg/constants"
	"github.com/dylancorbus/go-database/internal/pkg/file/operations"
	constants2 "github.com/dylancorbus/go-database/internal/pkg/index/constants"
)

func Delete(id string) []string {
	arr := operations.DeleteLine(constants.Index, constants2.Indexes[id])
	delete(constants2.Indexes, id)
	return arr
}
