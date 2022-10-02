package main

// Author: Nicolas Sá de Souza - 1072113016

import (
	"fmt"
	"net"
)

func main() {
	// Inicializando um novo servidor:
	s := newServer()
	// Chamando uma goroutine para rodar a func run do server, que fica esperando comandos do cliente:
	go s.run()

	// Utilizando o net.Listen para poder esperar conexões em um thread separado, para poder aceitar mais conexões ao mesmo tempo;
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Printf("\n--> Erro ao iniciar servidor")
		fmt.Println("\n", err.Error())
	}

	// Impedir que a conexão seja encerrada:
	defer listener.Close()
	fmt.Println("--> Ouvindo na porta 6666")

	// Criando um loop infinito para cuidar daa conexões:
	for {
		// Aceitando a conexão:
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("\n--> Erro ao estabelecer conexão")
			fmt.Println("\n", err.Error())
			// Continue ao invés do return para continuar no for loop e continuar ouvindo;
			continue
		}
		defer conn.Close()

		// Usando a conexão para inicializar um novo cliente:
		c := s.newClient(conn)

		// Chamando uma goroutine para ficar lendo o input do cliente:
		go c.readInput()
	}
}
