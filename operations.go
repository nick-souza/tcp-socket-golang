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
		matrixEmpty := buildMatrix("A", c)

		// Pegando o tamanha da matrix:
		var rows int = len(matrixEmpty)
		var columns int = len(matrixEmpty[0])

		// Chamando a função para popular a matriz:
		matrix := populateMatrix("A", matrixEmpty, c)

		// Variável para segurar o valor a ser multiplicado:
		var multiplyByInt int

		// Loop para o cliente inserir o valor:
	MULTIPLY_BY:
		for {
			c.msg("Insira o número que deseja multiplicar a matrix A:\n")
			var multiplyBy, _ = bufio.NewReader(c.conn).ReadString('\n')

			// Validando o input do cliente:
			var err error
			multiplyByInt, err = checkNumber(multiplyBy)
			if err != nil {
				c.err(fmt.Errorf("para fazer multiplicação precisa ser um número"))
				// Voltando para o início do loop:
				continue MULTIPLY_BY
			}

			// Saindo do loop caso não ocorra mais nenhum erro:
			break MULTIPLY_BY
		}

		// Criando a matriz resultado; Tendo o mesmo tamanho da matriz original:
		var result = make([][]int, rows)
		for i := 0; i < rows; i++ {
			result[i] = make([]int, columns)
		}

		// Fazendo o loop sobre a matriz para poder imprimir os valores:
		for i, row := range result {
			for j := range row {
				// Multiplicando cada elemento da matriz pelo valor do cliente:
				result[i][j] = matrix[i][j] * multiplyByInt
			}
		}

		// Ecoando o resultado para o cliente:
		c.msg(fmt.Sprintf("Resultado da matriz A multiplicada por %v é:\n", multiplyByInt))
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
		matrixAEmpty := buildMatrix("A", c)

		// Pegando o tamanha da matrix:
		var rowsA int = len(matrixAEmpty)
		var columnsA int = len(matrixAEmpty[0])

		matrixBEmpty := buildMatrix("B", c)

		// Pegando o tamanha da matrix:
		var rowsB int = len(matrixBEmpty)
		var columnsB int = len(matrixBEmpty[0])

		if columnsA != rowsB {
			c.err(fmt.Errorf("multiplicação não é possível, o número de colunas da matriz A é diferente no número de linhas da matriz B"))
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
		c.msg("Resultado da multiplicação da matriz A pela matriz B é:\n")
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
CALCULATE:
	for {
		// Chamando a função de criar matrizes:
		matrixEmpty := buildMatrix("A", c)

		// Pegando o tamanha da matrix:
		var rows int = len(matrixEmpty)
		var columns int = len(matrixEmpty[0])

		// Chamando a função para popular a matriz:
		matrix := populateMatrix("A", matrixEmpty, c)

		// Variável para segurar o valor a ser multiplicado:
		var addToInt int

		// Loop para o cliente inserir o valor:
	ADD:
		for {
			c.msg("Insira o número que deseja somar com os elementos da matriz A:\n")
			var addTo, _ = bufio.NewReader(c.conn).ReadString('\n')

			// Validando o input do cliente:
			var err error
			addToInt, err = checkNumber(addTo)
			if err != nil {
				c.err(fmt.Errorf("para somar precisa ser um número"))
				// Voltando para o início do loop:
				continue ADD
			}

			// Saindo do loop caso não ocorra mais nenhum erro:
			break ADD
		}

		// Criando a matriz resultado; Tendo o mesmo tamanho da matriz original:
		var result = make([][]int, rows)
		for i := 0; i < rows; i++ {
			result[i] = make([]int, columns)
		}

		// Fazendo o loop sobre a matriz para poder imprimir os valores:
		for i, row := range result {
			for j := range row {
				// Multiplicando cada elemento da matriz pelo valor do cliente:
				result[i][j] = matrix[i][j] + addToInt
			}
		}

		// Ecoando o resultado para o cliente:
		c.msg(fmt.Sprintf("Resultado da matriz A somada por %v é:\n", addToInt))
		// Ecoando o resultado para o servidor:
		fmt.Printf("--> Cliente: %v - resultado é:\n", c.conn.RemoteAddr())
		// Chamando a função para dar print:
		printMatrix(result, c)

		// Caso não ocorra mais erros, saindo do loop:
		break CALCULATE
	}

	// Quando terminar de calcular, voltar para o menu para o cliente escolher outra operação:
	go c.readInput()
}

