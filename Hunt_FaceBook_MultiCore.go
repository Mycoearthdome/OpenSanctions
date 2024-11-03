package main

import (
    "bufio"
    "fmt"
    "os"
    "sync"
	"strings"
)

const (
    PersonsFile = "Persons_Names_Alias.txt"
    FaceBookFile = "FaceBook_2019.txt"
    OutFile = "Persons_Found.txt"
    MAX_RESULTS = 5
)

type Result struct {
    Person string
    Entries []string
}

func loadFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    const maxCapacity = 1024 * 1024 // 1MB
    buf := make([]byte, maxCapacity)
    scanner.Buffer(buf, maxCapacity) // Increase the buffer size
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines, scanner.Err()
}

func processPerson(person string, facebookEntries []string, results chan Result, wg *sync.WaitGroup) {
    defer wg.Done()

    found := []string{}
    counter := 0
    for _, entry := range facebookEntries {
        if entry == person {
            found = append(found, entry)
            counter++
            if counter == MAX_RESULTS {
                break
            }
        }
    }

	if counter < MAX_RESULTS {
    	results <- Result{person, found}
	} else {
		results <- Result{person, []string{""}}
	}
}

func writeResults(results []Result, outFile string) error {
    file, err := os.OpenFile(outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    for _, result := range results {
        for _, entry := range result.Entries {
			if entry == "" {
				_, err := file.WriteString(entry + ":")
				if err != nil {
					return err
				}
			    _, err = file.WriteString("\n")
			    if err != nil {
				    return err
			    }
            }

        }
        
    }

    return nil
}

func main() {
    fmt.Println("Reading POI file...Please wait")
    personsNames, err := loadFile(PersonsFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    personsList := []string{}
    for _, entry := range personsNames {
        parts := strings.Split(entry, ":")
        if len(parts) > 1 {
            personName := strings.TrimSpace(parts[1])
            personsList = append(personsList, personName)
        }
    }

    fmt.Println("Loading FB_2019 data...Please wait!")
    facebookEntries, err := loadFile(FaceBookFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Hunting...Please wait!")

    var wg sync.WaitGroup
    results := make(chan Result, len(personsList))

    for _, person := range personsList {
        wg.Add(1)
        go processPerson(person, facebookEntries, results, &wg)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    var resultSlice []Result
    for result := range results {
        resultSlice = append(resultSlice, result)
    }

    err = writeResults(resultSlice, OutFile)
    if err != nil {
        fmt.Println(err)
    }
}
