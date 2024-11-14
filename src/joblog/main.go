package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	// "io"
	"log"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/dao"
) 

func getRoot(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("home"))
}

func getDash(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("Dashboard"))
}

func main(){
	// data.InitDb()
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/dash", getDash)
	var test models.Job 
	test.Title = "test title"
	test.Description = "test Description"
	dao.CreateJob(&test)

	client := data.InitS3()
	
	file, fileErr := os.Open("./assets/BurlakaAssignment9.pdf")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	data.UploadS3(client, file)

	

	println("server running")
	serverErr := http.ListenAndServe("127.0.0.1:3333", nil)

	if errors.Is(serverErr, http.ErrServerClosed){
		fmt.Printf("server closed\n")
	}else if serverErr != nil {
		fmt.Print("error starting server: %s\n", serverErr)
		os.Exit(1)
	}


}