func (s *server) addMatrixToAnother(c *client) {
	c.msg("Você escolheu somar uma matriz com outra.\n")

	var matrixAEmpty, matrixBEmpty [][]int
	// var matrixAError, matrixBError error
	var rowsA, colsA, rowsB, colsB int

BUILD_MATRICES:
	for {

		// Chamando a função de criar matrizes:
		matrixAEmpty = buildMatrix("A", c)

		// Pegando o tamanha da matrix:
		rowsA = len(matrixAEmpty)
		colsA = len(matrixAEmpty[0])

		// Chamando a função de criar matrizes:
		matrixBEmpty = buildMatrix("B", c)

		// Pegando o tamanha da matrix:
		rowsB = len(matrixBEmpty)
		colsB = len(matrixBEmpty[0])

		if colsA != colsB || rowsA != rowsB {
			c.err(fmt.Errorf("adição não é possível, as matrizes inseridas são de ordens diferentes"))
			// Voltar para o início:
			continue BUILD_MATRICES
		}

		// Chamando a função para popular as matrizes:
		matrixA := populateMatrix("A", matrixAEmpty, c)
		matrixB := populateMatrix("B", matrixBEmpty, c)

		// Criando a matriz resultado, com a mesma ordem das outras:
		var result = make([][]int, rowsA)
		for i := 0; i < rowsA; i++ {
			result[i] = make([]int, colsA)
		}

		// Fazendo o loop sobre a matriz para poder imprimir os valores:
		for i, row := range result {
			for j := range row {
				// Multiplicando cada elemento da matriz pelo valor do cliente:
				result[i][j] = matrixA[i][j] + matrixB[i][j]
			}
		}

		// Ecoando o resultado para o cliente:
		c.msg("Resultado da soma da matriz A com a matriz B é:\n")
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
		c.msg("Para multiplicar uma matriz por um número: /m1\n")
		time.Sleep(time.Millisecond * 10) // Usando sleep aqui pois as vezes ele não ecoava todas as mensagens;
		c.msg("Para multiplicar uma matriz por outra matriz: /m2\n")
		time.Sleep(time.Millisecond * 10)
		c.msg("Para somar uma matriz por um número: /m3\n")
		time.Sleep(time.Millisecond * 10)
		c.msg("Para somar uma matriz com outra matriz: /m4\n")
		time.Sleep(time.Millisecond * 10)
		c.msg("Para sair: /sair\n")
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

// Função que pega os inputs do usuário e cria uma matriz, retorna essa matriz e um potencial err:
func buildMatrix(name string, c *client) [][]int {
	for {
		var rowsString, columnsString string
		var rows, columns int

	GET_ROWS:
		for {
			// Número de linhas da matriz:
			c.msg(fmt.Sprintf("Coloque o número de linhas da matriz %s: \n", name))
			// Número de linhas da matriz:
			rowsString, _ = bufio.NewReader(c.conn).ReadString('\n')

			// Validando o input do cliente:
			var err error
			rows, err = checkNumber(rowsString)
			if err != nil {
				// Criando o erro para retornar:
				c.err(fmt.Errorf("número de linhas precisa ser um número"))
				continue GET_ROWS
			}
			break GET_ROWS
		}

	GET_COLS:
		for {
			c.msg(fmt.Sprintf("Coloque o número de colunas da matriz %s: \n", name))
			// Número do colunas da matriz:
			columnsString, _ = bufio.NewReader(c.conn).ReadString('\n')

			// Validando o input do cliente:
			var err error
			columns, err = checkNumber(columnsString)
			if err != nil {
				// Criando o erro para retornar:
				c.err(fmt.Errorf("número colunas precisa ser um número"))
				continue GET_COLS
			}
			break GET_COLS
		}

		// Criando uma matriz bidimensional de acordo com os valores do cliente:
		var matrix = make([][]int, rows)
		for i := 0; i < rows; i++ {
			matrix[i] = make([]int, columns)
		}

		// Retornando a matriz:
		return matrix
	}
}

// Função para popular uma matriz com inputs do cliente:
func populateMatrix(name string, matrix [][]int, c *client) [][]int {
	c.msg(fmt.Sprintf("Insira os valores da matriz %s: \n", name))

	// Para o cliente inserir os valores da matriz:
POPULATE:
	for {
		// Loop sobre os elementos da matriz para poder definir os valores:
		for k, r := range matrix {
			for l := range r {
				c.msg(fmt.Sprintf("Valores: matrix %s[%d][%d]: \n", name, k, l))
				var value, _ = bufio.NewReader(c.conn).ReadString('\n')

				// Validando o input do cliente:
				valueInt, err := checkNumber(value)
				if err != nil {
					c.err(fmt.Errorf("favor inserir apenas números"))
					// Voltando para o início do loop:
					continue POPULATE
				}
				// Atribuindo os valores:
				matrix[k][l] = valueInt
			}
		}
		// Saindo do loop caso não ocorra mais nenhum erro:
		break POPULATE
	}
	// Retornando a matriz com os elementos:
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
			c.msg(fmt.Sprint(matrix[i][j], "\t"))
		}
		// Para dividir as linhas:
		c.msg(fmt.Sprintln(""))
		fmt.Println()
	}
}

// Função para validar se o input é de fato um número. Retornando o número, e caso ocorra, um erro:
func checkNumber(str string) (int, error) {
	// Usando trim para limpar o input:
	str = strings.Trim(str, "\r\n")
	// Convertendo para int:
	num, err := strconv.Atoi(str)
	if err != nil {
		// Retornando erro caso ocorra:
		return 0, err
	}
	// Retornando o número e sem erros:
	return num, nil
}
