package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseAuth adalah instance Firebase Auth global yang bisa diakses paket lain
var FirebaseAuth *auth.Client

// InitFirebase melakukan inisialisasi Firebase App dengan service account credentials
func InitFirebase() {
	ctx := context.Background()

	// 1. Ambil path path credentials dari environment variable
	credPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credPath == "" {
		log.Fatal("FIREBASE_CREDENTIALS_PATH tidak ditemukan di environment variables")
	}

	// 2. Inisialisasi opsi dengan file service account
	opt := option.WithCredentialsFile(credPath)

	// 3. Inisialisasi Firebase App
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Gagal inisialisasi Firebase App: %v", err)
	}

	// 4. Dapatkan Firebase Auth client dan masukkan ke variabel global
	FirebaseAuth, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("Gagal mendapatkan Firebase Auth client: %v", err)
	}

	log.Println("✅ Firebase Admin SDK berhasil diinisialisasi")
}