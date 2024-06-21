.PHONY:api
api:
	@goctl1.6 api format -dir=.
	@goctl1.6 api go -api desc.api -dir . -style go_zero  --home  ./goctl
	@goctl1.6 api plugin -plugin goctl-swagger="swagger -filename docs/docs.json" -api desc.api -dir .

.PHONY:model
model:
	@go run cmd/gorm_gen/main.go