package main

import (
	"bufio"
	"fmt"
	"github.com/dylancorbus/go-database/internal/pkg/index/constants"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dylancorbus/go-database/internal/pkg/db/operations"
	fileConst "github.com/dylancorbus/go-database/internal/pkg/constants"
)

func start() {
	initIndex()
	fmt.Println(constants.Indexes)
}

func initIndex() {
	file, _ := os.Open(fileConst.Index)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		empArr := strings.Split(scanner.Text(), ";")
		x, err := strconv.Atoi(empArr[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		constants.Indexes[empArr[0]] = x
	}
	file.Close()
}

func main() {
	start()
	//fmt.Println(operations.Create("John II", 600))
	//fmt.Println(operations.Create("Jim", 200))
	//fmt.Println(operations.Create("Karol", 150))
	//fmt.Println(operations.Create("Joe", 50))
	//fmt.Println(operations.Create("Jermy", 60))
	//fmt.Println(operations.Create("Justin", 75))
	//fmt.Println(operations.Create("Jay", 80))
	//fmt.Println(operations.Create("Jess", 120))
	//fmt.Println(operations.Create("Jen", 110))
	//fmt.Println(operations.Create("Jorge", 130))
	//operations.Update("Karol", "vacation", "true")
	//operations.Read("Jess")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, " ")
		fmt.Println("EXECUTING QUERY:", text)
		switch args[0] {
		//SELECT someName
		case "SELECT":
			emp, emps, err := operations.Read(args[1])
			if err != nil {
				log.Fatal(err)
			} else if emp != nil {
				fmt.Println(emp)
			} else {
				fmt.Println(emps)
			}
		//UPDATE someName field value
		case "UPDATE":
			operations.Update(args[1], args[2], args[3])
		//INSERT Employee name: someName salary: value
		case "INSERT":
			i, _ := strconv.Atoi(args[5])
			operations.Create(args[3], i)
		//DELETE someName
		case "DELETE":
			operations.DeleteItem(args[1])
		case "exit()":
			os.Exit(1)
		}
	}

}
