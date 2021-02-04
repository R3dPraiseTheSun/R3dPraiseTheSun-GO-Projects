package main

import (
	"fmt"
	"path/filepath"
    "context"
    "io/ioutil"
    "log"
    "math"
    "os"
    "strings"

    "github.com/chromedp/cdproto/emulation"
    "github.com/chromedp/cdproto/page"
    "github.com/chromedp/chromedp"
)

func main() {
    // create context
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()
    // used to get absolute path of a file
   	abs,err := filepath.Abs("./demos/")
   	if err != nil {
   	   log.Fatal(err)
   	}

   	//name of files inside html's folder
   	files, err := ioutil.ReadDir("./demos/")
	if err != nil {
	    log.Fatal(err)
	}

	//creates a folder named ScreenShots if the folder does not exist
	if _, err := os.Stat("./ScreenShots"); os.IsNotExist(err) {
    	errdir := os.Mkdir("ScreenShots", 0755)
    	if errdir != nil {
    		log.Fatal(errdir)
    	}
	}

	for _, f := range files {
	    if strings.Contains(f.Name(),".html"){
	    	fmt.Println(f.Name())
    		//if you want to use html from your local filesystem use file:/// + absolute path to your html file
	   		url :=  "file:///" + abs + "/" + f.Name()	
		   	// capture screenshot of an element
		   	var buf []byte
		  	// capture entire browser viewport, returning png with quality=90
	   		if err := chromedp.Run(ctx, fullScreenshot(url, 90, &buf)); err != nil {
		    log.Fatal(err)
		   	}
		   	screenshotName := "ScreenShots/" + strings.Replace(f.Name(),".html", ".png", 1)	
	   		if err := ioutil.WriteFile(screenshotName, buf, 0644); err != nil {
	   		    log.Fatal(err)
    		}
    	}
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
    return chromedp.Tasks{
        chromedp.Navigate(urlstr),
        chromedp.ActionFunc(func(ctx context.Context) error {
            // get layout metrics
            _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
            if err != nil {
                return err
            }

            width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

            // force viewport emulation
            err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
                WithScreenOrientation(&emulation.ScreenOrientation{
                    Type:  emulation.OrientationTypePortraitPrimary,
                    Angle: 0,
                }).
                Do(ctx)
            if err != nil {
                return err
            }

            // capture screenshot
            *res, err = page.CaptureScreenshot().
                WithQuality(quality).
                WithClip(&page.Viewport{
                    X:      contentSize.X,
                    Y:      contentSize.Y,
                    Width:  contentSize.Width,
                    Height: contentSize.Height,
                    Scale:  1,
                }).Do(ctx)
            if err != nil {
                return err
            }
            return err
        }),
    }
}