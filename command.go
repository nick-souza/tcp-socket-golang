package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JOIN

	CMD_GRADES
	CMD_MATRIX
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
