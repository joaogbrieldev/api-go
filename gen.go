package main

//go:generate tern migrate --migrations ./internal/pgstore/migrations --config ./internal/pgstore/migrations/tern.conf

//go:generate sqlc generate -f ./internal/pgstore/sqlc.yaml

//go:generate goapi-gen --package=spec --out ./internal/spec/journey.gen.spec.go ./internal/spec/journey.spec.json