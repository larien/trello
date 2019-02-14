trello:
	@echo "Building and running trello-go"
	docker build -t trello .
	docker run trello