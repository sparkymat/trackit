migrate:
	migrate -database "postgres://localhost:5432/trackit_development?sslmode=disable" -path migrations up
