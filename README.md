# botverifier

Let's you know if a requester is a known good bot or not

## Setup

You can fetch this project using `go get`:

```
$ go get github.com/abedra/botverifier
```

or you can build the project directly from source:

```
$ go build
```

## Usage

```
$  ./botverifier -address "66.249.90.77" -useragent "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
YES
```
