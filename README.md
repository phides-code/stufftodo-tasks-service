# go-dynamodb-service-template

A Go project template providing CRUD services for a DynamoDB table, using AWS Lambda and API Gateway, deployed with AWS SAM and GitHub Actions.

### Deploy manually

-   `make deploy`

### Run locally

-   `make build && sam local start-api --port 8000`

### Setup GitHub actions

Once the repo is setup on GitHub, add AWS secrets to GitHub Actions for this repo:

-   `gh secret set AWS_ACCESS_KEY_ID`
-   `gh secret set AWS_SECRET_ACCESS_KEY`

### Test

-   `curl -X POST http://localhost:8000/tasks -H "Content-Type: application/json" -d @post-data.json |jq .`
-   `curl http://localhost:8000/tasks |jq .`
