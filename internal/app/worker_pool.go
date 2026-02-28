package app

import (
	"bufio"
	"fmt"
	"os"
	"pix-consulta-chaves-worker/internal/client"
	"pix-consulta-chaves-worker/internal/config"
	"pix-consulta-chaves-worker/internal/domain"
	"pix-consulta-chaves-worker/internal/output"
	"sync"
)

func ProcessChaves(scanner *bufio.Scanner) {
	chavesChan := make(chan string)
	resultChan := make(chan *model.Chave)
	done := make(chan struct{})

	go util.ConsumidorCSV(resultChan, done)

	var wg sync.WaitGroup
	numWorkers := config.NumWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for chave := range chavesChan {
				if info, err := client.ConsultarChavePix(chave); err == nil {
					resultChan <- info
				} else {
					fmt.Fprintf(os.Stderr, "Erro ao consultar chave %s: %v\n", chave, err)
				}
			}
		}()
	}

	for scanner.Scan() {
		chave := scanner.Text()
		if chave != "" {
			chavesChan <- chave
		}
	}
	close(chavesChan)
	wg.Wait()
	close(resultChan)
	<-done
}
