package main

// Author: Nicolas Sá de Souza - 1072113016

type commandID int

// Variáveis para usar como IDs dos comandos:
const (
	// Usando iota para popular o commandID e ir incrementando:
	CMD_LIST commandID = iota
	CMD_MATRIX_MULTI_BY_ANOTHER
	CMD_MATRIX_MULTI_BY_NUM
	CMD_ADD_NUM_TO_MATRIX
	CMD_MATRIX_ADD_BY_ANOTHER
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
}
