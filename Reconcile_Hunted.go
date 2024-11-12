package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"sync"
)

const maxCapacity = 1024 * 2024 //1MB

func main() {
	var wg sync.WaitGroup

    personsFile := "Persons_Names_Alias.txt"
    personsFoundFile := "Persons_Found.txt"
    outFile := "Reconciled_Persons.txt"

    personsHunted, err := readLines(personsFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    personsFound, err := readLines(personsFoundFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    personNameCheck := make([]string, 0)
    for _, person := range personsHunted {
        if strings.Contains(person, "firstName: ") {
            personNameCheck = append(personNameCheck, strings.TrimSpace(strings.Split(person, "firstName: ")[1]))
        }
        if strings.Contains(person, "lastName: ") {
            personNameCheck = append(personNameCheck, strings.TrimSpace(strings.Split(person, "lastName: ")[1]))
        }
        if strings.Contains(person, "firstName:") {
            personNameCheck = append(personNameCheck, strings.TrimSpace(strings.Split(person, "firstName:")[1]))
        }
        if strings.Contains(person, "lastName:") {
            personNameCheck = append(personNameCheck, strings.TrimSpace(strings.Split(person, "lastName:")[1]))
        }
    }

    fmt.Println("Reconcilliation...")
    for _, fbDetails := range personsFound {
        if fbDetails != ":" {
            details := strings.Split(fbDetails, ":")
            wg.Add(1)
			go ProcessData(details, personNameCheck, fbDetails, outFile, &wg)
        }
    }
}

func ProcessData(details []string, personNameCheck []string, fbDetails string, outFile string, wg *sync.WaitGroup){
	defer wg.Done()

	if len(details) > 7 {
		for i, detail := range details {
			if detail == "male" || detail == "female" {
				firstName := strings.Replace(details[i-2], "\"", "",-1)
				lastName := strings.Replace(details[i-1], "\"", "",-1)
				if firstName != "None" {
					for j:=0;j<len(personNameCheck);j++{
						if strings.Contains(personNameCheck[j], firstName) || strings.Contains(personNameCheck[j], lastName) {
							fmt.Printf("Probable match --> %s %s [%s]\n", firstName, lastName, fbDetails)
							writeToFile(outFile, fmt.Sprintf("%s %s [%s]\n", firstName, lastName, fbDetails))
							break
						}
					}
				}
				break
			}
		}
	} else {
		if strings.Contains(details[0], "female") {
			details = strings.Split(strings.Split(details[0], "female")[0], ",")
		} else if strings.Contains(details[0], "male") {
			details = strings.Split(strings.Split(details[0], "male")[0], ",")
		} else {
			details = []string{details[0]}
		}
		if len(details) > 4{
			firstName := strings.Replace(details[len(details)-3], "\"", "", -1)
			lastName := strings.Replace(details[len(details)-2], "\"", "", -1)

			if firstName == "None" && lastName == "None" {
				firstName = strings.Replace(details[len(details)-5], "\"", "", -1)
				lastName = strings.Replace(details[len(details)-4], "\"", "", -1)
			}
			for j:=0;j<len(personNameCheck);j++{
				if strings.Contains(personNameCheck[j], firstName) || strings.Contains(personNameCheck[j], lastName) {
					fmt.Printf("Probable match --> %s %s [%s]\n", firstName, lastName, fbDetails)
					writeToFile(outFile, fmt.Sprintf("%s %s [%s]\n", firstName, lastName, fbDetails))
					break
				}
			}
		}
	}
}

func readLines(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
    scanner.Buffer(buf, maxCapacity) // Increase the buffer size
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func writeToFile(filename string, content string) {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        fmt.Println(err)
    }
}
