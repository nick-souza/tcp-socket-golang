package main

type commandID int

const (
	CMD_LIST commandID = iota
	CMD_MATRIX_MULTI_BY_ANOTHER
	CMD_MATRIX_MULTI_BY_NUM
	CMD_MATRIX_ADD_BY_ANOTHER
	CMD_MATRIX_ADD_BY_NUM
	CMD_QUIT
)

// Struct para cuidar das propriedades de um comando:
type command struct {
	// Id para depois poder identificar cada um:
	id commandID
	// Variável para guardar o cliente que fez o comando:
	client *client
	// Array de string para guardar os comandos:
	args []string

	msg string
}
