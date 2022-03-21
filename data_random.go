package dbseed

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
	"time"
	"regexp"
	"strconv"
)

type Names struct {
	XMLName 	xml.Name 	`xml:"ol"`
	Names 		[]Name 		`xml:"li"`
	count 		int
}

type Name struct {
	XMLName 	xml.Name 	`xml:"li"`
	Name 		string		`xml:"div"`
	First 		string
	Last		string
}

var (
	phMin = 2020000000
	phMax = 9999999999

	isbn10Min = 1000000000
	isbn10Max = 9999999999

	isbn13Min = 1000000000000
	isbn13Max = 9999999999999

	ccMin = 1000000000000000
	ccMax = 9999999999999999
)

func DummyFmt() {
	fmt.Println("placeholder until fmt no longer needed.")
}

// Loads files saved from this webpage
// https://www.randomlists.com/random-names?qty=10000
func LoadNames(filepath string) Names {
	var startOfList = regexp.MustCompile(`<ol><li><div class="rand_large">`)
	var endOfList = regexp.MustCompile(`</li></ol>`)
	var reSpace = regexp.MustCompile(` `)
	var names Names

	f, ferr := os.Open(filepath)
	if ferr != nil {
		panic(ferr)
	}
	defer f.Close()

	fStat, statErr := f.Stat()
	if statErr != nil {
		panic(statErr)
	}

	bdata := make([]byte, fStat.Size())
	_, dataErr := f.Read(bdata)
	if dataErr != nil {
		panic(dataErr)
	}

	nameList1 := startOfList.Split(string(bdata), 2)
	nameList2 := endOfList.Split(nameList1[1], 3)
	nameList := nameList2[0]
	// Fixing XML to make it well formed for unmarshaling.
	nameList = "<ol><li><div>" + nameList + "</li></ol>"	

	xml.Unmarshal([]byte(nameList), &names)

	// Splitting the name string into first and last
	names.count = len(names.Names)
	for i := 0; i < names.count; i++ {
		firstLast := reSpace.Split(names.Names[i].Name, 3)
		names.Names[i].First = firstLast[0]
		names.Names[i].Last = firstLast[1]
	}

	return names
}

// Random first name generator.  I would like the names to be real names.
func (randPi *Names) RandNameFirst() string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	return randPi.Names[r1.Intn(len(randPi.Names))].First
}

// Random last name generator.  Real names are potentially better.
func (randPi *Names) RandNameLast() string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	return randPi.Names[r1.Intn(len(randPi.Names))].Last
}

// Random number generator for phone numbers.  Above 2* and below all 9's
func RandPhoneNumUS() string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r1.Intn(phMax - phMin) + phMin)
}

// ISBN10.  Same logic as phone numbers with different range.
func RandIsbn10() string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r1.Intn(isbn10Max - isbn10Min) + isbn10Min)
}

// ISBN13.  Same logic as ISBN with the potential for a trailing X.
func RandIsbn13() string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r1.Intn(isbn13Max - isbn13Min) + isbn13Min)
}

// CC Number.  Random number generator.  Don't care what the range is.
// Should return all digits
// Length of number should be 16
func RandCCNum() string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r1.Intn(ccMax - ccMin) + ccMin)
}

// Blurb.  Totally random text.
// Min length: 0
// Max length: 2000
func RandBlurb() string {
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	length := r1.Intn(1000) + 1000
	
	var byteLetters []byte

	for i := 0; i < length; i += 2 {
		byteLetters = append(byteLetters, byte(r1.Intn(90 - 65) + 65))
		byteLetters = append(byteLetters, byte(r1.Intn(122 - 97) + 97))
		if r1.Intn(10) == 5 {
			i++	// Increment the counter by one to maintain length
			byteLetters = append(byteLetters, byte(32))	// Adding a space character.
		}
	}

	return string(byteLetters)
}

// Returns a random integer
func RandInt() int {

	return 1
}

// Returns a random decimal
func RandMoney() float64 {
	return 0.0
}

// Returns a random date
func RandDate() time.Time {
	return time.Now()
}

var topLevelDomains = [...]string {
	"com",
	"us",
	"eu",
	"de",
	"org",
	"net",
}

var domains = [...]string {
	"gmail",
	"yahoo",
	"outlook",
	"facebook",
	"apple",
	"microsoft",
}

// Uses the name Names type for more realistic emails.
func (randPi *Names) RandEmail() string {
	var email, firstName, lastName, domain, tld string
	var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
	
	// Write an if statement to determine if a first name will be used
	if r1.Intn(100) > 5 {
		firstName = randPi.Names[r1.Intn(randPi.count)].First
	}

	// Write an if statement if a last name should be used.
	if r1.Intn(100) > 15 {
		lastName = randPi.Names[r1.Intn(randPi.count)].Last
	}

	if firstName == "" && lastName == "" {
		firstName = randPi.Names[r1.Intn(randPi.count)].First
	}

	// First name first or lastname first.
	if r1.Intn(1) == 1 {
		if firstName == "" {
			email += lastName
		} else {
			email += lastName + "." + firstName
		}
	} else {
		if lastName == "" {
			email += firstName
		} else {			
			if firstName == "" {
				email += lastName
			} else {
				email += firstName + "." + lastName
			}
		}
	}

	domain = domains[r1.Intn(len(domains))]
	tld = topLevelDomains[r1.Intn(len(topLevelDomains))]

	email += "@" + domain + "." + tld

	return email
}

func DataType(colType string) string  {
	fmt.Printf("colType: %v\n", colType)

	return "colType will go here"
}
