TARGET=floppnet@echeclus.uberspace.de:/var/www/virtual/floppnet/invisible-characters.com/

.PHONY: .build
build:
	go run main.go

.PHONY: clean
clean:
	rm -rf .out

.phony: upload
upload:
	rsync .out/*  static/.htaccess static/favicon.ico static/robots.txt $(TARGET)
