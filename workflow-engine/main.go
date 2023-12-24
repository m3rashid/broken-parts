package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Message struct {
	Type      string
	Data      interface{}
	Timestamp time.Time
}

func (m *Message) GetWorkflow() (Workflow, error) {
	for _, wf := range AllWorkflows {
		if wf.Trigger == m.Type {
			return wf, nil
		}
	}

	return Workflow{}, nil
}

type Workflow struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Trigger     string `json:"trigger"`
	Job         struct {
		Nodes []struct {
			ID   string      `json:"id"`
			Type string      `json:"type"`
			Meta interface{} `json:"meta"`
		} `json:"nodes"`
		Edges []struct {
			From []string `json:"from"`
			To   []string `json:"to"`
			Type string   `json:"type"`
		} `json:"edges"`
	} `json:"job"`
}

var AllWorkflows []Workflow

func (w *Workflow) Execute() ActivityLog {
	activity := ActivityLog{}
	activity.StartTime = time.Now()
	activity.Status = "pending"

	// construct the tree on the basis of json configuration
	for _, edge := range w.Job.Edges {

	}

	for _, node := range w.Job.Nodes {

	}

	return activity
}

func (w Workflow) Add() {
	AllWorkflows = append(AllWorkflows, w)
}

type ActivityStatus string

var (
	ActivityStatusCompleted ActivityStatus = "completed"
	ActivityStatusPending   ActivityStatus = "pending"
	ActivityStatusError     ActivityStatus = "error"
)

type ActivityLog struct {
	StartTime time.Time      `json:"startTime"`
	EndTime   time.Time      `json:"endTime"`
	Status    ActivityStatus `json:"status"`
	Result    interface{}    `json:"result"`
}

func (a *ActivityLog) ToJson() (string, error) {
	jsonStr, err := json.Marshal(*a)
	return string(jsonStr), err
}

var ActivityLogs []ActivityLog

func main() {
	file, err := os.Open("sample.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var workflow Workflow
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&workflow)
	if err != nil {
		log.Fatal(err)
	}

	activity := workflow.Execute()
	json, err := activity.ToJson()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("The corresponding data obtained: ", json)
}
