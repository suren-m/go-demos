module github.com/suren-m/go-demos

go 1.17

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

require services v1.0.0

replace services v1.0.0 => ./services
