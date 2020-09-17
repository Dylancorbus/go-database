package operations

import (
	"github.com/dylancorbus/go-database/internal/pkg/constants"
	constants2 "github.com/dylancorbus/go-database/internal/pkg/index/constants"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateIndex(id string, v int) {
	idx, err := os.OpenFile(constants.Index, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//create entry in index
	constants2.Indexes[id] = v
	str := []string{id, strconv.Itoa(v)}
	idx.WriteString(strings.Join(str, ";") + "\n")
	idx.Close()
}
