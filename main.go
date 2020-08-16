package main

import (
	"log"
	"net/http"
)

const scheme = "https://"

func main() {
	inMemStore := newStore()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		upstreamURI := scheme + r.Host + r.RequestURI

		// Cache GET requests only.
		if r.Method != http.MethodGet {
			_, _, err := proxyAndRespond(upstreamURI, w, r)
			if err != nil {
				log.Println("Error in req: ", err)
			}
			return
		}

		// Check response in in-memory cache.
		data, ok := inMemStore.get(upstreamURI)
		if ok {
			log.Println("Cache hit! URI = ", upstreamURI)
			addHeaders(w, data.headers)
			w.Write(data.body)
			return
		}

		log.Println("Cache miss! URI = ", upstreamURI)

		content, headers, err := proxyAndRespond(upstreamURI, w, r)
		if err != nil {
			log.Print("Error in GET req: ", err)
			return
		}

		// Update in-memory cache
		inMemStore.set(upstreamURI, content, headers)
	})

	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
