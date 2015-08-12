/*
* main
 */
package main

import (
	"controller"
	"fmt"
	"utils"
)

func main() {
	fmt.Println("start run main ...")
	utils.InitializeWidow()
	controller.CrawlController()
}
