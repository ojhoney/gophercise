package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//---------
// Fast IO
//---------

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

func init() {
	sample := ``

	sc = bufio.NewScanner(strings.NewReader(sample))
	//sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	sc.Split(bufio.ScanWords)
}

func scanInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}
