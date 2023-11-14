package main

import (
	"fmt"
	"strings"
)

func main() {

	// INTRODUCTION TO STRINGS
	// Unlike javascript or python, strings in Go are written in double quotes.
	// What ever is in the quotes is called string literal.

	// var s1 = "Learning Golang" // using the var keyword
	// s2 := "Enjoying Golang"    // using the shorthand symbol

	//_, _ = s1, s2

	// Trying to use double quotes inside a string literal double quotes must be escaped by using (\)
	//s2 = "He says : \"Xup, you gaming\"\n" // He says : "Xup, you gaming"

	// Another way , which isn't quite hard is to use single quotes inside a string literal double quotes
	// that would be easier than using (\)
	//s1 = "He says : 'Xup, you gaming'\n"

	// Another way is to use double quotes inside a backticks (``)
	//s3 := `She said : "Hey, That's my man right there."`
	//fmt.Println(s1, s2, s3)

	//STRING CONCATENATION

	//Strings are immutable so addition of 2 or more strings using (+) will create a new string
	// s3 = s1 + " " + s2
	// fmt.Println(s3)

	// Trying to change a value of a string at any index will not be possible because strings are immutable
	//e.g
	//fmt.Println(s1[0]) // 'H' form (s1 = "He says : 'Xup, you gaming'\n")

	//s1[0] = I // Not Possible because string is immutable

	// Using %v,%s, %q in Printf to print a string
	//fmt.Printf("%v", s1)

	//
	//
	//

	//................................................................

	// Learning and Coding Runes and Strings
	//a, b := 'G', 'H' // Rune which infact is an aliais of int32.

	// In Golang, a string is a sequence of byte and  not characters
	// ss := "Golang"
	// fmt.Println(len(ss)) //  the result is 6. because ss contains 6 bytes of  ascii letters.

	// sd := "Codruța"
	// fmt.Println(len(sd)) // the result is 8, rune ț holds 2 bytes

	//
	//
	// To return the number of runes(actual character) in a string and not bytes is also possible using a utf-8 function
	// se := utf8.RuneCountInString(sd) // prints 7
	// sf := utf8.RuneCountInString(ss) // prints 6
	// fmt.Println(se, sf)

	//NB: when the function len() is used on a string, what it retuns is the number of bytes in that string,
	// normally, each character in ascii letters occupy one byte so it might also give the number of characters
	// but if the string holds one or more non-ASCII characters, the number returned from len() will not be equal
	// to the number of characters in the string , instead using the utf8.RuneCountInString() function is the correct way
	// to get the accurate number of characters in that string.

	//................................................................
	//................................................................
	//................................................................
	//................................................................

	//SLICING STRINGS.
	//................................................................

	// Slicing of a string returns the number of bytes in that string and not the characters, but
	// for ascii characters, the number of characters will be displayed e.g
	// aa := "I love Golang"
	// fmt.Println(aa[2:5]) // lov // but these are actually the unicode bytes

	// for a non-ASCII string
	// ab := "汉字/漢字"
	// fmt.Println(ab[0:3]) // 汉. this returned a new rune entirely, that has one byte.

	// so to actually return a slice of runes(characters), and not the slice of bytes, we have to take 2 steps to achieve that:

	// STEP 1:
	// We convert the string which was a slice of bytes to a slice of runes
	// ac := []rune(ab)
	// fmt.Println(string(ac)) // "汉字/漢字"

	// STEP 2:
	// slicing the rune slice
	//fmt.Println(string(ac[0:3])) // "汉字/"

	//N.B: This method is not effecient because each time these steps are executed, a new back array is created.

	//................................................................
	//................................................................
	//................................................................
	//................................................................

	//STANDARD PACKAGE FUNCTIONS ON STRINGS UTF 8

	// 1 - contains (to check wether a substring is within a string)

	p := fmt.Println

	// scd := strings.Contains("I love Golang", "love")
	// sfe := strings.Contains("I love Golang", "hate")

	//p(scd) // true. the word love is in "I love Golang"

	//p(sfe) // false

	// 2- ContainsAny (to check wether a character in a substring is within a string)
	//sfe = strings.ContainsAny("I love Golang", "f")   // false
	//sfe = strings.ContainsAny("I love Golang", "gft") // true why? g is in the string

	// 3- conatinsRune
	//sfe = strings.ContainsRune("I love Golang", 'f') // false
	//p(sfe)

	// 4 - Count - retuns the number of occurrences of a substring in the string
	// n := strings.Count("I love Golang", "o")
	// p(n)

	// 5- toLower and toUpper
	// newS := strings.ToLower("I love Golang")

	// newS = strings.ToUpper("I love Golang")
	// p(newS)

	// 6 - Compating strings
	// Convert both strings you want to compare to lowercase strings or uppercase strings then use the == sign to return
	// true or false
	//p(strings.ToLower("GO") == strings.ToLower("go")) // This method is not efficient

	//7 - Repeat
	myStr := strings.Repeat("#", 10) // ########## (prints the strings 10 times)

	//
	//
	//

	// Instead a function called EqualFold is used
	//p(strings.EqualFold("go", "GO")) // true. #####turns case insensitivity off.
	// BUT if you wanna check the sensitivity , just use  == sign
	//p("GO" == "go") // false

	//
	//
	//
	// MANIPULATION OF STRINGS

	//8 - Replace
	//myStr = "123.5657.364734.57684.2637"
	// myStr = strings.Replace(myStr, ".", "-", 3) // "123-5657-364734-57684.2637"

	//NB - if the no of occurrences is -1, it replaces all the . to -.
	// Another way to replace all is to use the ReplaceAll function

	// 9 - Split
	//myStr = strings.Split(myStr, "-")[0]
	//p(myStr) // 123

	// 10 - Join
	// newStr := []string{"I", "am", "learning", "Golang"}
	// p(strings.Join(newStr, " ")) // "I am learning Golang"

	// 11 - Field // will automatically split the string at the space character (" "). a split function on stereiod
	myStr = "I am learning Golang"
	p(strings.Fields(myStr)[3]) // "Golang"

}
