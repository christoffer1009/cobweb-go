package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/christoffer1009/cobweb-go/node"
	"github.com/christoffer1009/cobweb-go/occurrence"
	"github.com/christoffer1009/cobweb-go/tree"
)

func main() {
	// Abrir o arquivo CSV
	file, err := os.Open("tabela.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}

	defer file.Close()

	// Criar um leitor de CSV
	reader := csv.NewReader(file)

  // Ler e descartar a primeira linha
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Erro ao ler a primeira linha:", err)
		return
	}


	// Ler as linhas do arquivo CSV
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler as linhas do arquivo:", err)
		return
	}

	var occurences []*occurrence.Occurrence

	for _, line := range lines {

		nucleus, _ := strconv.Atoi(line[1])
		tail, _ := strconv.Atoi(line[2])

		occurence := &occurrence.Occurrence{
			Color:   line[0],
			Nucleus: nucleus,
			Tail:    tail,
		}

		occurences = append(occurences, occurence)
	}

	root := node.NewNode(0)
	cobwebtree := tree.NewTree(root)

	for _, occ := range occurences {
		cobwebtree.Cobweb(cobwebtree.Root, occ)

	}

	fmt.Println("FIM")

}
