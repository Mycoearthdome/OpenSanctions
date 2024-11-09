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


func ProcessingPersons(personName string, CHUNK_SIZE int, facebookEntries []string, FacebookEntriesTotal int, results chan Result, wg *sync.WaitGroup){
    wg.Done()
    
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
    

    var resultSlice []string
        

    for i := 0; i < FacebookEntriesTotal; i += (CHUNK_SIZE * NB_CORES){ 
       
            
        chunk0 := facebookEntries[i:min(i+CHUNK_SIZE, len(facebookEntries))]
        chunk1 := facebookEntries[i + CHUNK_SIZE * 2:min(i+CHUNK_SIZE * 2, FacebookEntriesTotal)]
        chunk2 := facebookEntries[i + CHUNK_SIZE * 3:min(i+CHUNK_SIZE * 3, FacebookEntriesTotal)]
        chunk3 := facebookEntries[i + CHUNK_SIZE * 4:min(i+CHUNK_SIZE * 4, FacebookEntriesTotal)]
        chunk4 := facebookEntries[i + CHUNK_SIZE * 5:min(i+CHUNK_SIZE * 5, FacebookEntriesTotal)]
        chunk5 := facebookEntries[i + CHUNK_SIZE * 6:min(i+CHUNK_SIZE * 6, FacebookEntriesTotal)]
        chunk6 := facebookEntries[i + CHUNK_SIZE * 7:min(i+CHUNK_SIZE * 7, FacebookEntriesTotal)]
        chunk7 := facebookEntries[i + CHUNK_SIZE * 8:min(i+CHUNK_SIZE * 8, FacebookEntriesTotal)]
        chunk8 := facebookEntries[i + CHUNK_SIZE * 9:min(i+CHUNK_SIZE * 9, FacebookEntriesTotal)]
        chunk9 := facebookEntries[i + CHUNK_SIZE * 10:min(i+CHUNK_SIZE * 10, FacebookEntriesTotal)]
        chunk10 := facebookEntries[i + CHUNK_SIZE * 11:min(i+CHUNK_SIZE * 11, FacebookEntriesTotal)]
        chunk11 := facebookEntries[i + CHUNK_SIZE * 12:min(i+CHUNK_SIZE * 12, FacebookEntriesTotal)]
        chunk12 := facebookEntries[i + CHUNK_SIZE * 13:min(i+CHUNK_SIZE * 13, FacebookEntriesTotal)]
        chunk13 := facebookEntries[i + CHUNK_SIZE * 14:min(i+CHUNK_SIZE * 14, FacebookEntriesTotal)]
        chunk14 := facebookEntries[i + CHUNK_SIZE * 15:min(i+CHUNK_SIZE * 15, FacebookEntriesTotal)]
        chunk15 := facebookEntries[i + CHUNK_SIZE * 16:min(i+CHUNK_SIZE * 16, FacebookEntriesTotal)]

        wg.Add(16)

        go worker(workChan0, results, wg)
        go worker(workChan1, results, wg)
        go worker(workChan2, results, wg)
        go worker(workChan3, results, wg)
        go worker(workChan4, results, wg)
        go worker(workChan5, results, wg)
        go worker(workChan6, results, wg)
        go worker(workChan7, results, wg)
        go worker(workChan8, results, wg)
        go worker(workChan9, results, wg)
        go worker(workChan10, results, wg)
        go worker(workChan11, results, wg)
        go worker(workChan12, results, wg)
        go worker(workChan13, results, wg)
        go worker(workChan14, results, wg)
        go worker(workChan15, results, wg)

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
        for j:=0;j<NB_CORES;j++{
            select {
            case result := <-results:
                if result.Entries != nil{
                    for _, entry := range result.Entries{
                        resultSlice = append(resultSlice, entry)
                        //fmt.Println(entry)
                    }
                    //fmt.Println("\nWriting to file...", result.Person)
                    err := writeResults(resultSlice, OutFile)
                    if err != nil {
                        fmt.Println(err)
                    }
                    clear(resultSlice)
                }
            }
            
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

    //fmt.Println("Removing duplicates...Please wait!")

    //duplicateFree := removeDuplicates(personsNames)
    
    fmt.Println("Hunting...Please wait!")

    var wg sync.WaitGroup
    var personsTotal int64 = int64(len(personsNames)) //duplicateFree
    var FacebookEntriesTotal int = len(facebookEntries)
    
    results := make(chan Result, personsTotal)
    
    count := int64(0)


    CHUNK_SIZE := int(math.Floor(float64(FacebookEntriesTotal / NB_CORES))) //TODO:Leaves the last results behind (< NB_CORES) address that.
    FacebookEntriesTotal = CHUNK_SIZE * NB_CORES //Abandonning last results.

    fmt.Printf("CHUNK SIZE= %d TOTAL_ENTRIES= %d\n", CHUNK_SIZE, FacebookEntriesTotal)
    fmt.Printf("Processing %d Persons\n", personsTotal)
    // Prepare work to workers
    
    for _, personName := range personsNames {
        wg.Add(1)
        go ProcessingPersons(personName, CHUNK_SIZE, facebookEntries, FacebookEntriesTotal, results, &wg)

        count++
				
		fmt.Printf("\rStatus: %d/%d", count, personsTotal)
		
    }

    wg.Wait()

    close(results)
   
    //Clear the remaining results
    for result := range results{
        var resultSlice []string
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