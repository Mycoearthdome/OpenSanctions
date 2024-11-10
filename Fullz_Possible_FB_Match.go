package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
	"sync"
)

const maxCapacity= 1024 * 1024 // 1MB

var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
    entitiesFile := "entities.ftm.json"
    personsFile := "Reconciled_Persons.txt"
    outFile := "Probable_FB_Sanctions.txt"

    fmt.Println("Reading FOUND POI file...Please wait")
    personsNames, err := readLines(personsFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    personsList := make([][]string, 0)
    for _, entry := range personsNames {
        if entry != "]" {
			//fmt.Println(entry)
            entry = strings.TrimSpace(entry)
            parts := strings.Split(entry, " ")
			//fmt.Println(parts)
            name := parts[0] + " " + parts[1]
            personsList = append(personsList, []string{name, strings.Join(parts[2:], ",")})
        }
    }

    fmt.Println("Loading Entities data...Please wait!")
    entitiesEntries, err := readLines(entitiesFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Concluding Hunt...Please wait!")

    for _, infos := range personsList {
		wg.Add(1)
        go HuntConclusions(infos, entitiesEntries, outFile, &wg)
    }
	wg.Wait()
}

func HuntConclusions(infos []string, entitiesEntries []string, outFile string, wg *sync.WaitGroup){
	defer wg.Done()
	for _, entry := range entitiesEntries {
		if strings.Contains(entry, infos[0]) {
			fmt.Printf("+NAME=%s\n[%s]\n+Facebook->%s\n\n", infos[0], entry, infos[1])
			writeToFile(outFile, fmt.Sprintf("+NAME=%s\n%s\n+Facebook->%s\n\n", infos[0], entry, infos[1]))
			break
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

func writeToFile(filename, content string) {
	mu.Lock()
	defer mu.Unlock()
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
