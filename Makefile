.PHONY: site clean

site:
	go run main.go
	@cp snowz.html public/snowz.html
	@cp -r snowz public/snowz

clean:
	rm -rf public
