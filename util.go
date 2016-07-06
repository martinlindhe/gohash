package gohash

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	termutil "github.com/andrew-d/go-termutil"
)

// used in tests
type expectedForms map[string]string

// strings used in tests
var (
	blank = ""
	fox   = "The quick brown fox jumps over the lazy dog"
)

func byteArrayEquals(a []byte, b []byte) bool {

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isByteInSlice(a byte, list []byte) bool {

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func strToDistinctByteSlice(s string) []byte {

	res := []byte{}

	ptr := 0
	for i := 0; i < len(s); i++ {
		if isByteInSlice(s[i], res) {
			continue
		}
		res = append(res, s[i])
		ptr++
	}

	// sort it too
	sort.Sort(byteSlice(res))

	return res
}

// byteSlice implements sort.Interface to sort a []byte
type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }

// AppInputData is the captured input to the app, either from a pipe or a file
type AppInputData struct {
	Data   []byte
	IsPipe bool
}

// ReadPipeOrFile reads from stdin if pipe exists, else from provided file
func ReadPipeOrFile(fileName string) (*AppInputData, error) {

	res := AppInputData{}

	if !termutil.Isatty(os.Stdin.Fd()) {
		res.Data, _ = ioutil.ReadAll(os.Stdin)
		res.IsPipe = true
	} else {
		if fileName == "" {
			return nil, fmt.Errorf("no piped data and no file provided")
		}
		var err error
		res.Data, err = readBinaryFile(fileName)
		if err != nil {
			return &res, err
		}
	}
	return &res, nil
}

func readBinaryFile(filename string) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	size := stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}
