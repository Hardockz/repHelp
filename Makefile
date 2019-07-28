compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/mattermost-bot-test .

compile_and_build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/mattermost-bot-test . && ./build/mattermost-bot-test
