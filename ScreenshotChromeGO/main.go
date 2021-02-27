package main

import (
    "os"
    "fmt"
    "sync"
    "./cmd/ChromeHandler"
    "./cmd/FilesHandler"
    "./cmd/CreateFolder"
)

func main() {

	var customPath string
	if args := os.Args[1:]; len(args) > 0 {
		for _,arg := range args{
			customPath = arg
		}
	}

	//create a folder if needed
	createfolder.CreateFolder()

	Program(customPath)
}

func Program(customPath string){
	var path string
	//get default ./demos path or custom path
	fmt.Println(customPath)
	if customPath == "" {
		fmt.Println("Getting path...")
		path = fileshandler.GetPath()
	} else {
		fmt.Println("Getting path...")
		path = fileshandler.GetPath(customPath)
	}
	fmt.Println("Getting HTML names inside ", path, "...")
	var filesName []os.FileInfo = fileshandler.GetHTMLFiles(path)

	var wg sync.WaitGroup
	var filesPath []string = fileshandler.FileNames(filesName, path)
	var index int = 0
	for _,path := range filesPath{
		wg.Add(1)
		go chromehandle.ScreenShotSave(path,filesName[index].Name(), &wg)
		index++
	}
	wg.Wait()
}
