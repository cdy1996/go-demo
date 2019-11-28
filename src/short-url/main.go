package main

func main() {
	a := App{}
	a.Initialize(getEnv())
	a.Run("127.0.0.1:8000")

}
