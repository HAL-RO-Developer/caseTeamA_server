package main

import (
	"fmt"

	"github.com/HAL-RO-Developer/caseTeamA_server/service"
)

func main() {
	message, find := service.GetMessageInfoFromSame("sample", 1, 3, 10)
	if !find {
		fmt.Println("not found")
	}
	fmt.Println(len(message))

}
