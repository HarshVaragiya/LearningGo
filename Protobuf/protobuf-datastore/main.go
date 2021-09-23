package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	task "github.com/HarshVaragiya/LearningGo/Protobuf/protobuf-datastore/taskproto"
	"google.golang.org/protobuf/proto"
)

var dataStore = "protobuf-datastore/task-datastore.pb"

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
	taskBytes, err := proto.Marshal(taskObject)
	if err != nil {
		log.Printf("error marshalling the task object. error = %v", err)
		return err
	}
	file, err := os.OpenFile(dataStore, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening data storage file %s : %v ", dataStore, err)
		return err
	}
	return nil
}

func listTasks() error {
	return nil
}
