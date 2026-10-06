.PHONY: site clean

site:
	go run main.go

clean:
	rm -rf public
