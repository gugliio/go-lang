package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandlerRoot)
	server.Handle("/api", HandlerHome)
	server.Listen()
}
