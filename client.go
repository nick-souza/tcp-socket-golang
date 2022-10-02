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
	// E um canal para troca dos comandos:
	commands chan<- command
}

// Func con loop parar poder ler as mensagens do cliente
func (c *client) readInput() {
READINPUT:
	for {
		c.msg("Para a lista de comandos digite /cmd\n")

		// NewReader:
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		// Dividindo entre "/" e o resto da mensagem:
		args := strings.Split(msg, " ")
		// Criando uma variável para segurar o comando do cliente:
		cmd := strings.TrimSpace(args[0])

		// Switch case para construir o comando certo:
		switch cmd {
		case "/cmd":
			c.commands <- command{
				id:     CMD_LIST,
				client: c,
				args:   args,
			}

		case "/m1":
			c.commands <- command{
				id:     CMD_MATRIX_MULTI_BY_NUM,
				client: c,
				args:   args,
			}
			break READINPUT

		case "/m2":
			c.commands <- command{
				id:     CMD_MATRIX_MULTI_BY_ANOTHER,
				client: c,
				args:   args,
			}
			break READINPUT

		case "/m3":
			c.commands <- command{
				id:     CMD_ADD_NUM_TO_MATRIX,
				client: c,
				args:   args,
			}
			break READINPUT

		case "/m4":
			c.commands <- command{
				id:     CMD_MATRIX_ADD_BY_ANOTHER,
				client: c,
				args:   args,
			}
			break READINPUT

		case "/sair":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
			break READINPUT

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
