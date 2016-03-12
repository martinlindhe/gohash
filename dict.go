package gohash

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
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

	// runtime stats
	try    uint64
	tick   uint64
	buffer []byte
	algo   string
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

// Find searches for a cleartext for expected hash
func (d *Dictionary) Find() (string, string, error) {

	err := d.decidePossibleAlgos()
	if err != nil {
		return "", "", err
	}

	fmt.Println("Trying with", d.possibleAlgos)

	go d.statusReport()

	for _, line := range d.lines {
		if line == "" {
			continue
		}

		d.buffer = []byte(d.prefix + line + d.suffix)

		guesses := [][]byte{
			d.buffer,
			reverse(d.buffer),
		}

		for _, algo := range d.possibleAlgos {
			d.algo = algo // XXX slow to copy in hot path

			for _, guess := range guesses {
				if d.equals(algo, &guess) {
					return line, algo, nil
				}
			}
		}
		d.try++
	}

	return "", "", nil
}

func (d *Dictionary) statusReport() {

	for {
		time.Sleep(1 * time.Second)
		d.tick++
		avg := d.try / d.tick

		fmt.Printf("%s ~%d/s %s\n", d.algo, avg, string(d.buffer))
	}
}

func reverse(b []byte) []byte {

	len := len(b)
	res := make([]byte, len)

	for i := 0; i < len; i++ {
		res[i] = b[len-1-i]
	}

	return res
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

	sort.Strings(d.possibleAlgos)

	return nil
}
