package server

type evento struct {
	ID		string `json:"id" firestore:"-"`
	Nome	string `json:"nome" firestore:"nome"`
	Local	string `json:"local" firestore:"local"`
	Cidade	string `json:"cidade" firestore:"cidade"`
	Horario	string `json:"horario" firestore:"horario"`
	Dia		string `json:"dia" firestore:"dia"`
	Tipo	string `json:"tipo do role" firestore:"tipo_role"`
}

type responseError struct {
	Message	string `json:"message"`
}

type responseDelete struct {
	Message	string `json:"message"`
}