package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Struct para cuidar das propriedades do servidor:
type server struct {
	// rooms    map[string]*room

	// Usando channel para poder mover os comandos do usuário entre as goroutines
	commands chan command
}

// Func para criar um novo servidor:
func newServer() *server {
	// Retornando um novo servidor:
	return &server{
		// rooms:    make(map[string]*room),

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
		// TODO: Trocar para os comandos certos:
		case CMD_MATRIX:
			go server.matrixMultiplyByAnother(command.client, command.args)
		case CMD_GRADES:
			server.grades(command.client, command.args, command.msg)
		case CMD_QUIT:
			server.quit(command.client)
		}
	}
}

// Função para ser chamada quando um cliente entra no servidor, iniciando um novo cliente:
func (s *server) newClient(conn net.Conn) *client {
	fmt.Printf("\n--> Cliente entrou no servidor: %s", conn.RemoteAddr().String())

	// Retornando um novo cliente:
	return &client{
		// Passando a conexão:
		conn: conn,
		// E o canal para os comandos:
		commands: s.commands,
	}
}

// func (s *server) nick(c *client, args []string) {
// 	if len(args) < 2 {
// 		c.msg("nick is required. usage: /nick NAME")
// 		return
// 	}

// 	c.nick = args[1]
// 	c.msg(fmt.Sprintf("all right, I will call you %s", c.nick))
// }

// func (s *server) join(c *client, args []string) {
// 	if len(args) < 2 {
// 		c.msg("room name is required. usage: /join ROOM_NAME")
// 		return
// 	}

// 	roomName := args[1]

// 	r, ok := s.rooms[roomName]
// 	if !ok {
// 		r = &room{
// 			name:    roomName,
// 			members: make(map[net.Addr]*client),
// 		}
// 		s.rooms[roomName] = r
// 	}
// 	r.members[c.conn.RemoteAddr()] = c

// 	s.quitCurrentRoom(c)
// 	c.room = r

// 	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))

// 	c.msg(fmt.Sprintf("welcome to %s", roomName))
// }

// func (s *server) listRooms(c *client) {
// 	var rooms []string
// 	for name := range s.rooms {
// 		rooms = append(rooms, name)
// 	}

// 	c.msg(fmt.Sprintf("available rooms: %s", strings.Join(rooms, ", ")))
// }

// func (s *server) msg(c *client, args []string) {
// 	if len(args) < 2 {
// 		c.msg("message is required, usage: /msg MSG")
// 		return
// 	}

// 	msg := strings.Join(args[1:], " ")
// 	c.room.broadcast(c, c.nick+": "+msg)
// }

// Função para calcular matrizes:
func (s *server) matrixMultiplyByAnother(c *client, args []string) {
	c.msg("Você escolheu multiplicar uma matriz por outra\n")

WAITINGS:
	for {
		var reader = bufio.NewReader(c.conn)
		c.msg("Coloque o número de linhas matriz A: \n")

		var rowsAString, _ = reader.ReadString('\n')

		rowsA, err := checkNumber(rowsAString)

		if err != nil {
			c.msg("Não é um número")
			c.err(err)
			continue WAITINGS
		}

		c.msg(fmt.Sprintf("Inserido: %v", rowsA))
	}
}

func checkNumber(num string) (int, error) {
	num = strings.Trim(num, "\r\n")
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}
	return numInt, nil
}

func (server *server) grades(client *client, args []string, message string) {
	msg := strings.Fields(message)

	// Caso o usuário não tenha inserido mais que uma nota:
	if len(msg) < 3 {
		client.msg("Você precisa especificar pelo menos duas notas: /notas nota1 nota2")
		return
	}

	client.msg("Você escolheu normalização de notas")

	// Chamando a função para calcular e salvando na variável:
	result, largest, err := normalizeGrades(msg[1:])

	if err != nil {
		client.msg("Erro ao calcular notas, inserir somente números")
		return
	}

	client.msg(fmt.Sprintf("Soma é: %v. E o maior numero é: %v", result, largest))
}

// Função para calcular a normalização das notas
func normalizeGrades(grades []string) (int, int, error) {
	// Variável para guardar o resultado:
	var result int
	var largest int = 0

	// Fazendo um loop sobre a array de string que contém as notas:
	for _, i := range grades {
		// Fazendo um cast para int:
		j, err := strconv.Atoi(i)

		if err != nil {
			return 0, 0, err
		}
		if j > largest {
			largest = j
		}

		result += j
	}

	return result, largest, nil
}

// Função para cliente sair do servidor:
func (s *server) quit(c *client) {
	fmt.Printf("\n--> Cliente saiu do servidor: %s", c.conn.RemoteAddr().String())

	// Fechando a conexão do cliente:
	c.conn.Close()
}
