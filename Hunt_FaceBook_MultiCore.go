package main

import (
    "bufio"
    "fmt"
    "os"
    "sync"
    "strings"
    "math"
)

const (
    PersonsFile = "Persons_Names_Alias.txt"
    FaceBookFile = "FaceBook_2019.txt"
    OutFile = "Persons_Found.txt"
    MAX_RESULTS = 5
    NB_CORES = 16
)

var (
    mu sync.Mutex
)

type Result struct {
    Person string
    Entries []string
}

type Work struct {
    Person string
    Chunk []string
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

func worker(workChan chan Work, results chan Result, wg *sync.WaitGroup) {
    defer wg.Done()
    for work := range workChan {
        counter := 0
        var found []string = []string{}
        for _, entry := range work.Chunk{
            if strings.Contains(entry, strings.ReplaceAll(work.Person, " ", ":")) {
                counter++
                found = append(found, entry)
                if counter == MAX_RESULTS {
                    break
                }
            }

        }

        if counter > 0 && counter < MAX_RESULTS{
            results <- Result{work.Person, found}
        } else {
            results <- Result{"", nil}
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func writeResults(results []string, outFile string) error {
    mu.Lock()
    defer mu.Unlock()

    file, err := os.OpenFile(outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    for _, result := range results {
        _, err := file.WriteString(result + ":")
        if err != nil {
            return err
        }
        _, err = file.WriteString("\n")
        if err != nil {
            return err
        }
    }

    return nil
}

func removeDuplicates(personsNames []string) []string {

    var duplicateFree []string
    var duplicate bool

    personsList := []string{}
    for _, entry := range personsNames {
        parts := strings.Split(entry, ":")
        if len(parts) > 1 {
            personName := strings.TrimSpace(parts[1])
            personsList = append(personsList, personName)
        }
    }

    for _, single := range personsList {
        duplicate = false
        for _, dup := range duplicateFree{
            if single == dup{
                duplicate = true
                break
            }
        }
        if !duplicate{
            duplicateFree = append(duplicateFree, single)
        }
    }

    file, err := os.OpenFile(PersonsFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        panic(err)
    }
    defer file.Close()

    for _, person := range duplicateFree{
        _, err = file.WriteString(person+"\n")
        if err != nil {
            fmt.Println("Error writing to file:", err)
            panic(err)
        }
    }
    return duplicateFree
}

func main() {
    fmt.Println("Reading POI file...Please wait")
    personsNames, err := loadFile(PersonsFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Loading FB_2019 data...Please wait!")
    facebookEntries, err := loadFile(FaceBookFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Hunting...Please wait!")

    var wg sync.WaitGroup
    var personsTotal int64 = int64(len(personsNames))
    var resultCount int = NB_CORES
    
    results := make(chan Result, personsTotal)
    workChan := make(chan Work, personsTotal)
    
    CHUNK_SIZE := int(math.Floor(float64(len(facebookEntries) / NB_CORES)))

    for i := 0; i < NB_CORES; i++ {
        wg.Add(1)
        go worker(workChan, results, &wg)
    }

    var count int64 = 0

    for _, personName := range personsNames {
        var resultSlice []string
        for i := 0; i < len(facebookEntries); i += CHUNK_SIZE {
            chunk := facebookEntries[i:min(i+CHUNK_SIZE, len(facebookEntries))]
            workChan <- Work{personName, chunk}
        }
        //TradeOff for RAM limitations (REMOVE if you have enough)
        for j:=0;j<resultCount;j++{
            select {
            case result := <-results:
                if result.Entries != nil{
                    for _, entry := range result.Entries{
                        resultSlice = append(resultSlice, entry)
                        //fmt.Println(entry)
                    }
                    //fmt.Println("\nWriting to file...", result.Person)
                    err = writeResults(resultSlice, OutFile)
                    if err != nil {
                        fmt.Println(err)
                    }
                    clear(resultSlice)
                    resultCount--
                    
                }
            }
            
        }
        count++
        resultCount = NB_CORES
				
		fmt.Printf("\rStatus: %d/%d", count, personsTotal)
    }

    close(workChan)
    wg.Wait()

    close(results)

    for result := range results {
        var resultSlice []string
        if result.Entries != nil {
            for _, entry := range result.Entries {
                resultSlice = append(resultSlice, entry)
            }
            err = writeResults(resultSlice, OutFile)
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}
