package main

// Author: Nicolas Sá de Souza - 1072113016

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Função para multiplicar uma matriz por um único número:
func (s *server) multiplyMatrixByNumber(c *client) {
	c.msg("Você escolheu multiplicar uma matriz por um número.\n")

	// Label para poder continuar caso haja erro:
BUILD_MATRIX:
	for {
		// Chamando a função de criar matrizes:
		matrixEmpty, err := buildMatrix("A", c)
		if err != nil {
			// Informar o cliente do erro:
			c.err(err)
			// Em caso de erro, voltando para o início do loop:
			continue BUILD_MATRIX
		}

		// Pegando o tamanha da matrix:
		var rowsA int = len(matrixEmpty)
		var columnsA int = len(matrixEmpty[0])

		// Chamando a função para popular a matriz:
		matrixA := populateMatrix("A", matrixEmpty, c)

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

		// Fazendo o loop sobre a matriz para poder imprimir os valores:
		for i, row := range result {
			for j := range row {
				// Multiplicando cada elemento da matriz pelo valor do cliente:
				result[i][j] = matrixA[i][j] * multiplyByInt
			}
		}

		// Ecoando o resultado para o cliente:
		c.msg(fmt.Sprintf("Resultado da matriz A multiplicada por %v é:", multiplyByInt))
		// Ecoando o resultado para o servidor:
		fmt.Printf("--> Cliente: %v - resultado é:\n", c.conn.RemoteAddr())

		// Chamando a função para dar print:
		printMatrix(result, c)

		// Caso não ocorra mais erros, saindo do loop:
		break BUILD_MATRIX
	}

	// Quando terminar de calcular, voltar para o menu para o cliente escolher outra operação:
	go c.readInput()
}

// Função para calcular matrizes:
func (s *server) multiplyOneMatrixByAnother(c *client) {
	c.msg("Você escolheu multiplicar uma matriz por outra.\n")

BUILD_MATRICES:
	for {
		// Chamando a função de criar matrizes:
		matrixAEmpty, err := buildMatrix("A", c)
		if err != nil {
			// Informar o cliente do erro:
			c.err(err)
			// Em caso de erro, voltando para o início do loop:
			continue BUILD_MATRICES
		}

		// Pegando o tamanha da matrix:
		var rowsA int = len(matrixAEmpty)
		var columnsA int = len(matrixAEmpty[0])

		matrixBEmpty, err := buildMatrix("B", c)
		if err != nil {
			// Informar o cliente do erro:
			c.err(err)
			// Em caso de erro, voltando para o início do loop:
			continue BUILD_MATRICES
		}

		// Pegando o tamanha da matrix:
		var rowsB int = len(matrixBEmpty)
		var columnsB int = len(matrixBEmpty[0])

		if columnsA != rowsB {
			c.err(fmt.Errorf("Multiplicação não é possível, o número de colunas da matrix A é diferente no número de linhas da matrix B.\n"))
			continue BUILD_MATRICES
		}

		// Chamando a função para popular as matrizes:
		matrixA := populateMatrix("A", matrixAEmpty, c)
		matrixB := populateMatrix("B", matrixBEmpty, c)

		// Criando a matriz resultado, com o número do linhas da matriz A, e o número de colunas da matriz B:
		var result = make([][]int, rowsA)
		for i := 0; i < rowsA; i++ {
			result[i] = make([]int, columnsB)
		}

		// Variável para ser somada:
		var total int
		total = 0

		// Loop entre as duas matrizes:
		for i := 0; i < rowsA; i++ {
			for j := 0; j < columnsB; j++ {
				for k := 0; k < rowsB; k++ {
					// Fazendo a multiplicação:
					total = total + matrixA[i][k]*matrixB[k][j]
				}
				// Atribuindo o valor para o elemento certo:
				result[i][j] = total
				// Voltando a zero para o próximo cálculo:
				total = 0
			}
		}

		// Ecoando o resultado para o cliente:
		c.msg(fmt.Sprintf("Resultado da multiplicação da matriz A pela matriz B é:"))
		// Ecoando o resultado para o servidor:
		fmt.Printf("--> Cliente: %v - resultado é:\n", c.conn.RemoteAddr())
		// Chamando a função para dar print:
		printMatrix(result, c)

		// Caso não ocorra mais erros, saindo do loop:
		break BUILD_MATRICES
	}

	// Quando terminar de calcular, voltar para o menu para o cliente escolher outra operação:
	go c.readInput()
}

