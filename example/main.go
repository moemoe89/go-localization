//
//  main.go
//  go-localization
//
//  Copyright © 2019. All rights reserved.
//

package main

import (
	"fmt"

	"github.com/moemoe89/go-localization"
)

var lang *language.Config

func init() {
	// initiate the go-localization & bind some config
	cfg := language.New()
	cfg.BindPath("./language.json")
	cfg.BindMainLocale("en")
	var err error
	lang, err = cfg.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	// print some message in different languages
	fmt.Println(lang.Lookup("en", "Good morning"))
	fmt.Println(lang.Lookup("id", "Good morning"))
	fmt.Println(lang.Lookup("jp", "Good morning"))
	// if language not found, will return default language
	fmt.Println(lang.Lookup("fr", "Good morning"))
	// print unlisted message
	fmt.Println(lang.Lookup("en", "Good night"))
}
