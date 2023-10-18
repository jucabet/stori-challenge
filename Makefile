install-deps:
	npm install -g aws-cdk-local aws-cdk && \
	pip3 install awscli-local && \
	cd ./infra/stori-challenge-cdk && \
	npm i

up-infra-local:
	docker-compose -f ./infra/localstack/localstack-docker-compose.yml up -d && \
	cd ./infra/stori-challenge-cdk && \
	npm run deploy-local && \
	awslocal dynamodb put-item \
    --table-name stori-transactions-db \
    --item '{"type":{"S":"CONTACT"},"id":{"S":"1"},"contactName":{"S":"Juan Camilo"},"email":{"S":"camilobg1546@gmail.com"}}' \
    --return-consumed-capacity TOTAL && \
	cd ../.. && \
	awslocal s3 cp transactions.csv s3://transactions-bucket-stori/incoming_files/

run-docker-process-tx:
	docker-compose -f ./infra/localstack/process-tx-docker-compose.yml up -d

run-docker-send-reports:
	docker-compose -f ./infra/localstack/send-reports-docker-compose.yml up -d

up-infra-prod:
	cd src/process-transactions && \
	GOOS=linux GOARCH=amd64 go build -o main cmd/main.go && \
	cd ../send-reports && \
	GOOS=linux GOARCH=amd64 go build -o main cmd/main.go && \
	cd ../../infra/stori-challenge-cdk && \
	npm run deploy-prod && \
	aws dynamodb put-item \
    --table-name stori-transactions-db \
    --item '{"type":{"S":"CONTACT"},"id":{"S":"1"},"contactName":{"S":"Juan Camilo"},"email":{"S":"camilobg1546@gmail.com"}}' \
    --return-consumed-capacity TOTAL && \
	cd ../.. && \
	aws s3 cp transactions.csv s3://transactions-bucket-stori/incoming_files/

load-file-local:
	awslocal s3 cp transactions.csv s3://transactions-bucket-stori/incoming_files/

load-file-prod:
	aws s3 cp transactions.csv s3://transactions-bucket-stori/incoming_files/