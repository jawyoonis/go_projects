package main

func main() {

	slice := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(slice); i++ {
		// println(slice[i])
	}

	for i, v := range slice { // v=value, i=index
		println(i, v)
	}

	jamalBio := map[string]string{"name": "jamal", "age": "20", "accupation": "software enginear", "nationality": "Somali"}
	for k, v := range jamalBio {
		println(k, v)
	}
}
