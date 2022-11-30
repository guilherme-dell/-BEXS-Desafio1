package server

type evento struct {
	ID		string `json:"id" firestore:"-"`
	Nome	string `json:"nome" firestore:"nome" binding:"required,min=5,max=50"`
	Local	string `json:"local" firestore:"local" binding:"required,min=10,max=80"`
	Cidade	string `json:"cidade" firestore:"cidade" binding:"required,min=10,max=50"`
	Horario	string `json:"horario" firestore:"horario" binding:"required,min=5,max=15"`
	Dia		string `json:"dia" firestore:"dia" binding:"required,datetime=02-01-2006"`
	Tipo	string `json:"tipo do role" firestore:"tipo_role" binding:"required,min=5,max=40"`
}

type responseError struct {
	Message	string `json:"message"`
}

type responseDelete struct {
	Message	string `json:"message"`
}