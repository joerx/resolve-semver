default: build

build:
	go build -o ${BINARY}

release:
# ifndef GPG_FINGERPRINT
# 	@ echo "GPG_FINGERPRINT must be set"
# 	@ exit 1
# endif
	goreleaser release --rm-dist -p2

snapshot-release:
	goreleaser release --rm-dist -p2 --snapshot

test: 
	go test .
