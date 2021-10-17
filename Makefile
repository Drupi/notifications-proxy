CF_PATH=deployments/cloudformation

.PHONY: build
build:
	sam build -t ${CF_PATH}/template.yaml

.PHONY: test
test:
	sam local start-api

.PHONY: deploy
deploy:
	sam deploy -t .aws-sam/build/template.yaml --stack-name notifications-proxy --resolve-s3 --capabilities CAPABILITY_IAM