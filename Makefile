.PHONY: site clean

site:
	go run main.go
	@cp -R static/. public/

clean:
	rm -rf public
