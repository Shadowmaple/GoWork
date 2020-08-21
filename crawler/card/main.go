package main

func main() {
	sid := ""
	password := ""

	if err := GetCardInfo(sid, password); err != nil {
		panic(err)
	}
}
