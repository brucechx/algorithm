package kmp

import (
	"errors"
	"fmt"
)

type KMP struct {
	pattern string
	next  []int
	size    int
}

// compile new next-array given argument
func NewKMP(pattern string) (*KMP, error) {
	prefix, err := next(pattern)
	if err != nil {
		return nil, err
	}
	return &KMP{
			pattern: pattern,
			next:  prefix,
			size:    len(pattern)},
		nil
}

// return index of first occurrence of kmp.pattern in argument 's'
// - if not found, returns -1
func (kmp *KMP) FindStringIndex(s string) int {
	if len(s) < kmp.size{
		return -1
	}
	i, j := 0, 0
	sLen := len(s)
	pLen := kmp.size
	for i < sLen && j < pLen{
		if j == -1 || s[i] == kmp.pattern[j]{
			i++
			j++
		}else {
			j = kmp.next[j]
		}
	}
	if j == pLen{
		return i - j
	}
	return -1

}


// find every occurence of the kmp.pattern in 's'
func (kmp *KMP) FindAllStringIndex(s string) []int {
	if len(s) < kmp.size{
		return []int{}
	}
	i, j := 0, 0
	sLen, pLen := len(s), kmp.size
	allIndex := make([]int, 0, 10)
	for i < sLen{
		for i<sLen && j<pLen{
			if j == -1 || s[i] == kmp.pattern[j]{
				i++
				j++
			}else {
				j = kmp.next[j]
			}
		}
		if j == pLen{
			allIndex = append(allIndex, i - j)
			i++
			j = 0
		}
	}
	return allIndex
}

// returns true if pattern i matched at least once
func (kmp *KMP) ContainedIn(s string) bool {
	return kmp.FindStringIndex(s) >= 0
}

// returns the number of occurrences of pattern in argument
func (kmp *KMP) Occurrences(s string) int {
	return len(kmp.FindAllStringIndex(s))
}

// For debugging
func (kmp *KMP) String() string {
	return fmt.Sprintf("pattern: %v\nnext: %v", kmp.pattern, kmp.next)
}

// 传统的next数组
func next(p string) ([]int, error) {
	pLen := len(p)
	if pLen < 2{
		if pLen == 0{
			return nil, errors.New("'pattern' must contain at least one character")
		}
		return []int{-1}, nil
	}
	next := make([]int, pLen, pLen)
	next[0] = -1
	next[1] = 0
	i, j := 0, 1  // i 表示前缀，j 表示后缀
	for j < (pLen - 1){
		if i == -1 || p[i] == p[j]{ //-1代表了起始位不匹配，i=0,s[0]!=s[j]=>i=next[0]=-1
			i++
			j++
			next[j] = i
		}else {
			i = next[i]
		}
	}
	return next, nil
}

// 优化后的next数组
func optimizeNext(p string) ([]int, error){
	pLen := len(p)
	if pLen < 2{
		if pLen == 0{
			return nil, errors.New("'pattern' must contain at least one character")
		}
		return []int{-1}, nil
	}
	next := make([]int, pLen, pLen)
	next[0] = -1
	next[1] = 0
	i, j := 0, 1  // i 表示前缀，j 表示后缀
	for j < (pLen - 1){
		if i == -1 || p[i] == p[j]{ //-1代表了起始位不匹配，i=0,s[0]!=s[j]=>i=next[0]=-1
			i++
			j++
			if p[i] != p[j]{ //因为出现在j位置不匹配的话会跳到next[j]=i位置去匹配,p[i] == p[j]肯定又是不匹配（优化核心点）
				next[j] = i
			}else {
				next[j] = next[i]
			}
		}else {
			i = next[i]
		}
	}
	return next, nil
}

// normal
func findStringIndex(s, p string) (int, error){
	i, j := 0, 0
	sLen := len(s)
	pLen := len(p)
	next, err := optimizeNext(p)
	if err != nil{
		return -1, err
	}
	for i < sLen && j < pLen{
		if j == -1 || s[i] == p[j]{
			i++
			j++
		}else {
			j = next[j]
		}
	}
	if j == pLen{
		return i - j, nil
	}
	return -1, errors.New("not exist")
}

func findAllStringIndex(s, p string) ([]int, error){
	i, j := 0, 0
	sLen, pLen := len(s), len(p)
	next, err := optimizeNext(p)
	if err != nil{
		return nil, err
	}
	allIndex := make([]int, 0, 10)
	for i < sLen{
		for i<sLen && j<pLen{
			if j == -1 || s[i] == p[j]{
				i++
				j++
			}else {
				j = next[j]
			}
		}
		if j == pLen{
			allIndex = append(allIndex, i - j)
			i++
			j = 0
		}
	}
	return allIndex, nil
}