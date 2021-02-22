package chromehandle

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"path/filepath"
	"strings"
	"sync"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func ScreenShotSave(filePath string, fileName string, wg *sync.WaitGroup) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	fmt.Println("Opening chrome to page: ", getAbsolutePath(filePath), "...")
	if err := chromedp.Run(ctx, fullScreenshot("file:///"+getAbsolutePath(filePath), 90, &buf)); err != nil {
		log.Fatal(err)
	}
	screenshotName := "./" + strings.Replace(fileName, ".html", ".png", 1)
	if err := ioutil.WriteFile(screenshotName, buf, 0644); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func fullScreenshot(filePath string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(filePath),
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

func getAbsolutePath(filePath string) string {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return abs
}
