// No req semicolons, except combining statements, decl
// Echo1 prints command-line args
package main

import (
	"fmt"
	"os"
)

// slice = dynam sized seq s
// of array elements
// s[i] indiv elements
// [m:n] contig subseq

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//!-
