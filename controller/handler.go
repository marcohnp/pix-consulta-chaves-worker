package controller

import (
	"bufio"
	"os"
	"pix-consulta-chaves-worker/service"
)

func StartProcessing() {
	scanner := bufio.NewScanner(os.Stdin)
	service.ProcessChaves(scanner)
}
