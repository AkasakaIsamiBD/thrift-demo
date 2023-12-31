// Code generated by Thrift Compiler (0.18.1). DO NOT EDIT.

package tutorial

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gen-go/shared"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"regexp"
	"strings"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

var _ = shared.GoUnusedProtection__

const INT32CONSTANT = 9853

var MAPCONSTANT map[string]string

func init() {
	MAPCONSTANT = map[string]string{
		"goodnight": "moon",
		"hello":     "world",
	}

}
