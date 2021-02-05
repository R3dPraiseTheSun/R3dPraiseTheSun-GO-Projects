package createfolder

import (
    "log"
    "os"
)

func CreateFolder(){
    //creates a folder named ScreenShots if the folder does not exist
    if _, err := os.Stat("./ScreenShots"); os.IsNotExist(err) {
        errdir := os.Mkdir("ScreenShots", 0755)
        if errdir != nil {
            log.Fatal(errdir)
        }
    }
}