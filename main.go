package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// HTML content
		html := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Sample Go App</title>
				<style>
					body {
						margin: 0;
						padding: 0;
						font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
						background: #f5f5f5;
						display: flex;
						align-items: center;
						justify-content: center;
						height: 100vh;
					}
					.container {
						background: white;
						padding: 40px;
						border-radius: 12px;
						box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
						text-align: center;
					}
					h1 {
						color: #333;
						margin-bottom: 10px;
					}
					p {
						color: #666;
						font-size: 16px;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Welcome to the Sample Go App</h1>
					<p>This is a simple static site served using Go.</p>
				</div>
			</body>
			</html>`)

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
	})

	port := ":8000"
	fmt.Println("Server is running at : " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
