release-prep:
	@gh changelog new --next-version "$(version)"

tag:
	@git tag -a $(version) -m "Release $(version)"
	@git push --follow-tags

lint:
	@golangci-lint run ./...
