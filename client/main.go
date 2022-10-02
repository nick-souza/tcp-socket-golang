package main

// Author: Nicolas Sá de Souza - 1072113016

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// Criando os canais para poder passar os dados entre as goroutines:
var (
	output    = make(chan string)
	input     = make(chan string)
	errorChan = make(chan error)
)

func main() {
	// Iniciando uma goroutine com a função para ler input do usuário:
	go readFromInput()

	// For loop com label para poder usar o 'continue':
RECONNECT:
	for {
		// Chamando a função connect e atribuindo o return para a variável:
		conn := connect()

		// Iniciando uma goroutine com a função para receber os dados do servidor:
		go readFromServer(conn)

		// For loop:
		for {
			// Usando select para poder esperar até que as cominuição dos canais seja finalizada:
			select {
			// Case para receber dados do servidor. Imprimindo na tela do cliente:
			case m := <-output:
				// fmt.Printf("--> Servidor: %s\n", strings.Trim(m, "\r\n"))
				fmt.Printf("--> Servidor: %s", m)

				// Case para ler o input do cliente e mandar para o servidor:
			case m := <-input:
				// Usando net.con.Write:
				_, err := conn.Write([]byte(m + "\n"))
				// Em caso de erro, fechar a conexão e continuar para o label RECONNECT:
				if err != nil {
					fmt.Println(err)
					conn.Close()
					continue RECONNECT
				}
				// Case para o canal de erro:
			case err := <-errorChan:
				// Caso o erro seja end of file (ex: o cliente digitou /sair), feche a conexão:
				if err == io.EOF {
					fmt.Println("\nFechando o cliente...")
					conn.Close()
					return
				}
				// Em caso de erro, fechar a conexão e continuar para o label RECONNECT:
				fmt.Println("Error:", err)
				conn.Close()
				continue RECONNECT
			}
		}
	}
}

// Função para conectar com o servidor. Retornando a conexão em si:
func connect() net.Conn {
	// Variáveis:
	var (
		// Para a conexão:
		conn net.Conn
		// Para eventuais erros:
		err error
	)

	// Loop infinito para ficar ouvindo até poder se conectar:
	for {
		fmt.Println("\nTentando se conectar ao servidor...")
		// Usando net.Dial para poder se conectar. Passando o tipo de conexão e nesse caso (local) apenas a porta:
		conn, err = net.Dial("tcp", ":6666")

		// Verificação de erro invertida. Caso o erro seja nil (ou seja, não houve nenhum erro) saia do loop:
		if err == nil {
			break
		}
		// Ecoando o erro:
		fmt.Println("\nErro: ", err)
		// Usando sleep de 1 segundo para tentar fazer a conexão de novo:
		time.Sleep(time.Second * 1)
	}

	// Saindo do loop caso a conexão seja aceita:
	fmt.Println("\nConexão bem sucedida")
	// Retornando a conexão:
	return conn
}

// Função com loop para ler o input do cliente
func readFromInput() {
	for {
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		input <- userInput
	}
}

// Função com loop para receber o que o servidor mandar:
func readFromServer(conn net.Conn) {
	for {
		serverInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			errorChan <- err
			return
		}
		output <- serverInput
	}
}