func (s *server) addNumToMatrix(c *client) {
	c.msg("Você escolheu somar uma matriz por um número.\n")

	// Label para poder continuar caso haja erro:
BUILD_MATRIX:
	for {
		// Chamando a função de criar matrizes:
		matrixEmpty, err := buildMatrix("A", c)
		if err != nil {
			// Informar o cliente do erro:
			c.err(err)
			// Em caso de erro, voltando para o início do loop:
			continue BUILD_MATRIX
		}

		// Pegando o tamanha da matrix:
		var rowsA int = len(matrixEmpty)
		var columnsA int = len(matrixEmpty[0])

		// Chamando a função para popular a matriz:
		matrixA := populateMatrix("A", matrixEmpty, c)

		// Variável para segurar o valor a ser multiplicado:
		var addToInt int

		// Loop para o cliente inserir o valor:
	ADD:
		for {
			c.msg("Insira o número que deseja somar na matrix A:\n")
			var addTo, _ = bufio.NewReader(c.conn).ReadString('\n')

			// Validando o input do cliente:
			addToInt, err = checkNumber(addTo)
			if err != nil {
				c.err(fmt.Errorf("Erro: para fazer multiplicação precisa ser um número\n"))
				// Voltando para o início do loop:
				continue ADD
			}

			// Saindo do loop caso não ocorra mais nenhum erro:
			break ADD
		}

		// Criando a matriz resultado; Tendo o mesmo tamanho da matriz original:
		var result = make([][]int, rowsA)
		for i := 0; i < rowsA; i++ {
			result[i] = make([]int, columnsA)
		}

		// Fazendo o loop sobre a matriz para poder imprimir os valores:
		for i, row := range result {
			for j := range row {
				// Multiplicando cada elemento da matriz pelo valor do cliente:
				result[i][j] = matrixA[i][j] + addToInt
			}
		}

		// Ecoando o resultado para o cliente:
		c.msg(fmt.Sprintf("Resultado da matriz A somada por %v é:", addToInt))
		// Ecoando o resultado para o servidor:
		fmt.Printf("--> Cliente: %v - resultado é:\n", c.conn.RemoteAddr())
		// Chamando a função para dar print:
		printMatrix(result, c)

		// Caso não ocorra mais erros, saindo do loop:
		break BUILD_MATRIX
	}

	// Quando terminar de calcular, voltar para o menu para o cliente escolher outra operação:
	go c.readInput()
}

func (s *server) addMatrixToAnother(c *client) {
	c.msg("Você escolheu somar uma matriz com outra.\n")

BUILD_MATRICES:
	for {
		// Chamando a função de criar matrizes:
		matrixAEmpty, err := buildMatrix("A", c)
		if err != nil {
			// Informar o cliente do erro:
			c.err(err)
			// Em caso de erro, voltando para o início do loop:
			continue BUILD_MATRICES
		}

		// Pegando o tamanha da matrix:
		var rowsA int = len(matrixAEmpty)
		var columnsA int = len(matrixAEmpty[0])

		matrixBEmpty, err := buildMatrix("B", c)
		if err != nil {
			// Informar o cliente do erro:
			c.err(err)
			// Em caso de erro, voltando para o início do loop:
			continue BUILD_MATRICES
		}

		// Pegando o tamanha da matrix:
		var rowsB int = len(matrixBEmpty)
		var columnsB int = len(matrixBEmpty[0])

		if columnsA != columnsB || rowsA != rowsB {
			c.err(fmt.Errorf("Adição não é possível, as matrizes inseridas são de ordens diferentes.\n"))
			continue BUILD_MATRICES
		}

		// Chamando a função para popular as matrizes:
		matrixA := populateMatrix("A", matrixAEmpty, c)
		matrixB := populateMatrix("B", matrixBEmpty, c)

		// Criando a matriz resultado, com a mesma ordem das outras:
		var result = make([][]int, rowsA)
		for i := 0; i < rowsA; i++ {
			result[i] = make([]int, columnsA)
		}

		// Fazendo o loop sobre a matriz para poder imprimir os valores:
		for i, row := range result {
			for j := range row {
				// Multiplicando cada elemento da matriz pelo valor do cliente:
				result[i][j] = matrixA[i][j] + matrixB[i][j]
			}
		}

		// Ecoando o resultado para o cliente:
		c.msg(fmt.Sprintf("Resultado da soma da matriz A com a matriz B é:"))
		// Ecoando o resultado para o servidor:
		fmt.Printf("--> Cliente: %v - resultado é:\n", c.conn.RemoteAddr())

		// Chamando a função para dar print:
		printMatrix(result, c)

		// Caso não ocorra mais erros, saindo do loop:
		break BUILD_MATRICES
	}

	// Quando terminar de calcular, voltar para o menu para o cliente escolher outra operação:
	go c.readInput()
}

