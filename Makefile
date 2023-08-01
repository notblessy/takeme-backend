.PHONY: run migration

run: check-modd-exists
	@modd -f ./.modd/httpsrv.modd.conf

migrate_up=go run main.go migrate --direction=up
migrate_down=go run main.go migrate --direction=down

migration:
	@if [ "$(direction)" = "" ] && [ "$(new)" = "" ]; then\
		$(migrate_up);\
	elif [ "$(new)" != "" ] && [ "$(direction)" = "" ]; then\
		go run main.go migrate --new=$(new);\
	elif [ "$(direction)" != "" ] && [ "$(new)" = "" ]; then\
		go run main.go migrate --direction=$(direction);\
	else\
		go run main.go migrate --direction=$(direction);\
	fi

check-modd-exists:
	@modd --version > /dev/null

model/mock/mock_user_repository.go:
	mockgen -destination=model/mock/mock_user_repository.go -package=mock github.com/notblessy/takeme-backend/model UserRepository

model/mock/mock_user_usecase.go:
	mockgen -destination=model/mock/mock_user_usecase.go -package=mock github.com/notblessy/takeme-backend/model UserUsecase

mockgen: model/mock/mock_user_repository.go \
	model/mock/mock_user_usecase.go

unit-test: mockgen
	go test ./... -v --cover

test: unit-test