type data struct {
	// ...
}

func handleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// ...
	c := make(chan data)
	defer close(c)
	defer r.Body.Close()
	go saveDataFromChannel(c)
	c <- calculateStuff(getData(r.Body))
}

func getData(r io.Reader) data {
	// ...
	return data{}
}

func calculateStuff(d data) data {
	// ...
	return d
}

func saveDataFromChannel(c chan data) {
	for d := range c {
		saveToDatabase(d)
		// ...
	}
}

func saveToDatabase(d data) {
	// ...
}