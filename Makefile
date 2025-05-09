TARGET=floppnet@echeclus.uberspace.de:/var/www/virtual/floppnet/invisible-characters.com/

.PHONY: .build
build:
	go run main.go

.PHONY: clean
clean:
	rm -rf .out

.phony: upload
upload:
	rsync \
		.out/* \
		static/style.css static/.htaccess static/favicon.ico static/robots.txt static/yg3fyzua3sdbgfzwf6xjv5tu483buz49.txt \
		$(TARGET)
