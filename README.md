# gospritz - Spritz program in Golang

gospritz is a Spritz program in Golang. To know Spritz, see the following links:

* [Official site](http://www.spritzinc.com/) (English)
* [単語が目に飛び込んできてすごい速度で文章を読めるようになる「Spritz」 - GIGAZINE](http://gigazine.net/news/20140228-spritz/) (Japanese)

## Installation

```sh
$ go get github.com/yosssi/gospritz
```

## Usage

```sh
$ gospritz -w WPM(WORDS PER MINUTE) FILEPATH 
```

## Example

```sh
$ cd $GOPATH/src/github.com/yosssi/gospritz
$ gospritz -w 250 ./sample.txt
```
