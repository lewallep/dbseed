package dbseed

import (
	"fmt"
	"testing"
	"strconv"
	"regexp"

	"dbbeagle/finder"
)

func TestPrintPlaceholder(t *testing.T) {
	fmt.Println("Testing has begun.")
}

func TestPhoneNumbers (t *testing.T) {
	for i := 0; i < 10000; i++ {
		phoneNum, err := strconv.Atoi(RandPhoneNumUS())
		if err != nil {
			panic(err)
		}
		if phoneNum < phMin || phoneNum > phMax {
			t.Errorf("The phone number generated was out of range: %d\n", phoneNum)
		}
	}
}

var invalidPhoneNumsLow = []int {
	2000000000,
	1999999999,
	0000000000,
	0,
	958478,
	911,
}

// Ensure that logic in above test is correct.  Find cases outside desired range
func TestPhoneNumRangeLow (t *testing.T) {
	for i := 0; i < len(invalidPhoneNumsLow); i++ {
		if invalidPhoneNumsLow[i] > phMin {
			t.Errorf("phoneNumber low logic has a bug.  Num tested was: %d\n", invalidPhoneNumsLow[i])
		}
	}
}

var invalidPhoneNumsHigh = []int {
	10000000000,
	9999999999,
	10000000001,
	99999999999,
	22222222222,
	1000000000000000,
}

func TestPhoneNumRangeHigh (t *testing.T) {
	for i := 0; i < len(invalidPhoneNumsHigh); i++ {
		if invalidPhoneNumsHigh[i] < phMax {
			t.Errorf("phoneNumber high logic has a bug.  Num tested was: %d\n", invalidPhoneNumsHigh[i])
		}
	}
}

var validPhoneNums = []int {
	5032301206,
	7025673409,
	9097659876,
	2026753345,
	8672843785,
	2031111111,
	2030000000,
	2020000000,
}

func TestPhoneNumValid (t *testing.T) {
	for i := 0; i < len(validPhoneNums); i++ {
		if validPhoneNums[i] < phMin || validPhoneNums[i] >= phMax {
			t.Errorf("phoneNumber logic is not correct for valid phone numbers.  Num tested was: %d\n", validPhoneNums[i])
		}
	}
}

func TestRandPhoneNums (t *testing.T) {
	// t.Parallel()
	for i := 0; i < 5000000; i++ {
		phNum := RandPhoneNumUS()
		if  _, err := strconv.Atoi(phNum); err != nil || phNum < "202000000" || phNum > "999999999" || len(phNum) != 10 {
			t.Errorf("RandPhoneNumUS() out of range: %s\n", phNum)
		}
	}
}

var isbn10Valid = []string {
	"1000000001",
	"8575728364",
	"7777777777",
	"3823724375",
	"9999999999",
}

func TestIsbn10Valid (t *testing.T) {
	for i := 0; i < len(isbn10Valid); i++ {
		if len(isbn10Valid[i]) < 10 || isbn10Valid[i] <= "1000000000" || isbn10Valid[i] > "9999999999" {
			t.Errorf("Isbn10 logic is not working.  Isbn10 num: %s\n", isbn10Valid[i])
		}
	}
}

func TestIsbn10ValidRand (t *testing.T) {
	for i := 0; i < 5000000; i++ {
		num := RandIsbn10()
		if len(num) < 10 || num <= "1000000000" || num > "9999999999" {
			t.Errorf("RandIsbn10 generated an invalid isbn10: %s\n", num)
		}
	}
}

var isbn10Invalid = []string {
	"",
	"1",
	"2",
	"adfgfauiejs",
	"gfhjrfhhnytgnjfgljkhbdlfkjahbsdjh",
	"4756757a96",
	"547Z674857",
}

func TestIsbn10Invalid (t *testing.T) {
	for i := 0; i < len(isbn10Invalid); i++ {
		if len(isbn10Invalid[i]) == 10 && isbn10Invalid[i] < "0000000000" && isbn10Invalid[i] > "9999999999" {
			t.Errorf("Valid isbn10 detected when it should be invalid: %s\n", isbn10Invalid[i])
		}

		if _, err := strconv.Atoi(isbn10Invalid[i]); err == nil && len(isbn10Invalid[i]) == 10 {
			t.Errorf("Valid isbn10 detected with correct length and only numersl: %s\n", isbn10Invalid[i])
		}
	}
}

func TestIsbn13Rand (t *testing.T) {
	var isbn string
	var err error
	var isbnInt int

	for i := 0; i < 100000; i++ {
		isbn = RandIsbn13()
		isbnInt, err = strconv.Atoi(isbn)
		if err != nil {
			t.Errorf("Error converting string to int: %v\n", err)
		}

		if isbnInt < isbn13Min || isbnInt > isbn13Max || len(isbn) != 13 {
			t.Errorf("ERROR RandIsbn13 generated: %s\n", isbn)
		}
	}
}

func TestCCNumRand (t *testing.T) {
	var ccNum string
	var err error
	var ccNumInt int

	for i := 0; i < 1000000; i++ {
		ccNum = RandCCNum()
		ccNumInt, err = strconv.Atoi(ccNum)
		if err != nil {
			t.Errorf("ERROR converting CCNum string to int: %v\n", err)
		}
		if ccNumInt < ccMin || ccNumInt > ccMax || len(ccNum) != 16 {
			t.Errorf("ERROR RandCCNum() produced an incorrect number: %s\t len: %d\n", ccNum, len(ccNum))
		}
	}
}

func TestBlurbRand (t *testing.T) {
	
	for i := 0; i < 10000; i++ {
		blurb := RandBlurb()
		len := len(blurb)
		if len > 2000 || len < 1000 {
			t.Errorf("RandBlurb length incorrect, len: %d\n", len)
		}
	}
}

func TestEmailRand(t *testing.T) {
	names := LoadNames("random_names_01.html")
	rEmail := regexp.MustCompile(finder.REGEX_EMAIL)

	for i := 0; i < 1000000; i++ {
		email := names.RandEmail()
		if !rEmail.Match([]byte(email)) {
			t.Errorf("None email: %s\n", email)
		}	
	}
}