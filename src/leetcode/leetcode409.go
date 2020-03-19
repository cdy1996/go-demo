package main

import "fmt"

func main() {
	fmt.Printf("%v \n", longestPalindrome("bb"))
	//fmt.Printf("%v \n",longestPalindrome("abccccdd"))
	//fmt.Printf("%v \n", longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}

// 最长回文串
func longestPalindrome(s string) int {
	maps := map[rune]int{}
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		a, b := maps[r]
		if b {
			a++
			maps[r] = a
		} else {
			maps[r] = 1
		}
	}
	total := 0
	odd := 0
	for _, v := range maps {
		i := v % 2

		if i != 0 {
			odd++
		}
		total += v
	}
	if odd >= 2 {
		return total - odd + 1
	} else {
		return total
	}

}
