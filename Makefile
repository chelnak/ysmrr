release-prep:
	@gh changelog new --next-version "$(version)"

tag:
	@git checkout main && git pull
	@git tag -a $(version) -m "Release $(version)"
	@git push --follow-tags
	@gh release create $(version) -F CHANGELOG.md

lint:
	@golangci-lint run ./...
