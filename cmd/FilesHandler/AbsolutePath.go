package fileshandler

import(

	"path/filepath"
    "log"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func GetPath(htmlPath ...string) string {
	var path string
	if len(htmlPath) == 0{
    	abs,err := filepath.Abs("./demos/")
    	if err != nil {
       		log.Fatal(err)
       	}
       	path = abs
    } else {
        for _,customPath := range htmlPath {
    	   abs,err := filepath.Abs(customPath)
    	   if err != nil {
          log.Fatal(err)
   		   }
   		   path = abs
        }
      }
    return path
}

func GetHTMLFiles(path string) []os.FileInfo{
	files, err := ioutil.ReadDir(path)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println("I found ", len(files), " files inside", path)
	return files
}

func FileNames(files []os.FileInfo, path string) []string{
	var filePaths []string
	for _, f := range files {
	   if strings.Contains(f.Name(),".html"){
    		//if you want to use html from your local filesystem use file:/// + absolute path to your html file
	   		thisPath :=  "file:///" + path + "/" + f.Name()	
	   		filePaths = append(filePaths,thisPath)
    	}
	}
	return filePaths
}