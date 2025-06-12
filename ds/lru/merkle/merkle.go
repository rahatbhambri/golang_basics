package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"log"
)

type Hashf interface {
	getHash(s string) []byte
}

type simpleHash struct {
	dateOfImplement string
}

type NewHash struct {
	dateOfImplement string
	modNum          int
}

func (sh simpleHash) getHash(s string) []byte {
	// log.Println("date implemented", sh.dateOfImplement)
	if len(s) == 0 {
		return []byte{}
	}
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return bs
}

func (mh NewHash) getHash(s string) []byte {
	// log.Println("date implemented", mh.dateOfImplement)
	if len(s) == 0 {
		return []byte{}
	}
	h := sha512.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return bs
}

func main() {
	// h1 := simpleHash{"21.10.2023"}
	h2 := NewHash{"31.02.2022", 3}
	calc(h2)
}
func calc(hf Hashf) {
	// recursive function which takes list of hashes and returns a list of merged hashes
	// do this till list len is 1

	s1 := []string{"Hi", "this", "i'm", "simple"}
	v1 := ComputeMerkle(s1, hf)

	s2 := []string{"Even", "Array"}
	v2 := ComputeMerkle(s2, hf)

	log.Println(v1, v2)
	log.Printf("%x\n %x\n, %t", v1, v2, bytes.Equal(v1, v2))
}

func ComputeMerkle(s1 []string, hf Hashf) []byte {
	if len(s1) == 0 {
		return []byte{}
	}
	var initHash [][]byte
	for _, s := range s1 {
		initHash = append(initHash, hf.getHash(s))
	}

	v := GetMerkle(initHash, hf)[0]
	return v
}
func GetMerkle(arr [][]byte, hf Hashf) [][]byte {
	n := len(arr)
	if n == 1 {
		return arr
	}

	if n%2 != 0 {
		arr = append(arr, arr[n-1])
	}

	m := len(arr)
	var arr2 [][]byte

	for i := 0; i < m; i += 2 {
		a := hf.getHash(string(arr[i]))
		b := hf.getHash(string(arr[i+1]))
		c := append(a, b...)
		arr2 = append(arr2, c)
	}

	return GetMerkle(arr2, hf)
}
