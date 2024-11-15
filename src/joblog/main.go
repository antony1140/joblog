package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
) 

func getRoot(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("root"))
}

func getDash(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("Dashboard"))
}



func main(){
	// data.InitDb()
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/dash", getDash)
	http.HandleFunc("/home", func(resp http.ResponseWriter, req *http.Request){
		resp.Write([]byte("home"))
	})
	var test models.Job 
	test.Title = "test title"
	test.Description = "test Description"
	dao.CreateJob(&test)
	client := data.InitS3()
	data.DownloadS3(client, "testDownload")


	

	println("server running")
	serverErr := http.ListenAndServe("127.0.0.1:3333", nil)

	if errors.Is(serverErr, http.ErrServerClosed){
		fmt.Printf("server closed\n")
	}else if serverErr != nil {
		fmt.Print("error starting server: %s\n", serverErr)
		os.Exit(1)
	}


}

