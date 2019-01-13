package SDCC_esercizio

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {

	/****************************
		// First element in os.Args is always the program name,
		// So we need at least 2 arguments to have a file name argument.

		if len(os.Args) < 2 {
			fmt.Println("Missing parameter, provide file name!")
			return
		}
	****************************/

	contentFile := ReadFile("testoA.txt")
	MapFileWords(contentFile)

	//PrintFile(os.Args[1])
	//PrintFile(os.Args[2])

}

/*ReadFile : it reads a file and return its content*/
func ReadFile(fileName string) string {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Can't read file:", fileName)
		panic(err)
	}
	PrintStatusFile(string(data), "ReadFile")
	return string(data)
}

/*MapFileWords : it counts the words of a string  ***********************/
func MapFileWords(contentFile string) map[string]int {

	onlyWords := RemoveAllNonAlphanumericCharacters(contentFile)
	PrintStatusFile(onlyWords, "CountFileWords 1")

	onlyLowerWords := strings.ToLower(onlyWords)
	PrintStatusFile(onlyLowerWords, "CountFileWords 2")

	separetedWords := strings.Fields(onlyLowerWords)
	for i := range separetedWords {
		fmt.Println(separetedWords[i])
	}

	fmt.Println(" ***************** Start map ***************** ")
	mappedWords := make(map[string]int)

	for _, word := range separetedWords {
		mappedWords[word]++
	}

	for index, element := range mappedWords {
		fmt.Println("Index =", index, "--- Element =", element)
	}

	return mappedWords
}

/*SortFileWords : **********************/
func SortFileWords(contentFile string) {

	// mapWords := WordCount(onlyWords)
	// fmt.Println(" ========= The number of words in the file is :", len(mapWords))

	/*
		// To store the keys in slice in sorted order
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		//sort.Ints(keys)

		// To perform the opertion you want
		for _, k := range keys {
			fmt.Println("Key:", k, "Value:", m[k])
		}
	*/

}

/*RemoveAllNonAlphanumericCharacters : ssssss*/
func RemoveAllNonAlphanumericCharacters(myString string) string {

	// Make a Regex to say we only want
	reg, err := regexp.Compile("[^a-zA-Z0-9\t\n\f\r ]+")
	if err != nil {
		log.Fatal(err)
	}
	// processedString
	myString = reg.ReplaceAllString(myString, "")

	return myString

}

/*PrintStatusFile : ffffffffffffffffffffffffffffffff*/
func PrintStatusFile(contentFile string, functionLocated string) {

	fmt.Printf(" ====================== FILE CONTENT  ===================== \n")
	fmt.Printf(" ====================== FUNCTION %s ======================= \n", functionLocated)
	fmt.Printf("%s", contentFile)
	fmt.Printf(" ========================================================== \n")

}
