package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dylancorbus/go-database/internal/pkg/constants"
)

func start() {
	initIndex()
	fmt.Println(indexes)
}

var indexes = make(map[string]int)

type Employee struct {
	name     string
	salary   int
	vacation bool
}

func initIndex() {
	file, _ := os.Open(constants.Index)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		empArr := strings.Split(scanner.Text(), ";")
		x, err := strconv.Atoi(empArr[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		indexes[empArr[0]] = x
	}
	file.Close()
}

func newEmp(str string) (*Employee, error) {
	if str == "" {
		return nil, errors.New("empty string")
	}
	strArr := strings.Split(str, ";")
	b, err := strconv.ParseBool(strArr[2])
	i, err := strconv.ParseInt(strArr[1], 0, 36)
	if err != nil {
		return nil, err
	}
	emp := Employee{strArr[0], int(i), b}
	return &emp, nil
}

func read(id string) (*Employee, error) {
	file, _ := os.Open(constants.Transaction)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var str string
	for i := 0; scanner.Scan(); i++ {
		x := scanner.Text()
		if i != indexes[id]-1 {
			continue
		}
		if arr := strings.Split(x, ";"); arr[0] == id {
			str = x
			break
		}
	}
	file.Close()
	emp, err := newEmp(str)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return emp, nil
}

func update(id string, field string, value string) {
	emp, err := read(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch field {
	case "salary":
		fmt.Println("Changed ", emp.salary, field, "to", value)
		emp.salary, _ = strconv.Atoi(value)
	case "name":
		fmt.Println("Changed ", emp.name, field, "to", value)
		updateIndex(emp.name, value)
		emp.name = value
	case "vacation":
		fmt.Println("Changed ", emp.name, field, "to", value)
		emp.vacation, _ = strconv.ParseBool(value)
	}
	s := []string{emp.name, strconv.Itoa(emp.salary), strconv.FormatBool(emp.vacation)}
	v := strings.Join(s, ";")
	updateLine(constants.Transaction, indexes[emp.name]-1, v)
}

func deleteItem(id string) {
	arr := removeLine(constants.Index, indexes[id])
	removeLine(constants.Transaction, indexes[id])
	delete(indexes, id)
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
		indexes[textArr[0]] = line
		updateLine(constants.Index, i, strings.Join(textArr, ";"))
	}
}

func removeLine(path string, lineNumber int) []string {
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
func updateLine(path string, lineNumber int, update string) {
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

func updateIndex(old string, new string) {
	file, _ := os.Open(constants.Index)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for i := 0; scanner.Scan(); i++ {
		if i != indexes[old]-1 {
			continue
		}
		x := scanner.Text()
		if arr := strings.Split(x, ";"); arr[0] == old {
			arr[0] = new
			s := strings.Join(arr, ";")
			updateLine(constants.Index, i, s)
			break
		}
	}
	x := indexes[old]
	delete(indexes, old)
	indexes[new] = x
	file.Close()
}

func makeIndex(id string, v int) {
	idx, err := os.OpenFile(constants.Index, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//create entry in index
	indexes[id] = v
	str := []string{id, strconv.Itoa(v)}
	idx.WriteString(strings.Join(str, ";") + "\n")
	idx.Close()
}

func create(name string, salary int) error {
	var emp = Employee{name: name, salary: salary}
	if indexes[name] != 0 {
		return errors.New("record already exists")
	}
	// If the file doesn't exist, create it, or append to the file
	txn, err := os.OpenFile("/Users/dylancorbus/Desktop/test/go-database/internal/app/db/transaction.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	s := []string{emp.name, strconv.Itoa(emp.salary), strconv.FormatBool(emp.vacation)}
	v := strings.Join(s, ";")
	//write to file
	txn.WriteString(v + "\n")
	makeIndex(name, len(indexes)+1)
	txn.Close()
	//create transaction log
	fmt.Println("Writing to file ", emp)
	return nil
}

func main() {
	start()
	fmt.Println(create("John II", 600))
	fmt.Println(create("Jim", 200))
	fmt.Println(create("Karol", 150))
	fmt.Println(create("Joe", 50))
	fmt.Println(create("Jermy", 60))
	fmt.Println(create("Justin", 75))
	fmt.Println(create("Jay", 80))
	fmt.Println(create("Jess", 120))
	fmt.Println(create("Jen", 110))
	fmt.Println(create("Jorge", 130))
	update("Karol", "vacation", "true")
	read("Jess")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, " ")
		fmt.Println("EXECUTING QUERY:", text)
		switch args[0] {
		case "SELECT":
			emp, _ := read(args[1])
			fmt.Println(emp)
		case "UPDATE":
			update(args[1], args[2], args[3])
		case "INSERT":
			i, _ := strconv.Atoi(args[5])
			create(args[3], i)
		case "DELETE":
			deleteItem(args[1])
		case "exit()":
			os.Exit(1)
		}
	}

}

func parse(query string) {

}
