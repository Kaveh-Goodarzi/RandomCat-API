package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Cat struct {
	ID          string   `json:"id"`
	ImageURL    string   `json:"image_url"`
	Breed       string   `json:"breed,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Width       int      `json:"width,omitempty"`
	Height      int      `json:"height,omitempty"`
}

var catBreeds = []string{
	"Siamese", "Persian", "Maine Coon", "Bengal", "Sphynx",
	"British Shorthair", "Scottish Fold", "Ragdoll", "Abyssinian",
	"Russian Blue", "Burmese", "Norwegian Forest", "Turkish Angora",
}

var catDescriptions= []string{
	"A fluffy and adorable feline friend.",
	"Master of the art of napping.",
	"Professional mouse hunter and cuddle expert.",
	"Has an attitude but loves treats.",
	"Loves cardboard boxes more than expensive beds.",
	"Will judge you silently from across the room.",
	"Purring machine and lap warmer.",
	"Expert in knocking things off tables.",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	http.HandleFunc("/random-cat", randomCatHandler)
	http.HandleFunc("/health", healthHandler)

	port := ":8080"
	fmt.Println(strings.Repeat("🐱", 30))
	fmt.Println("Random Cat API Server")
	fmt.Println(strings.Repeat("🐱", 30))
	fmt.Printf("Server running on http://localhost%s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET /           → frontend")
	fmt.Println("  GET /random-cat → Random cat data")
	fmt.Println("  GET /health     → Health check")
	fmt.Println(strings.Repeat("🐱", 30))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}


func randomCatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	cat := getRandomCat()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cat)
}

func getRandomCat() Cat {
	imageURL := getRandomCatImage()
	
	return Cat{
		ID: fmt.Sprintf("cat_%d", rand.Intn(10000)),
		ImageURL: imageURL,
		Breed: catBreeds[rand.Intn(len(catBreeds))],
		Description: catDescriptions[rand.Intn(len(catDescriptions))],
		Tags: generateRandomTags(),
		Width: 400 + rand.Intn(400),
		Height: 400 + rand.Intn(400),
	}
}

func getRandomCatImage() string {
	apis := []string {
		"https://cataas.com/cat?json=true",
		"https://api.thecatapi.com/v1/images/search",
	}
	apiURL := apis[rand.Intn(len(apis))]

	resp, err := http.Get(apiURL)

	if err != nil {
		// fallback to placeholder
		return "https://cataas.com/cat"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if apiURL == "https://cataas.com/cat?json=true" {
		var data map[string]interface{}
		if json.Unmarshal(body, &data) == nil {
			if url, ok := data["url"].(string); ok {
				return "https://cataas.com" + url
			}
		}
	} else if apiURL == "https://api.thecatapi.com/v1/images/search" {
		var data []map[string]interface{}
		if json.Unmarshal(body, &data) == nil && len(data) > 0 {
			if url, ok := data[0]["url"].(string); ok {
				return url
			}
		}
	}

	// Fallbach
	return "https://cataas.com/cat"
}

func generateRandomTags() []string {
	tags := []string{
		"cute", "fluffy", "sleepy", "playful", "curious",
		"hungry", "lazy", "active", "funny", "adorable",
	}

	// 2-4 tag for each cat
	numTags := 2 + rand.Intn(3)
	selected := make([]string, numTags)

	for i := 0; i < numTags; i++ {
		selected[i] = tags[rand.Intn(len(tags))]
	}

	return selected
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "healthy",
		"service": "Random Cat API",
		"timestamp": time.Now().Unix(),
		"cat_served": rand.Intn(1000),
	})
}