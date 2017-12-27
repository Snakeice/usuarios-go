package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/snakeice/usuarios/models"
	"strconv"
)

func Home(context echo.Context) error {
	var usuarios []models.Usuario

	if err := models.UsuarioModel.Find().All(&usuarios); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"menssagem": "Erro ao recupera dados.",
		})
	}

	data := map[string]interface{}{
		"titulo":   "Listagem de usuarios",
		"usuarios": usuarios,
	}
	return context.Render(http.StatusOK, "index.html", data)
}

func Add(c echo.Context) error {

	return c.Render(http.StatusOK, "add.html", nil)
}

func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")
	if nome != "" && email != "" {
		var usuario models.Usuario
		usuario.Email = email
		usuario.Nome = nome
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"menssagem": "Não foi possivel salvar registro!!",
			})
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"menssagem": "Os campos precisam ser informados!",
	})

}
func Deletar(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	resultado := models.UsuarioModel.Find("id=?", userID)

	if count, _ := resultado.Count(); count > 0 {
		if err := resultado.Delete(); err != nil {
			return c.JSON(http.StatusAccepted, map[string]string{
				"menssagem": "Não foi possivel deletar o usuario!",
			})

		}
		return c.JSON(http.StatusAccepted, map[string]string{
			"menssagem": "Usuário deletado com sucesso!",
		})
	}
	return c.JSON(http.StatusNotFound, map[string]string{
		"menssagem": "Usuário não encontrado!",
	})

}

func Atualizar(c echo.Context) error {
	uId, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("nome")
	email := c.FormValue("email")
	resultado := models.UsuarioModel.Find("id=?", uId)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"menssagem": "Usuário não existe!",
		})
	}

	user := models.Usuario{
		Nome:  nome,
		Email: email,
		Id:    uId,
	}

	if err := resultado.Update(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"menssagem": "Erro ao tentar atualizar o registro!",
		})
	}
	return c.JSON(http.StatusOK, user)
}

func Update(c echo.Context) error {
	uId, _ := strconv.Atoi(c.Param("id"))
	var usuario models.Usuario

	resultado := models.UsuarioModel.Find("id=?", uId)
	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"menssagem": "Usuário não existe!",
		})
	}

	if err := resultado.One(&usuario); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"menssagem": "Usuário não existe!",
		})
	}

	data := map[string]interface{}{
		"usuario": usuario,
	}
	return c.Render(http.StatusOK, "update.html", data)
}
