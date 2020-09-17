package operations

import (
	"errors"
	"fmt"
	"github.com/dylancorbus/go-database/internal/app/models"
	"github.com/dylancorbus/go-database/internal/pkg/constants"
	constants2 "github.com/dylancorbus/go-database/internal/pkg/index/constants"
	"github.com/dylancorbus/go-database/internal/pkg/index/operations"
	"log"
	"os"
	"strconv"
	"strings"
)

func Create(name string, salary int) error {
	var emp = models.Employee{Name: name, Salary: salary}
	if constants2.Indexes[name] != 0 {
		return errors.New("record already exists")
	}
	// If the file doesn't exist, create it, or append to the file
	txn, err := os.OpenFile(constants.Transaction, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := []string{emp.Name, strconv.Itoa(emp.Salary), strconv.FormatBool(emp.Vacation)}
	v := strings.Join(s, ";")
	//write to file
	txn.WriteString(v + "\n")
	operations.CreateIndex(name, len(constants2.Indexes)+1)
	txn.Close()
	//create transaction log
	fmt.Println("Writing to file ", emp)
	return nil
}
