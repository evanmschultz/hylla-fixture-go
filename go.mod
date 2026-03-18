module github.com/evanmschultz/hylla-fixture-go

go 1.24.0

require github.com/google/uuid v1.6.0

require (
	github.com/evanmschultz/hylla-fixture-go-2 v0.0.0
	github.com/evanmschultz/hylla-fixture-go-2/codec v0.0.0
	github.com/evanmschultz/hylla-fixture-go-2/x/clock v0.0.0
)

replace github.com/evanmschultz/hylla-fixture-go-2 => ../hylla-fixture-go-2

replace github.com/evanmschultz/hylla-fixture-go-2/codec => ../hylla-fixture-go-2/codec

replace github.com/evanmschultz/hylla-fixture-go-2/x/clock => ../hylla-fixture-go-2/x/clock
