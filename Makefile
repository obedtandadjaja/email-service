ENV             ?= "development"
APP_HOST        ?= "localhost"
APP_PORT        ?= "8080"
MAILGUN_DOMAIN  ?= "secret"
MAILGUN_API_KEY ?= "secret"

run:
	export ENV=$(ENV) \
         APP_HOST=$(APP_HOST) \
         APP_PORT=$(APP_PORT) \
         MAILGUN_DOMAIN=$(MAILGUN_DOMAIN) \
         MAILGUN_API_KEY=$(MAILGUN_API_KEY); \
	go clean; \
	go build; \
	./email-service
