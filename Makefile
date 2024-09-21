hello:
	echo "hello"

run:
	cd components
	templ generate
	cd ..
	go run main.go

build:
	cd components
	templ generate
	cd ..
	go build . -o site