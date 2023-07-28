run:
	go run main.go httpsrv

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