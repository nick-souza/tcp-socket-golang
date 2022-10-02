package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func buildMatrix(name string, c *client) ([][]int, int, int, error) {
	// Número de linhas da matriz:
	c.msg(fmt.Sprintf("Coloque o número de linhas da matriz %s: \n", name))
	var rowsAString, _ = bufio.NewReader(c.conn).ReadString('\n')

	// Validando o input do cliente:
	rowsA, err := checkNumber(rowsAString)
	if err != nil {
		// c.err(fmt.Errorf("Erro: linhas precisa ser um número\n"))
		// // Voltando para o início do loop:
		// continue BUILD_MATRIX

		return nil, 0, 0, fmt.Errorf("Erro: linhas precisa ser um número\n")
	}

	c.msg(fmt.Sprintf("Coloque o número de colunas da matriz %s: \n", name))
	var columnsAString, _ = bufio.NewReader(c.conn).ReadString('\n')

	columnsA, err := checkNumber(columnsAString)
	if err != nil {
		// c.err(fmt.Errorf("Erro: colunas precisa ser um número\n"))
		// // Voltando para o início do loop:
		// continue BUILD_MATRIX

		return nil, 0, 0, fmt.Errorf("Erro: linhas precisa ser um número\n")
	}

	var matrixA = make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		matrixA[i] = make([]int, columnsA)
	}

	return matrixA, rowsA, columnsA, nil
}

// Função para multiplicar uma matriz por um único número:
func (s *server) multiplyMatrixByNumber(c *client) {
	c.msg("Você escolheu multiplicar uma matriz por um número.\n")

	// Label para poder continuar caso haja erro:
BUILD_MATRIX:
	for {
		// Número de linhas da matriz:
		// c.msg("Coloque o número de linhas matriz A: \n")
		// var rowsAString, _ = bufio.NewReader(c.conn).ReadString('\n')

		// // Validando o input do cliente:
		// rowsA, err := checkNumber(rowsAString)
		// if err != nil {
		// 	c.err(fmt.Errorf("Erro: linhas precisa ser um número\n"))
		// 	// Voltando para o início do loop:
		// 	continue BUILD_MATRIX
		// }

		// // Número do colunas da matriz:
		// c.msg("Coloque o número de colunas matriz A: \n")
		// var columnsAString, _ = bufio.NewReader(c.conn).ReadString('\n')

		// // Validando o input do cliente:
		// columnsA, err := checkNumber(columnsAString)
		// if err != nil {
		// 	c.err(fmt.Errorf("Erro: colunas precisa ser um número\n"))
		// 	// Voltando para o início do loop:
		// 	continue BUILD_MATRIX
		// }

		// // Criando uma matriz bidimensional de acordo com os valores do cliente:
		// var matrixA = make([][]int, rowsA)
		// for i := 0; i < rowsA; i++ {
		// 	matrixA[i] = make([]int, columnsA)
		// }

		matrixA, rowsA, columnsA, err := buildMatrix("A", c)
		if err != nil {
			c.err(err)
			continue BUILD_MATRIX
		}

		// Loop para o cliente inserir os valores da matriz:
	MATRIX_A_ELEMENTS:
		for {
			c.msg("Insira os valores da matrix A: \n")

			// Loop sobre os elementos da matriz para poder definir os valores:
			for k, r := range matrixA {
				for l := range r {
					c.msg(fmt.Sprintf("Valores: matrixA[%d][%d]: ", k, l))
					var value, _ = bufio.NewReader(c.conn).ReadString('\n')

					// Validando o input do cliente:
					valueInt, err := checkNumber(value)
					if err != nil {
						c.err(fmt.Errorf("Erro: favor inserir apenas números\n"))
						// Voltando para o início do loop:
						continue MATRIX_A_ELEMENTS
					}

					// Atribuindo os valores:
					matrixA[k][l] = valueInt
				}
			}

			// Saindo do loop caso não ocorra mais nenhum erro:
			break MATRIX_A_ELEMENTS
		}

		// Variável para segurar o valor a ser multiplicado:
		var multiplyByInt int

		// Loop para o cliente inserir o valor:
	MULTIPLY_BY:
		for {
			c.msg("Insira o número que deseja multiplicar a matrix A:\n")
			var multiplyBy, _ = bufio.NewReader(c.conn).ReadString('\n')

			// Validando o input do cliente:
			multiplyByInt, err = checkNumber(multiplyBy)
			if err != nil {
				c.err(fmt.Errorf("Erro: para fazer multiplicação precisa ser um número\n"))
				// Voltando para o início do loop:
				continue MULTIPLY_BY
			}

			// Saindo do loop caso não ocorra mais nenhum erro:
			break MULTIPLY_BY
		}

		// Criando a matriz resultado; Tendo o mesmo tamanho da matriz original:
		var result = make([][]int, rowsA)
		for i := 0; i < rowsA; i++ {
			result[i] = make([]int, columnsA)
		}

		// Ecoando o resultado para o cliente:
		c.msg(fmt.Sprintf("Resultado da matriz A multiplicada por %v é:", multiplyByInt))
		// Ecoando o resultado para o servidor:
		fmt.Printf("--> Cliente: %v - resultado é:\n", c.conn.RemoteAddr())

		// Fazendo o loop sobre a matriz para poder imprimir os valores:
		for i, row := range result {
			for j := range row {
				// Multiplicando cada elemento da matriz pelo valor do cliente:
				result[i][j] = matrixA[i][j] * multiplyByInt

				//c.msg(fmt.Print("%d \t", matrixA[i][j]))
				//c.msg(fmt.Sprintf("%d ", matrixA[i][j]))
				//fmt.Print(matrixA[i][j], "\t")

				// Ecoando para o servido:
				fmt.Print(result[i][j], "\t")
				// Ecoando para o cliente:
				c.msg(fmt.Sprintf("Resultado[%d][%d]: %d\n", i, j, result[i][j]))
			}
			fmt.Println()
		}

		// Caso não ocorra mais erros, saindo do loop:
		break BUILD_MATRIX
	}

	// Quando terminar de calcular, voltar para o menu para o cliente escolher outra operação:
	go c.readInput()
}

// Função para calcular matrizes:
func (s *server) multiplyOneMatrixByAnother(c *client, args []string) {
	c.msg("Você escolheu multiplicar uma matriz por outra.\n")

BUILD_MATRICES:
	for {
		c.msg("Coloque o número de linhas matriz A: \n")

		var rowsAString, _ = bufio.NewReader(c.conn).ReadString('\n')

		rowsA, err := checkNumber(rowsAString)
		if err != nil {
			c.msg("Erro: favor inserir apenas números")
			continue BUILD_MATRICES
		}

		c.msg(fmt.Sprintf("Inserido: %v", rowsA))
	}
}

// Função para cliente sair do servidor:
func (s *server) quit(c *client) {
	fmt.Printf("\n--> Cliente saiu do servidor: %s", c.conn.RemoteAddr().String())
	// Fechando a conexão do cliente:
	c.conn.Close()
}

// Função para validar se o input é de fato um número. Retornando o número, e caso ocorra, um erro:
func checkNumber(num string) (int, error) {
	// Usando trim para limpar o input:
	num = strings.Trim(num, "\r\n")
	// Convertendo para int:
	numInt, err := strconv.Atoi(num)
	if err != nil {
		// Retornando erro caso ocorra:
		return 0, err
	}
	// Retornando o número e sem erros:
	return numInt, nil
}
