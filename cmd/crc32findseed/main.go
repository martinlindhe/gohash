package main

import (
	"fmt"
	"hash/crc32"
	"math"
)

// none of the supported crc32 hashes seem to correspond to the one php calls "crc32"
// this tool figures out what seed was used to generate php-src/ext/hash/php_hash_crc32_tables.h

func main() {

	for i := uint32(0); i < math.MaxUint32; i++ {
		tbl := crc32.MakeTable(i)
		if tbl[0] != 0 {
			fmt.Println("first value non-zero")
			continue
		}

		if tbl[1] == 0x04c11db7 {
			fmt.Printf("Possible match %08x", i)
			fmt.Println()
		}

		if i%100000 == 0 {
			fmt.Printf("it %08x, t1 = %08x", i, tbl[1])
			fmt.Println()
		}
	}
}
