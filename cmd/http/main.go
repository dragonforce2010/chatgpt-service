package main

func main() {
	r := InitializeRouter()
	r.Run("0.0.0.0:8080")
}
