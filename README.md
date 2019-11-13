[![Build Status](https://travis-ci.org/moemoe89/go-localization.svg?branch=master)](https://travis-ci.org/moemoe89/go-localization)
[![Coverage Status](https://coveralls.io/repos/github/moemoe89/go-localization/badge.svg?branch=master)](https://coveralls.io/github/moemoe89/go-localization?branch=master)
[![GoDoc](https://godoc.org/github.com/moemoe89/go-localization?status.svg)](https://godoc.org/github.com/moemoe89/go-localization)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/go-localization)](https://goreportcard.com/report/github.com/moemoe89/go-localization)

# go-localization

## Description

go-localization is a Go package that provides tool for showing message based on language code.
It has the following rules:

* configuration message written in JSON with specific format.
* if the message is found but the language code is not, will return by the default language code that was set up previously.
* if the message is not found, will return the message.

## Requirements

Go programming language

## JSON Format
```
{
  "en": {
    "Good morning": "Good morning"
  },
  "id": {
    "Good morning": "Selamat pagi"
  },
  "jp": {
    "Good morning": "おはようございます"
  }
}

```

## Getting Started

The go-localization package mainly includes a set of methods for managing the data. You use 
method `language.New()` to init the data, and you can call some method to bind config

```
cfg := language.New()
cfg.BindPath("./language.json")
cfg.BindMainLocale("en")
```
 
If you want to user the message convert function, start with the init function

````
lang, err := cfg.Init()
if err != nil {
	panic(err)
}

// print some message in different languages
fmt.Println(lang.Lookup("en", "Good morning"))
fmt.Println(lang.Lookup("id", "Good morning"))
fmt.Println(lang.Lookup("jp", "Good morning"))
// if language not found, will return default language
fmt.Println(lang.Lookup("fr", "Good morning"))
// print unlisted message
fmt.Println(lang.Lookup("en", "Good night"))
````

### Installation

Run the following command to install the package:

```
go get github.com/moemoe89/go-localization
```

### Example

You can find a example for the implementation at [example](https://github.com/moemoe89/go-localization/blob/master/example/main.go) repository.

### Unit Test

go-localization provide unit test with 100% code coverage.
Run the following command to run the unit test:
```
cd $GOPATH/src/github.com/moemoe89/go-localization
go test -v
```

With code coverage:
```
cd $GOPATH/src/github.com/moemoe89/go-localization
go test -cover
```

With html output:
```
cd $GOPATH/src/github.com/moemoe89/go-localization
go test -coverprofile=c.out && go tool cover -html=c.out
```