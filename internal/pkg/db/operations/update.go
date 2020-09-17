package operations

import (
	"fmt"
	"github.com/dylancorbus/go-database/internal/pkg/constants"
	operations2 "github.com/dylancorbus/go-database/internal/pkg/file/operations"
	constants2 "github.com/dylancorbus/go-database/internal/pkg/index/constants"
	"github.com/dylancorbus/go-database/internal/pkg/index/operations"
	"strconv"
	"strings"
)

func Update(id string, field string, value string) {
	emp, err := Read(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch field {
	case "salary":
		fmt.Println("Changed ", emp.Salary, field, "to", value)
		emp.Salary, _ = strconv.Atoi(value)
	case "name":
		fmt.Println("Changed ", emp.Name, field, "to", value)
		operations.UpdateIndex(emp.Name, value)
		emp.Name = value
	case "vacation":
		fmt.Println("Changed ", emp.Name, field, "to", value)
		emp.Vacation, _ = strconv.ParseBool(value)
	}
	s := []string{emp.Name, strconv.Itoa(emp.Salary), strconv.FormatBool(emp.Vacation)}
	v := strings.Join(s, ";")
	operations2.UpdateLine(constants.Transaction, constants2.Indexes[emp.Name]-1, v)
}