// Função para imprimir os comandos disponíveis:
func (s *server) showList(c *client) {
	for {
		c.msg(fmt.Sprint("Para multiplicar uma matriz por um número: /m1"))
		time.Sleep(time.Millisecond * 10) // Usando sleep aqui pois as vezes ele não ecoava todas as mensagens;
		c.msg(fmt.Sprint("Para multiplicar uma matriz por outra matriz: /m2"))
		time.Sleep(time.Millisecond * 10)
		c.msg(fmt.Sprint("Para somar uma matriz por um número: /m3"))
		time.Sleep(time.Millisecond * 10)
		c.msg(fmt.Sprint("Para somar uma matriz com outra matriz: /m4"))
		time.Sleep(time.Millisecond * 10)
		c.msg(fmt.Sprint("Para sair: /sair"))
		time.Sleep(time.Millisecond * 10)

		break
	}
}

// Função para cliente sair do servidor:
func (s *server) quit(c *client) {
	fmt.Printf("\n--> Cliente saiu do servidor: %s", c.conn.RemoteAddr().String())
	// Fechando a conexão do cliente:
	c.conn.Close()
}

//----- Helpers -----\\

// Função que pega os inputs do usuário e cria uma matriz. Retornando essa matriz, número de linhas e colunas e um potencial err:
func buildMatrix(name string, c *client) ([][]int, error) {
	// Número de linhas da matriz:
	c.msg(fmt.Sprintf("Coloque o número de linhas da matriz %s: \n", name))
	// Número de linhas da matriz:
	var rowsAString, _ = bufio.NewReader(c.conn).ReadString('\n')

	// Validando o input do cliente:
	rowsA, err := checkNumber(rowsAString)
	if err != nil {
		// Criando o erro para retornar:
		return nil, fmt.Errorf("Número de linhas precisa ser um número\n")
	}

	c.msg(fmt.Sprintf("Coloque o número de colunas da matriz %s: \n", name))
	// Número do colunas da matriz:
	var columnsAString, _ = bufio.NewReader(c.conn).ReadString('\n')

	// Validando o input do cliente:
	columnsA, err := checkNumber(columnsAString)
	if err != nil {
		// Criando o erro para retornar:
		return nil, fmt.Errorf("Número colunas precisa ser um número\n")
	}

	// Criando uma matriz bidimensional de acordo com os valores do cliente:
	var matrixA = make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		matrixA[i] = make([]int, columnsA)
	}

	// Retornando a matriz, número de linhas e colunas:
	return matrixA, nil
}

// Função para popular uma matriz com inputs do cliente:
func populateMatrix(name string, matrix [][]int, c *client) [][]int {
	// Loop para o cliente inserir os valores da matriz:
MATRIX_A_ELEMENTS:
	for {
		c.msg(fmt.Sprintf("Insira os valores da matrix %s: \n", name))

		// Loop sobre os elementos da matriz para poder definir os valores:
		for k, r := range matrix {
			for l := range r {
				c.msg(fmt.Sprintf("Valores: matrix %s[%d][%d]: ", name, k, l))
				var value, _ = bufio.NewReader(c.conn).ReadString('\n')

				// Validando o input do cliente:
				valueInt, err := checkNumber(value)
				if err != nil {
					c.err(fmt.Errorf("Erro: favor inserir apenas números\n"))
					// Voltando para o início do loop:
					continue MATRIX_A_ELEMENTS
				}

				// Atribuindo os valores:
				matrix[k][l] = valueInt
			}
		}

		// Saindo do loop caso não ocorra mais nenhum erro:
		break MATRIX_A_ELEMENTS
	}

	return matrix
}

func printMatrix(matrix [][]int, c *client) {
	// Fazendo loop para dar print no resultado:
	for i, row := range matrix {
		for j := range row {
			// Ecoando para o servidor:
			fmt.Print(matrix[i][j], "\t")
			// Ecoando para o cliente:
			time.Sleep(time.Millisecond * 10) // Usando sleep aqui pois as vezes ele não ecoava todas as mensagens;
			c.msg(fmt.Sprintf("Resultado[%d][%d]: %d\n", i, j, matrix[i][j]))
		}
		fmt.Println()
	}
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
