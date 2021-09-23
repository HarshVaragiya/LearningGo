package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	task "github.com/HarshVaragiya/LearningGo/Protobuf/protobuf-datastore/taskproto"
	"google.golang.org/protobuf/proto"
)

var dataStore = "task-datastore.pb"

func main() {

	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatalf("no arguments supplied. please supply add/list. exiting.")
	}

	switch action := flag.Arg(0); action {

	case "add":
		task := strings.Join(flag.Args()[1:], " ")
		if task == "" {
			log.Fatal("no task data supplied. exiting")
		}
		log.Printf("adding task '%s' ", task)
		err := addTask(task)
		if err != nil {
			log.Fatalf("error adding task. error = %v", err)
		}
	case "list":
		log.Print("stored tasks are listed below.")
		err := listTasks()
		if err != nil {
			log.Fatalf("error listing task. error = %v", err)
		}
	default:
		log.Fatalf("unrecognised action '%s'. exiting", action)
	}
}

func addTask(todo string) error {
	taskObject := &task.Task{Name: todo, Done: false, Added: time.Now().Format(time.RFC850)}
	file, err := os.OpenFile(dataStore, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening datastore file %v. error = %v", dataStore, err)
		return err
	}
	existingDataStoreBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error reading datastore file %v. err = %v", dataStore, err)
		return err
	} else {
		log.Printf("read %v bytes from datastore", len(existingDataStoreBytes))
	}
	var taskList task.TaskList
	if err := proto.Unmarshal(existingDataStoreBytes, &taskList); err != nil {
		log.Println("error decoding datastore information. existing data might be corrupted.")
		log.Printf("error = %v", err)
		return err
	}
	taskList.Tasks = append(taskList.Tasks, taskObject)
	outputBytes, err := proto.Marshal(&taskList)
	if err != nil {
		log.Printf("error marshalling the task list object. error = %v", err)
		return err
	}
	if n, err := file.WriteAt(outputBytes, 0x00); err != nil {
		log.Print("error writing bytes in the file. err = %v", err)
		return err
	} else {
		log.Printf("wrote %v bytes to datastore", n)
	}
	if err = file.Close(); err != nil {
		log.Printf("error closing datastore. error = %v", err)
		log.Printf("exiting to prevent data corruption.")
		return err
	}
	return nil
}

func listTasks() error {
	dataStoreBytes, err := ioutil.ReadFile(dataStore)
	if err != nil {
		log.Printf("error opening data storage file %s : %v ", dataStore, err)
		return err
	} else {
		log.Printf("read %v bytes from datastore", len(dataStoreBytes))
	}
	var taskList task.TaskList
	if err := proto.Unmarshal(dataStoreBytes, &taskList); err != nil {
		log.Println("error decoding datastore information. existing data might be corrupted.")
		log.Printf("error = %v", err)
		return err
	}
	log.Println("contents of datastore are : ")
	displayTasks(taskList)
	return nil
}

func displayTasks(tasks task.TaskList) {
	for index, task := range tasks.Tasks {
		chr := "[X]"
		if !task.Done {
			chr = "[ ]"
		}
		fmt.Printf("[Task %v] %v - %s -  %s \n", index, chr, task.Added, task.Name)
	}
}
