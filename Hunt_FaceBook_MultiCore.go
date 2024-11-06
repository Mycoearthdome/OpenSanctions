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
    CHUNK_SIZE = 1024
)

type Result struct {
    Person string
    Entries []string
}

type Work struct {
    Person string
    Chunk []string
}

var facebookEntriesPool = sync.Pool{
    New: func() interface{} {
        return make([]string, 0, 1024)
    },
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
	var work Work

	work = <- workChan
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
		//fmt.Println("FOUND! -->", work.Person)
		results <- Result{work.Person, found}
	} else {
		results <- Result{"", nil}
	}
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func writeResults(results []string, outFile string) error {
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
    var personsTotal int64 = int64(len(personsList))
    results := make(chan Result, personsTotal)
    done := make(chan bool, 16)
    
    workChan0 := make(chan Work)
    workChan1 := make(chan Work)
    workChan2 := make(chan Work)
    workChan3 := make(chan Work)
    workChan4 := make(chan Work)
    workChan5 := make(chan Work)
    workChan6 := make(chan Work)
    workChan7 := make(chan Work)
    workChan8 := make(chan Work)
    workChan9 := make(chan Work)
    workChan10 := make(chan Work)
    workChan11 := make(chan Work)
    workChan12 := make(chan Work)
    workChan13 := make(chan Work)
    workChan14 := make(chan Work)
    workChan15 := make(chan Work)
    
    count := int64(0)

    // Prepare work to workers
    for _, personName := range personsList {

        var resultSlice []string
        

        for i := 0; i + i < len(facebookEntries); i += CHUNK_SIZE * 16{
           
                
            chunk0 := facebookEntries[i:min(i+CHUNK_SIZE, len(facebookEntries))]
            chunk1 := facebookEntries[i + CHUNK_SIZE * 2:min(i+CHUNK_SIZE * 2, len(facebookEntries))]
            chunk2 := facebookEntries[i + CHUNK_SIZE * 3:min(i+CHUNK_SIZE * 3, len(facebookEntries))]
            chunk3 := facebookEntries[i + CHUNK_SIZE * 4:min(i+CHUNK_SIZE * 4, len(facebookEntries))]
            chunk4 := facebookEntries[i + CHUNK_SIZE * 5:min(i+CHUNK_SIZE * 5, len(facebookEntries))]
            chunk5 := facebookEntries[i + CHUNK_SIZE * 6:min(i+CHUNK_SIZE * 6, len(facebookEntries))]
            chunk6 := facebookEntries[i + CHUNK_SIZE * 7:min(i+CHUNK_SIZE * 7, len(facebookEntries))]
            chunk7 := facebookEntries[i + CHUNK_SIZE * 8:min(i+CHUNK_SIZE * 8, len(facebookEntries))]
            chunk8 := facebookEntries[i + CHUNK_SIZE * 9:min(i+CHUNK_SIZE * 9, len(facebookEntries))]
            chunk9 := facebookEntries[i + CHUNK_SIZE * 10:min(i+CHUNK_SIZE * 10, len(facebookEntries))]
            chunk10 := facebookEntries[i + CHUNK_SIZE * 11:min(i+CHUNK_SIZE * 11, len(facebookEntries))]
            chunk11 := facebookEntries[i + CHUNK_SIZE * 12:min(i+CHUNK_SIZE * 12, len(facebookEntries))]
            chunk12 := facebookEntries[i + CHUNK_SIZE * 13:min(i+CHUNK_SIZE * 13, len(facebookEntries))]
            chunk13 := facebookEntries[i + CHUNK_SIZE * 14:min(i+CHUNK_SIZE * 14, len(facebookEntries))]
            chunk14 := facebookEntries[i + CHUNK_SIZE * 15:min(i+CHUNK_SIZE * 15, len(facebookEntries))]
            chunk15 := facebookEntries[i + CHUNK_SIZE * 16:min(i+CHUNK_SIZE * 16, len(facebookEntries))]

            wg.Add(16)

            go worker(workChan0, results, &wg)
            go worker(workChan1, results, &wg)
            go worker(workChan2, results, &wg)
            go worker(workChan3, results, &wg)
            go worker(workChan4, results, &wg)
            go worker(workChan5, results, &wg)
            go worker(workChan6, results, &wg)
            go worker(workChan7, results, &wg)
            go worker(workChan8, results, &wg)
            go worker(workChan9, results, &wg)
            go worker(workChan10, results, &wg)
            go worker(workChan11, results, &wg)
            go worker(workChan12, results, &wg)
            go worker(workChan13, results, &wg)
            go worker(workChan14, results, &wg)
            go worker(workChan15, results, &wg)

            workChan0 <- Work{personName, chunk0}
            workChan1 <- Work{personName, chunk1}
            workChan2 <- Work{personName, chunk2}
            workChan3 <- Work{personName, chunk3}
            workChan4 <- Work{personName, chunk4}
            workChan5 <- Work{personName, chunk5}
            workChan6 <- Work{personName, chunk6}
            workChan7 <- Work{personName, chunk7}
            workChan8 <- Work{personName, chunk8}
            workChan9 <- Work{personName, chunk9}
            workChan10 <- Work{personName, chunk10}
            workChan11 <- Work{personName, chunk11}
            workChan12 <- Work{personName, chunk12}
            workChan13 <- Work{personName, chunk13}
            workChan14 <- Work{personName, chunk14}
            workChan15 <- Work{personName, chunk15}

            // Process results
			for j:=0;j<16;j++{
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
					}
				}
				
			}
        }
		count++
				
		fmt.Printf("\rStatus: %d/%d", count, personsTotal)
		
		if count == personsTotal {
			break
		}
    }

    wg.Wait()

    // Close channels
    close(workChan0)
    close(workChan1)
    close(workChan2)
    close(workChan3)
    close(workChan4)
    close(workChan5)
    close(workChan6)
    close(workChan7)
    close(workChan8)
    close(workChan9)
    close(workChan10)
    close(workChan11)
    close(workChan12)
    close(workChan13)
    close(workChan14)
    close(workChan15)
    close(results)
    close(done)
    
   

    
}