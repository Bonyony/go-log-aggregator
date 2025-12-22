package producer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)


func ProduceLog() {
	var wg sync.WaitGroup

	fmt.Println("Starting dummy log stream to Pub / Sub");
	entries, err := os.ReadDir("./data");
	if (err != nil) {
		log.Fatal(err);
	}

	var allLogs []string;
	for _, entry := range entries {
		if !entry.IsDir() {
			fullPath := filepath.Join("data", entry.Name())
			allLogs = append(allLogs, fullPath)
		}
	}

	errChan := make(chan error, len(allLogs))

	go func() {
		for err := range errChan {
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}()

	for _, logPath := range allLogs {
		wg.Add(1)
		go openAndScanLogs(logPath, &wg, errChan);
	}

	wg.Wait()
	close(errChan)
}

func openAndScanLogs(s string, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	log, err := os.Open(s);
	if (err != nil) {
		errChan <- fmt.Errorf("task %s encountered an error: %w", s, err)
		return
	}
	defer log.Close();

	scanner := bufio.NewScanner(log);
	for scanner.Scan() {
		// logic to pub / sub will go here
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        errChan <- fmt.Errorf("Scanner error in %s: %w", s, err)
    }
}