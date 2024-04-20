package main

import "fmt"

func main() {
	{
		fmt.Println(longestPalindrome("babad") == "bab" || longestPalindrome("babad") == "aba")
	}

	{
		fmt.Println(longestPalindrome("dfdftaratino") == "tarat")
	}

	{
		fmt.Println(longestPalindrome("cbbd") == "bb")
	}

	{
		fmt.Println(longestPalindrome("a") == "a")
	}

	{
		fmt.Println(longestPalindrome("ccc") == "ccc")
	}

	{
		fmt.Println(longestPalindrome("aaaa") == "aaaa")
	}

	{
		fmt.Println(longestPalindrome("aaaaaa") == "aaaaaa")
	}

	{
		fmt.Println(longestPalindrome("caaaas") == "aaaa")
	}

	{
		fmt.Println(longestPalindrome("caataas") == "aataa")
	}

	{
		fmt.Println(longestPalindrome("xaabacxcabaaxcabaax") == "xaabacxcabaax")
	}

	{
		fmt.Println(longestPalindrome("aacabdkacaa") == "aca")
	}
}

func longestPalindrome(s string) string {
	s2 := make([]byte, len(s)*2+1)

	for i := 0; i != len(s); i++ {
		s2[i*2+1] = s[i]
	}

	p := make([]int, len(s2))
	var r, c, index, iMir int
	maxLen := 1

	for i := 1; i != len(s2)-1; i++ {
		iMir = 2*c - i
		if r > i {
			if p[iMir] < r-i {
				p[i] = p[iMir]
			} else {
				p[i] = r - i
			}
		} else {
			p[i] = 0
		}

		for i > p[i] && (i+p[i]+1) < len(s2) && s2[i-p[i]-1] == s2[i+p[i]+1] {
			p[i]++
		}

		if p[i]+i > r {
			c = i
			r = i + p[i]
		}

		if maxLen < p[i] {
			maxLen = p[i]
			index = i
		}
	}

	return s[(index-maxLen)/2 : (index-maxLen)/2+maxLen]
}
