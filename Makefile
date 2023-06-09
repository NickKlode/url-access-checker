build:
	docker build -t inhousead-test .
run:
	docker run --name=inhousead-test-web -p 8080:8080 inhousead-test