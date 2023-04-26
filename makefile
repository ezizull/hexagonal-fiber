migrate:
	go run main.go postgres -m

testing:
	go test ./...

mocking:
	~/go/bin/mockery --all

update:
	# ~/go/bin/swag init
	git add .
	git commit -m "$(commit)"
	git push -u origin master