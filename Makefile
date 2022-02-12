NAME=godenticon
GITHUB_USERNAME=poseidon-code
PROJECT=github.com/$(GITHUB_USERNAME)/$(NAME)
VERSION=v0.2.0


init:
	@printf "\n\033[37;1m»\033[0m Initializing go.mod for '\033[36;1m$(PROJECT)\033[0m'...\n"
	@go mod init $(PROJECT)
	@printf "\033[32;1m»\033[0m Initialized '$(PROJECT)'\n"


clean:
	@printf "\033[37;1m»\033[0m Cleaning Golang cached packages...\n"
	@go clean -modcache
	@printf "\033[32;1m»\033[0m Cleaned\n"


tidy:
	@printf "\033[37;1m»\033[0m Tidying Up dependencies...\n"
	@go mod tidy
	@printf "\033[32;1m»\033[0m Finished\n"


publish:
	git tag $(VERSION)
	git push origin $(VERSION)
	GOPROXY=proxy.golang.org go list -m $(PROJECT)@$(VERSION)
