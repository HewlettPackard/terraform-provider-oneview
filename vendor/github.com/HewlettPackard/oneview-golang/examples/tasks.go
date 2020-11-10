package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		ClientOV *ov.OVClient
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")
	// Get all tasks present
	sort := "name:desc"
	count := "5"
	fmt.Println("\nGetting all tasks present: \n")
	task_list, err := ovc.GetTasks("", sort, count, "")
	if err != nil {
		fmt.Println("Error getting the storage volumes ", err)
	}
	for i := 0; i < len(task_list.Members); i++ {
		fmt.Println(task_list.Members[i].Name, task_list.Members[i].URI)
	}

	id := strings.Replace(string(task_list.Members[0].URI), "/rest/tasks/", "", 1)

	fmt.Println("Getting Details of the requested Task")
	task, err := ovc.GetTasksById("", "", "", "", id)
	if err != nil {
		fmt.Println("Error getting the task details ", err)
	}
	fmt.Println(task)

}
