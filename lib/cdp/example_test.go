package cdp_test

import (
	"context"
	"fmt"

	"github.com/go-rod/rod/lib/cdp"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/ysmood/kit"
)

func ExampleClient() {
	// launch a browser
	url := launcher.New().Headless(false).Launch()

	// create a controller
	client := cdp.New(url).Connect()

	// Such as call this endpoint on the api doc:
	// https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
	// This will create a new tab and navigate to the test.com
	res, err := client.Call(context.Background(), "", "Target.createTarget", map[string]string{
		"url": "https://google.com",
	})
	kit.E(err)

	fmt.Println(kit.JSON(res).Get("targetId").String())

	kit.Pause()

	// Skip
	// Output: id
}
