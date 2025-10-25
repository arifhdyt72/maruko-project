package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Maruko 🎵</title>
			<style>
				body {
					background-color: #ffeef8;
					display: flex;
					flex-direction: column;
					align-items: center;
					justify-content: center;
					height: 100vh;
					margin: 0;
					font-family: 'Comic Sans MS', cursive;
				}
				img {
					width: 300px;
					border-radius: 20px;
					box-shadow: 0 0 20px rgba(255, 105, 180, 0.6);
				}
				h1 {
					color: #e91e63;
					text-shadow: 1px 1px #fff;
				}
				audio {
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<h1>🎶 Maruko Style 🎶</h1>
			<img src="/static/maruko.gif" alt="Chibi Maruko Chan">
			<audio autoplay controls>
				<source src="/static/maruko.mp3" type="audio/mpeg">
				Your browser does not support the audio element.
			</audio>
		</body>
		</html>`
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"status":  "ok",
			"message": "server is running",
		}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("Server running on http://localhost:5544 🎵")
	http.ListenAndServe("0.0.0.0:5544", nil)
}
