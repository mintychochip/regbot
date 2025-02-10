package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
	"log"
	"os"
	"time"
)

type PlaywrightWrapper struct {
	playwright *playwright.Playwright
	err        error
}

type BrowserWrapper struct {
	browser *playwright.Browser
	err     error
}

type PageWrapper struct {
	page *playwright.Page
	err  error
}

func createPlaywrightWrapper() *PlaywrightWrapper {
	err := playwright.Install()
	if err != nil {
		return &PlaywrightWrapper{err: err}
	}
	pw, err := playwright.Run()
	return &PlaywrightWrapper{playwright: pw, err: err}
}

func (pw *PlaywrightWrapper) launchBrowser(bt playwright.BrowserType, options playwright.BrowserTypeLaunchOptions) *BrowserWrapper {
	browser, err := bt.Launch(options)
	if err != nil {
		return &BrowserWrapper{err: err}
	}
	return &BrowserWrapper{browser: &browser, err: nil}
}

func (bw *BrowserWrapper) newPage(uri string, options playwright.PageGotoOptions) *PageWrapper {
	page, err := (*bw.browser).NewPage()
	if err != nil {
		return &PageWrapper{err: err}
	}
	_, err = page.Goto(uri, options)
	if err != nil {
		return &PageWrapper{err: err}
	}
	return &PageWrapper{page: &page, err: nil}
}

func (pw *PageWrapper) locator(selector string, action func(playwright.Locator) error) {
	locator := (*pw.page).Locator(selector)
	if locator == nil {
		return
	}
	pw.err = action(locator)
}

var stateVisible = playwright.LocatorWaitForOptions{
	State: playwright.WaitForSelectorStateVisible,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	wrapper := createPlaywrightWrapper()
	if wrapper.err != nil {
		log.Fatal(wrapper.err)
	}

	browserWrapper := wrapper.launchBrowser(wrapper.playwright.Chromium, playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})

	if browserWrapper.err != nil {
		log.Fatal(browserWrapper.err)
	}
	pageWrapper := browserWrapper.newPage("https://cmsweb.cms.csub.edu/psp/CBAKPRD/?cmd=login", playwright.PageGotoOptions{})
	authCSUB(pageWrapper)
	if pageWrapper.err != nil {
		log.Fatal(pageWrapper.err)
	}

	time.Sleep(100 * time.Second)
}

func authCSUB(pageWrapper *PageWrapper) {
	pageWrapper.locator("[name='loginfmt']", func(locator playwright.Locator) error {
		err := locator.WaitFor(stateVisible)
		if err != nil {
			return err
		}
		return locator.Fill(os.Getenv("EMAIL"))
	})
	pageWrapper.locator("[value='Next']", func(locator playwright.Locator) error {
		err := locator.WaitFor(stateVisible)
		if err != nil {
			return err
		}
		return locator.Click()
	})
	pageWrapper.locator("[name='passwd']", func(locator playwright.Locator) error {
		err := locator.WaitFor(stateVisible)
		if err != nil {
			return err
		}
		return locator.Fill(os.Getenv("PASSWD"))
	})
	pageWrapper.locator("[value='Sign in']", func(locator playwright.Locator) error {
		err := locator.WaitFor(stateVisible)
		if err != nil {
			return err
		}
		return locator.Click()
	})
	time.Sleep(4 * time.Second)
	robotgo.TypeStr(os.Getenv("PIN"))
	pageWrapper.locator("#trust-browser-button", func(locator playwright.Locator) error {
		err := locator.WaitFor(stateVisible)
		if err != nil {
			return err
		}
		return locator.Click()
	})
	pageWrapper.locator("[value='Yes']", func(locator playwright.Locator) error {
		err := locator.WaitFor(stateVisible)
		if err != nil {
			return err
		}
		return locator.Click()
	})

	pageWrapper.locator("//a[text()='Enroll']", func(locator playwright.Locator) error {
		err := locator.WaitFor(stateVisible)
		if err != nil {
			return err
		}
		return locator.Click()
	})
	array := []string{"32540", "32544", "32548", "32550"}
	p := (*pageWrapper.page)

	for _, v := range array {
		fmt.Println(v)
		_, _ = p.WaitForSelector("[name='DERIVED_REGFRM1_CLASS_NBR']", playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateVisible,
			Timeout: playwright.Float(500),
		})
		locator := p.Locator("[name='DERIVED_REGFRM1_CLASS_NBR']")
		_ = locator.Fill(v)

		_, _ = p.WaitForSelector("[value='Enter']", playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateVisible,
			Timeout: playwright.Float(500),
		})
		locator = p.Locator("[value='Enter']")
		locator.Click()
		for {
			time.Sleep(300 * time.Millisecond)
			_, err := p.WaitForSelector(`a.gh-footer-item >> span:text("Next")`, playwright.PageWaitForSelectorOptions{
				State:   playwright.WaitForSelectorStateVisible,
				Timeout: playwright.Float(500),
			})
			if err != nil {
				fmt.Println("Next button not found or timeout reached:", err)
				break
			}
			l := p.Locator(`a.gh-footer-item >> span:text("Next")`)
			err = l.Click()
			if err != nil {
				break
			}

		}
	}
}
