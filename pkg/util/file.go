package util

import (
	"bailuoxi66/go-loggie/pkg/core/log"
	"bufio"
	"bytes"
	"github.com/mattn/go-zglob"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// LineCountTo calculates the number of lines to the offset
func LineCountTo(offset int64, fileName string) (int, error) {
	r, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	buf := make([]byte, 64*1024)
	count := 0
	lineSep := []byte{'\n'}
	totalReadBytes := int64(0)

	for totalReadBytes < offset {
		c, err := r.Read(buf)
		gap := totalReadBytes + int64(c) - offset
		if gap > 0 {
			c = c - int(gap) + 1
		}
		count += bytes.Count(buf[:c], lineSep)

		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}
		totalReadBytes += int64(c)
	}
	return count, nil
}

// LineCount returns the number of file lines
// better
func LineCount(r io.Reader) (int, error) {
	buf := make([]byte, 64*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}
	}
}

// LineCount1 returns the number of file lines
// deprecated
func LineCount1(r io.Reader) (int, error) {
	fileScanner := bufio.NewScanner(r)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount, nil
}

func WriteFileOrCreate(dir string, filename string, content []byte) error {
	f := filepath.Join(dir, filename)
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				log.Panic("mkdir %s error: %v", dir, err)
			}
		}
		return err
	}
	return ioutil.WriteFile(f, content, os.ModePerm)
}

func GlobWithRecursive(pattern string) (matches []string, err error) {
	if strings.Contains(pattern, "**") {
		// recursive lookup
		matches, err = zglob.Glob(pattern)
	} else {
		matches, err = filepath.Glob(pattern)
	}
	return matches, err
}

func MatchWithRecursive(pattern, name string) (matched bool, err error) {
	if strings.Contains(pattern, "**") {
		// recursive lookup
		matched, err = zglob.Match(pattern, name)
	} else {
		matched, err = filepath.Match(pattern, name)
	}
	return matched, err
}
