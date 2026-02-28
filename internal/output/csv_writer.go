package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"pix-consulta-chaves-worker/internal/domain"
)

func ConsumidorCSV(resultChan <-chan *model.Chave, done chan<- struct{}) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	writer.Write([]string{
		"Chave", "TpChave", "Ispb", "NrAgencia", "TpConta", "NrConta",
		"DtHrAberturaConta", "TpPessoa", "CpfCnpj", "Nome", "NomeFantasia",
		"DtHrCriacaoChave", "DtHrIncioPosseChave", "DtHrAberturaReivindicacao", "EndToEndId",
	})

	for info := range resultChan {
		err := writer.Write([]string{
			info.Chave,
			fmt.Sprintf("%d", info.TpChave),
			info.Ispb,
			info.NrAgencia,
			fmt.Sprintf("%d", info.TpConta),
			info.NrConta,
			info.DtHrAberturaConta,
			fmt.Sprintf("%d", info.TpPessoa),
			info.CpfCnpj,
			info.Nome,
			info.NomeFantasia,
			info.DtHrCriacaoChave,
			info.DtHrIncioPosseChave,
			info.DtHrAberturaReivindicacao,
			info.EndToEndId,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro escrevendo CSV: %v\n", err)
		}
	}
	done <- struct{}{}
}
