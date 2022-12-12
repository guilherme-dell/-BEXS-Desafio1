package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getRoleID (r *gin.Context) string {
	id, err := r.Params.Get("roleID"); if err != true {
		log.Panic(err)
	}
	return id
}

func pegarRole(r *gin.Context){
	roleID := getRoleID(r)
	collection := FireStore.client.Collection("Roles")
	doc, err := collection.Doc(roleID).Get(r.Request.Context())
	if err != nil {
		var response responseError
		response.Message = "Role não encontrado :("
		r.JSON(http.StatusNotFound, response)
		return
	}
	var response evento
	if err := doc.DataTo(&response); err != nil {
		r.String(http.StatusInternalServerError, err.Error())
		return
	}

	response.ID = doc.Ref.ID
	r.JSON(http.StatusOK, response)
	return
}

func todosRoles(r *gin.Context){
	var roles evento
	collection := FireStore.client.Collection("Roles").Documents(r)
	docs ,_ := collection.GetAll()
	for _, doc := range docs {
		doc.DataTo(&roles)
		roles.ID = doc.Ref.ID
		r.JSON(http.StatusOK, roles)
	}
	return
}

func marcarRole(r *gin.Context){
	var role evento
	if err := r.ShouldBindJSON(&role); err != nil {
		r.String(http.StatusBadRequest, err.Error())
		return
	}

	role.ID = uuid.NewString()
	collection := FireStore.client.Collection("Roles")
	_, err := collection.Doc(role.ID).Set(r, role)
	if err != nil {
		r.String(http.StatusInternalServerError, err.Error())
		return
	}
	r.JSON(http.StatusCreated, role)
}

func alteraRole(r *gin.Context){
	roleID := getRoleID(r)
	collection := FireStore.client.Collection("Roles")

	doc, err := collection.Doc(roleID).Get(r.Request.Context())
	if err != nil {
		var response responseError
		response.Message = "Role não encontrado :("
		r.JSON(http.StatusNotFound, response)
		return
	}
	var roleAtualizado evento
	if err := r.ShouldBindJSON(&roleAtualizado); err != nil {
		r.String(http.StatusBadRequest, err.Error())
		return
	}
	roleAtualizado.ID = doc.Ref.ID
	collection.Doc(doc.Ref.ID).Set(r, roleAtualizado)
	r.JSON(http.StatusOK, roleAtualizado)
}

func apagarRole(r *gin.Context){
	roleID := getRoleID(r)
	collection := FireStore.client.Collection("Roles")

	_, err := collection.Doc(roleID).Delete(r); if err != nil{
		var response responseError
		response.Message = "Role não encontrado :("
		r.JSON(http.StatusNotFound, response)
		return
	}
	var response responseDelete
	response.Message = "Role apagado"
	r.JSON(http.StatusOK, response)
}

func rolesHoje(r *gin.Context){
	var roles evento
	hoje := time.Now()
	collection := FireStore.client.Collection("Roles").Documents(r)
	docs ,_ := collection.GetAll()
	for _, doc := range docs {
		doc.DataTo(&roles)
		if (strings.Compare(roles.Dia, hoje.Format("02-01-2006")) == 0){
			roles.ID = doc.Ref.ID
			r.JSON(http.StatusOK, roles)
		}
	}
	return
}