migrate:
	migrate -database "postgres://localhost:5432/crimsonpix_development?sslmode=disable" -path migrations up
