package main

func main() {
	sid := "2018214830"
	password := "ccnuzmc"

	if err := GetCardInfo(sid, password); err != nil {
		panic(err)
	}
}
