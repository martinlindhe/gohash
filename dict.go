package gohash

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Dictionary ...
type Dictionary struct {
	dictFileName  string
	lines         []string
	expected      []byte
	possibleAlgos []string
}

// NewDictionary ...
func NewDictionary(dictFileName string) (*Dictionary, error) {

	data, err := ioutil.ReadFile(dictFileName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	return &Dictionary{
		dictFileName: dictFileName,
		lines:        lines,
	}, nil
}

// ExpectedHash sets the expected hash
func (d *Dictionary) ExpectedHash(expected string) {
	tmp := hexStringToBytes(expected)
	d.expected = tmp[:]
}

// Find ...
func (d *Dictionary) Find() (string, string, error) {

	err := d.decidePossibleAlgos()
	if err != nil {
		return "", "", err
	}

	fmt.Println("Trying with", d.possibleAlgos)

	for _, line := range d.lines {
		buf := []byte(line)
		for _, algo := range d.possibleAlgos {
			if d.equals(algo, &buf) {
				return line, algo, nil
			}
		}
	}

	return "", "", nil
}

func (d *Dictionary) equals(algo string, buffer *[]byte) bool {

	if equals, ok := algoEquals[algo]; ok {
		return equals(buffer, &d.expected)
	}

	// NOTE: ok to panic here, since code path can only occur
	// while adding a new algo to the lib
	panic(fmt.Errorf("Unknown algo %s", algo))
}

// derive possible hashes from bitsize
func (d *Dictionary) decidePossibleAlgos() error {

	bitSize := len(d.expected) * 8

	for algo, algoBitSize := range algos {
		if algoBitSize == bitSize {
			d.possibleAlgos = append(d.possibleAlgos, algo)
		}
	}

	if len(d.possibleAlgos) == 0 {
		return fmt.Errorf("No known hashes uses a bitsize of %d", bitSize)
	}

	return nil
}
