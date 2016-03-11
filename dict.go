package gohash

import (
	"crypto/sha256"
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
func (d *Dictionary) Find() (string, error) {

	err := d.decidePossibleAlgos()
	if err != nil {
		return "", err
	}

	fmt.Println("possibly using", d.possibleAlgos)

	for _, line := range d.lines {
		for _, algo := range d.possibleAlgos {
			if d.equals(algo, []byte(line)) {
				return line, nil
			}
		}
	}

	return "", nil
}

func (d *Dictionary) equals(algo string, buffer []byte) bool {

	// XXX use a  map[string]()func instead. first do that in hasher.go equals()
	if algo == "sha256" && byte32ArrayEquals(sha256.Sum256(buffer), d.expected) {
		return true
	}
	return false
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
