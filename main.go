package main

import (
	"fmt"

	"gopkg.in/headzoo/surf.v1"
)

// Source http://surf.readthedocs.io/#quick-start

func main() {
	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	err := bow.Open("http://indeed.com")
	if err != nil {
		panic(err)
	}
	// Outputs: "reddit: the front page of the internet"
	fmt.Println(bow.Title())

	// Click the link for the newest submissions.

	err = bow.Click(".signin")
	if err != nil {
		fmt.Println("Click failed", err)
	}

	// fmt.Println(el)

	// bow.Find(".signin").Each(func(_ int, s *goquery.Selection) {
	// 	fmt.Println(s.Text(), *s)

	// })

	// bow.Find("ul.icl-DesktopGlobalHeader-items icl-DesktopGlobalHeader-items--right").Each(func(_ int, s *goquery.Selection) {
	// 	fmt.Println(s.Text(), *s)

	// })

	// Outputs: "newest submissions: reddit.com"
	//fmt.Println(bow.Title())

	// Log in to the site.

	// fm, _ := bow.Form("form.loginform")
	// fm.Input("Email Address", "doronesk@gmail.com")
	// fm.Input("Password", "%08EaowRR2")
	// if fm.Submit() != nil {
	// 	panic(err)
	// }

	//bow.Click("nav.settings")

	// Go back to the "newest submissions" page, bookmark it, and
	// print the title of every link on the page.
	//bow.Back()
	//bow.Bookmark("reddit-new")
	// bow.Find("app-boot-bg").Each(func(_ int, s *goquery.Selection) {
	// 	fmt.Println(s.Text())

	// })
}
