Package blake512
=====================

	import "github.com/dchest/blake512"

Package blake256 implements BLAKE-512 and BLAKE-384 hash functions (SHA-3
candidate).

Public domain / Creative Commons Zero.


Constants
---------

``` go
const BlockSize = 128
```
The block size of the hash algorithm in bytes.

``` go
const Size = 64
```
The size of BLAKE-512 hash in bytes.

``` go
const Size384 = 48
```
The size of BLAKE-384 hash in bytes.


Functions
---------

### func New

	func New() hash.Hash

New returns a new hash.Hash computing the BLAKE-512 checksum.

### func New384

	func New384() hash.Hash

New384 returns a new hash.Hash computing the BLAKE-384 checksum.

### func New384Salt

	func New224Salt(salt []byte) hash.Hash

New384Salt is like New384 but initializes salt with the given 32-byte slice.

### func NewSalt

	func NewSalt(salt []byte) hash.Hash

NewSalt is like New but initializes salt with the given 32-byte slice.
