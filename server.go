package main

import (
	"fmt"
	"net"
)

// Struct para cuidar das propriedades do servidor:
type server struct {
	// Usando channel para poder mover os comandos do usuário entre as goroutines
	commands chan command
}

// Func para criar um novo servidor:
func newServer() *server {
	// Retornando um novo servidor:
	return &server{
		// Criando um channel, para poder comunicar entre as goroutines, nesse caso para poder passar os comandos:
		commands: make(chan command),
	}
}

// Func run, para poder entender os comandos e chamar as funções apropriadas
func (server *server) run() {
	// Usando for e range para poder fazer um loop pelo channel de comandos:
	for command := range server.commands {
		// Switch case no ID, chamando as funções adequadas:
		switch command.id {
		case CMD_MATRIX_MULTI_BY_ANOTHER:
			go server.multiplyOneMatrixByAnother(command.client, command.args)
		case CMD_MATRIX_MULTI_BY_NUM:
			go server.multiplyMatrixByNumber(command.client)
		case CMD_MATRIX_ADD_BY_ANOTHER:
			// go server.multiplyMatrixByNumber(command.client)
		case CMD_MATRIX_ADD_BY_NUM:
			// go server.multiplyMatrixByNumber(command.client)
		case CMD_QUIT:
			server.quit(command.client)
		}
	}
}

// Função para ser chamada quando um cliente entra no servidor, iniciando um novo cliente:
func (s *server) newClient(conn net.Conn) *client {
	fmt.Printf("\n--> Cliente entrou no servidor: %s\n", conn.RemoteAddr().String())
	// Retornando um novo cliente:
	return &client{
		// Passando a conexão:
		conn: conn,
		// E o canal para os comandos:
		commands: s.commands,
	}
}
