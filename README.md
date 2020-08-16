# Caching server

Basic HTTP caching web service. It transparently proxies all HTTP requests and only caches GET requests.

## Usage

- Run the server using `go run ./...`
- Do a network request with modified Host header of your choice.
  ```
  curl -H "Host: example.com" http://localhost:8080/
  ```
