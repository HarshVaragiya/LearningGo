package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	task "github.com/HarshVaragiya/LearningGo/Protobuf/gRPC/taskproto"
	"google.golang.org/grpc"
)

func main() {

	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatalf("no arguments supplied. please supply add/list. exiting.")
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect to server. error = %v", err)
	}

	client := task.NewTasksClient(conn)

	switch action := flag.Arg(0); action {

	case "add":
		task := strings.Join(flag.Args()[1:], " ")
		if task == "" {
			log.Fatal("no task data supplied. exiting")
		}
		log.Printf("adding task '%s' ", task)
		err := addTask(task, context.Background(), client)
		if err != nil {
			log.Fatalf("error adding task. error = %v", err)
		}
	case "list":
		err := listTasks(context.Background(), client)
		if err != nil {
			log.Fatalf("error listing task. error = %v", err)
		}
	default:
		log.Fatalf("unrecognised action '%s'. exiting", action)
	}
}

func addTask(text string, ctx context.Context, client task.TasksClient) (err error) {
	newTask := task.Task{Name: text, Done: false}
	if _, err = client.Add(ctx, &newTask); err != nil {
		log.Printf("could not add task to server. error = %v", err)
	} else {
		log.Println("adding task to server successful!")
	}
	return
}

func listTasks(ctx context.Context, client task.TasksClient) error {
	log.Println("fetching task list from server.")
	list, err := client.List(ctx, &task.Void{})
	if err != nil {
		log.Fatalf("error fetching task list from server. error = %v", err)
	} else {
		log.Println("Tasks are as follows: ")
		displayTasks(*list)
	}
	return err
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
