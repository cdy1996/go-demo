package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v \n", addBinary("100", "110010"))
	fmt.Printf("%v \n", addBinary("11", "1"))
	fmt.Printf("%v \n", addBinary("1010", "1011"))
	fmt.Printf("%v \n", addBinary("", "1011"))
	fmt.Printf("%v \n", addBinary("", ""))

}

func addBinary(a string, b string) string {
	al := len(a) - 1
	bl := len(b) - 1

	add := false
	c := ""
	for al >= 0 && bl >= 0 {
		add, c = addAB(a[al], b[bl], add, c)
		al--
		bl--
	}

	if al >= 0 {
		for i := al; i >= 0; i-- {
			add, c = addZERO(add, a[i], c)
		}
	}
	if bl >= 0 {
		for i := bl; i >= 0; i-- {
			add, c = addZERO(add, b[i], c)
		}
	}
	if add {
		c = "1" + c
	}
	return c

}

func addAB(a, b uint8, add bool, c string) (bool, string) {
	if a == '1' && b == '1' {
		if add {
			c = "1" + c
		} else {
			c = "0" + c
		}
		add = true
	} else if a == '0' && b == '0' {
		if add {
			c = "1" + c
		} else {
			c = "0" + c
		}
		add = false
	} else {
		if add {
			c = "0" + c
			add = true
		} else {
			c = "1" + c
			add = false
		}
	}
	return add, c
}

func addZERO(add bool, a uint8, c string) (bool, string) {
	return addAB(a, '0', add, c)
}

func addBinaryBest(a string, b string) string {
	const base = 2
	m, n := len(a), len(b)
	if m > n {
		return addBinaryBest(b, a)
	}
	buf := make([]byte, n+1)
	carry := 0
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 {
			carry += int(a[j] - '0')
			j--
		}
		carry += int(b[i] - '0')
		buf[i+1] = byte(carry%base + '0')
		carry /= base
	}
	if carry == 0 {
		return string(buf[1:])
	}
	buf[0] = '1'
	return string(buf)
}
