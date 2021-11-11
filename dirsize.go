package dirsize

import (
	"errors"
	"io/fs"
	"math"
	"path/filepath"
)

// ByteUnit ...
type ByteUnit float64

const (
	B  ByteUnit = 1
	KB ByteUnit = B * 1024.00
	MB ByteUnit = KB * 1024.00
	GB ByteUnit = MB * 1024.00
	TB ByteUnit = GB * 1024.00
)

// Calc the size of a folder and returns it in the unit specified by u.
// The calculation result is rounded to two decimal places.
func Calc(root string, u ByteUnit, opts ...OptionFunc) (float64, error) {
	var size int64
	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, opt := range opts {
				switch err := opt(path, info, err); err {
				case nil:
				case SkipFile:
					return nil
				default:
					return err
				}
			}
			size += info.Size()
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return math.Round(float64(size)/float64(u)*100) / 100, nil
}

// SkipFile is used as a return value from OptionFunc to indicate that
// the file named in the call is to be skipped. It is not returned
// as an error by any function.
var SkipFile = errors.New("skip this file")

// OptionFunc ...
type OptionFunc filepath.WalkFunc

// Ignore ...
func Ignore(pattern string) OptionFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if matched, _ := filepath.Match(pattern, info.Name()); matched {
			return SkipFile
		}
		return nil
	}
}
