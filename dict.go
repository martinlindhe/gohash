package gohash

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Dictionary is used to find cleartext for checksum in `expected`,
// using dictionary and any algorithm with matching bitsize
type Dictionary struct {
	dictFileName  string
	lines         []string
	expected      []byte
	possibleAlgos []string
	prefix        string
	suffix        string
}

// NewDictionary creates a new Dictionary
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

// Prefix sets a fixed prefix
func (d *Dictionary) Prefix(s string) { d.prefix = s }

// Suffix sets a fixed suffix
func (d *Dictionary) Suffix(s string) { d.suffix = s }

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
		if line == "" {
			continue
		}

		buf := []byte(d.prefix + line + d.suffix)

		for _, algo := range d.possibleAlgos {
			if d.equals(algo, &buf) {
				return line, algo, nil
			}
		}
	}

	return "", "", nil
}

func (d *Dictionary) equals(algo string, buffer *[]byte) bool {

	calc := NewCalculator(*buffer)
	return byteArrayEquals(*calc.Sum(algo), d.expected)
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
