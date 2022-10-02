package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Struct para cuidar das propriedades do cliente:
type client struct {
	// A conexão do usuário:
	conn net.Conn

	// nick     string
	// room     *room

	// E um canal para troca dos comandos:
	commands chan<- command
}

// Func con loop parar poder ler as mensagens do cliente
func (c *client) readInput() {
	for {
		// NewReader:
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		// Removendo os "/" e os espaços da mensagem:
		// msg = strings.Trim(msg, "\r\n")
		// Dividindo entre "/" e o resto da mensagem:
		args := strings.Split(msg, " ")
		// Criando uma variável para segurar o comando do cliente:
		cmd := strings.TrimSpace(args[0])

		// Switch case para construir o comando certo:
		switch cmd {
		// case "/nick":
		// 	c.commands <- command{
		// 		id:     CMD_NICK,
		// 		client: c,
		// 		args:   args,
		// 	}
		// case "/join":
		// 	c.commands <- command{
		// 		id:     CMD_JOIN,
		// 		client: c,
		// 		args:   args,
		// 	}
		// case "/rooms":
		// 	c.commands <- command{
		// 		id:     CMD_ROOMS,
		// 		client: c,
		// 	}
		// case "/msg":
		// 	c.commands <- command{
		// 		id:     CMD_MSG,
		// 		client: c,
		// 		args:   args,
		// 	}

		case "/notas":
			c.commands <- command{
				id:     CMD_GRADES,
				client: c,
				args:   args,
				msg:    msg,
			}

		case "/matriz":
			c.commands <- command{
				id:     CMD_MATRIX,
				client: c,
				args:   args,
			}

		case "/sair":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
			}

		// Case não caia nos cases:
		default:
			c.err(fmt.Errorf("Comando não reconhecido: %s", cmd))
		}
	}
}

// Função para escrever potenciais erros para o cliente:
func (c *client) err(err error) {
	c.conn.Write([]byte("Erro: " + err.Error() + "\n"))
}

// Função para escrever mensagens:
func (c *client) msg(msg string) {
	// Usando o net Write, que aceita apenas byte:
	c.conn.Write([]byte(msg + "\n"))
}
