goic
=======
[![GoDoc](https://godoc.org/github.com/hapoon/goic?status.png)](https://godoc.org/github.com/hapoon/goic)
[![Build Status](https://travis-ci.org/hapoon/goic.svg?branch=master)](https://travis-ci.org/hapoon/goic)
[![Coverage Status](https://coveralls.io/repos/github/hapoon/goic/badge.svg?branch=master)](https://coveralls.io/github/hapoon/goic?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hapoon/goic/master/LICENSE)

`goic` is a library for converting go's interface{} to variable type.

## Installation

Make sure that Go is installed on your computer. Type the following command in your terminal:

`go get gopkg.in/hapoon/goic.v1`

After it the package is ready to use.

Add following line in your `*.go` file:

```go
import "gopkg.in/hapoon/goic.v1"
```

## Usage

### interface{} to uint

```go
func Sample(val interface{}) {
    u,err := goic.ToUint()
    if err != nil {
        // error handling
    }
}
```

## License

[MIT License](LICENSE)
