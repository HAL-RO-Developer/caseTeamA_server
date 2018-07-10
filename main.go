package main

import "github.com/HAL-RO-Developer/caseTeamA_server/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
