package main

func main() {
	homeTheater := newHomeTheaterFacade()

	homeTheater.watchMovie("Raiders of the Lost Ark")
	homeTheater.endMovie()
}
