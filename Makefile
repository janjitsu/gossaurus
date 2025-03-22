test_json:
	go run cmd/build/build.go -from=hr -to=en -title="Serb Dict Devtest" -entries="docs/example.json"
test_tsv:
	go run cmd/build/build.go -from=hr -to=en -title="Serb Dict Devtest" -entries="docs/example.tsv"
help:
	go run cmd/build/build.go -h
install:


