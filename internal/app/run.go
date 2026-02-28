package app

import (
	"bufio"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	ProcessChaves(scanner)
}
