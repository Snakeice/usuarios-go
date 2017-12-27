package models

import "github.com/snakeice/usuarios/lib"

type Usuario struct {
	Id int `db:"id" json:"id"`
	Nome string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

var UsuarioModel = lib.Conn.Collection("usuarios")

