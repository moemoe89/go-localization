//
//  language_test.go
//  go-localization
//
//  Copyright © 2019. All rights reserved.
//

package language

import (
	"testing"
)

func TestBindPath(t *testing.T) {
	cfg := New()

	path := "./example/language.json"
	cfg.BindPath(path)
	if cfg.path[0] != path {
		t.Errorf("Should return %s, got %s", path, cfg.path)
	}
}

func TestBindMainLocale(t *testing.T) {
	cfg := New()

	mainLocale := "en"
	cfg.BindMainLocale(mainLocale)
	if cfg.mainLocale != mainLocale {
		t.Errorf("Should return %s, got %s", mainLocale, cfg.mainLocale)
	}
}

func TestInit(t *testing.T) {
	cfg := New()
	cfg.BindPath("./example/language.json")
	cfg.BindMainLocale("en")

	_, err := cfg.Init()
	if err != nil {
		t.Errorf("Should return %v, got %s", nil, err.Error())
	}
}

func TestInitFail(t *testing.T) {
	cfg := New()
	cfg.BindPath("./lang.json")
	cfg.BindMainLocale("en")

	_, err := cfg.Init()
	if err == nil {
		t.Errorf("Should return error, got %v", nil)
	}
}

func TestInitUnmarshalFail(t *testing.T) {
	cfg := New()
	cfg.BindPath("./README.md")
	cfg.BindMainLocale("en")

	_, err := cfg.Init()
	if err == nil {
		t.Errorf("Should return error, got %v", nil)
	}
}

func TestLookupLangCodeNotFound(t *testing.T) {
	cfg := New()
	cfg.BindPath("./example/language.json")
	cfg.BindMainLocale("en")

	lang, _ := cfg.Init()
	expectedResult := "Good night"
	result := lang.Lookup("fr", "Good night")
	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupMessageNotFound(t *testing.T) {
	cfg := New()
	cfg.BindPath("./example/language.json")
	cfg.BindMainLocale("en")

	lang, _ := cfg.Init()
	expectedResult := "Good night"
	result := lang.Lookup("en", "Good night")
	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}

func TestLookupMessageFound(t *testing.T) {
	cfg := New()
	cfg.BindPath("./example/language.json")
	cfg.BindMainLocale("en")

	lang, _ := cfg.Init()
	expectedResult := "Selamat pagi"
	result := lang.Lookup("id", "Good morning")
	if result != expectedResult {
		t.Errorf("Should return %s, got %s", expectedResult, result)
	}
}
