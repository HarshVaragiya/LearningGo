package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	task "github.com/HarshVaragiya/LearningGo/Protobuf/gRPC/taskproto"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var dataStore = "task-datastore.pb"

type taskServer struct {
}

func (s taskServer) List(ctx context.Context, void *task.Void) (*task.TaskList, error) {
	log.Println("client requested task list.")
	return listTasks()
}

func (s taskServer) Add(ctx context.Context, newTask *task.Task) (*task.Void, error) {
	log.Printf("client requested adding new task: %s", newTask.Name)
	newTask.Added = time.Now().Format(time.RFC850)
	err := addTask(newTask)
	return &task.Void{}, err
}

func main() {
	log.Println("Starting TaskManager gRPC Server")
	srv := grpc.NewServer()
	var taskserver taskServer
	log.Println("Registering task server")
	task.RegisterTasksServer(srv, taskserver)
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not start listener. error = %v", err)
	}
	log.Println("attached to listener port. serving gRPC service")
	err = srv.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

func addTask(taskObject *task.Task) error {
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

func listTasks() (*task.TaskList, error) {
	var taskList task.TaskList
	dataStoreBytes, err := ioutil.ReadFile(dataStore)
	if err != nil {
		log.Printf("error opening data storage file %s : %v ", dataStore, err)
		return &taskList, err
	} else {
		log.Printf("read %v bytes from datastore", len(dataStoreBytes))
	}
	if err = proto.Unmarshal(dataStoreBytes, &taskList); err != nil {
		log.Println("error decoding datastore information. existing data might be corrupted.")
		log.Printf("error = %v", err)
		return &taskList, err
	} else {
		log.Println("sucessfully retrieved task list from datastore")
	}
	return &taskList, nil
}
