package storage

import (
	"log"
	"os"
	"testing"
)

var s *Storage

func TestMain(m *testing.M) {
	pwd := os.Getenv("dbpass")
	if pwd == "" {
		m.Run()
	}
	var err error
	s, err = New("postgres://postgres:" + pwd + "@192.168.1.62:/tasks")
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}
func TestStorage_Tasks(t *testing.T) {
	taskID := 0
	authorID := -1
	data, err := s.Tasks(taskID, authorID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
	taskID = 0
	authorID = 0
	data, err = s.Tasks(taskID, authorID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
	taskID = 1
	authorID = 1
	data, err = s.Tasks(taskID, authorID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestStorage_NewTask(t *testing.T) {
	task := Task{Title: "New Test Task", Content: "New test task content"}
	id, err := s.NewTask(task)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создана задача с id: ", id)
	taskID := 0
	authorID := -1
	data, err := s.Tasks(taskID, authorID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestStorage_TasksLabels(t *testing.T) {
	labelName := "git"
	data, err := s.TasksLabels(labelName)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
	labelName = ""
	data, err = s.TasksLabels(labelName)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

func TestStorage_DeleteTask(t *testing.T) {
	taskID := 1
	err := s.DeleteTask(taskID)
	if err != nil {
		t.Fatal(err)
	}
	taskID = 15
	err = s.DeleteTask(taskID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStorage_UpdateTask(t *testing.T) {
	taskID := 2
	title := "Update task title"
	content := "Update task content"
	var closed int64 = 0
	err := s.UpdateTask(taskID, title, content, closed)
	if err != nil {
		t.Fatal(err)
	}
	taskID = 1
	title = ""
	content = ""
	closed = 0
	err = s.UpdateTask(taskID, title, content, closed)
	if err != nil {
		t.Fatal(err)
	}
	taskID = 0
	authorID := -1
	data, err := s.Tasks(taskID, authorID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
