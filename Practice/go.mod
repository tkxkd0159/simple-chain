module github.com/tkxkd0159/go-chain

go 1.16

require (
	rsc.io/quote/v3 v3.1.0
	rsc.io/sampler v1.99.99 // indirect
)


replace (
	github.com/tkxkd0159/go-chain/hello => "./hello"
	github.com/tkxkd0159/go-chain/hello/nums => "./hello/nums"
)