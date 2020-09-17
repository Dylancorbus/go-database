package operations

import (
	"bufio"
	"fmt"
	"github.com/dylancorbus/go-database/internal/app/models"
	"github.com/dylancorbus/go-database/internal/pkg/constants"
	constants2 "github.com/dylancorbus/go-database/internal/pkg/index/constants"
	"os"
	"strings"
)

func Read(id string) (*models.Employee, []models.Employee, error) {
	if(id == "*") {
		emps, err := readAll()
		if err != nil {
			return nil, nil, err
		}
		return nil, emps, nil
	}
	emp, err := read(id)
	if err != nil {
		return nil, nil, err
	}
	return emp, nil, nil

}

func read(id string) (*models.Employee, error) {
	file, _ := os.Open(constants.Transaction)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var str string
	for i := 0; scanner.Scan(); i++ {
		x := scanner.Text()
		if i != constants2.Indexes[id]-1 {
			continue
		}
		if arr := strings.Split(x, ";"); arr[0] == id {
			str = x
			break
		}
	}
	file.Close()
	emp, err := models.New(str)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return emp, nil
}

func readAll() ([]models.Employee, error) {
	file, _ := os.Open(constants.Transaction)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var emps []models.Employee
	for scanner.Scan() {
		x := scanner.Text()
		emp, err := models.New(x)
		if err != nil {
			return nil, err
		}
		emps = append(emps, *emp)
	}
	file.Close()
	return emps, nil
}
