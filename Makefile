swag-init:
	swag init -g api/main.go -o api/docs

run-test:
	cd test && go test -run=TestCrawler -v
