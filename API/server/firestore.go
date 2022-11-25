package server

import (
	"log"
	"os"
	"context"
	"cloud.google.com/go/firestore"
)

func StartFireEnvp(){
	err := os.Setenv("FIRESTORE_EMULATOR_HOST", "firebase:8081"); if err != nil {
		log.Fatal(err)
	}
}

func StartFire() {
	FireStore.ctx = context.Background()
	var err error

	FireStore.client, err = firestore.NewClient(FireStore.ctx, "bexs-challenge-1ea1e"); if err != nil {
		log.Println(err)
	}
}