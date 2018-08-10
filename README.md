# [mozey/gateway](https://github.com/mozey/gateway)

Serverless API example with Go and AWS Lambda using Apex Gateway


# Quick start

    go get github.com/mozey/gateway

    cd ${GOPATH}/src/github.com/mozey/gateway

Make scripts executable
 
    chmod u+x ./scripts/*.sh

Set env using `config` cmd, default env is dev.
The config files must be in the package root.
Remember to set your `AWS_PROFILE` in the prod config,  
see [aws cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-multiple-profiles.html)
The dev config uses [aws-local](https://github.com/mozey/aws-local)

    ./scripts/config.sh
    
    $(./config)
    
Print env

    printenv | sort | grep -E 'AWS_|APP_'
    
Run dev server

    go run ./cmd/dev/dev.go &

Test
    
    http localhost:${APP_PORT}
    http "localhost:${APP_PORT}/foo?foo=asdf"
    http "localhost:${APP_PORT}/foo?foo=panic"
    http "localhost:${APP_PORT}/bar?key=xxx"
    http "localhost:${APP_PORT}/bar?key=123"
    http "localhost:${APP_PORT}/echo_if_no_match"
    
    
# Create lambda fn and API

Clear env

    unset $(compgen -v APP_)
    
Set prod env
    
    $(./config -env prod)

Build the exe

    ./scripts/build.sh
    
Create lambda fn and API

    ./scripts/create.sh && $(./config -env prod)
    
Call lambda endpoint

    http ${APP_API_ENDPOINT}/foo?foo=foo
    
Add a custom domain to invoke the lambda fn via API gateway,
all request methods and paths are forwarded to the lambda fn
    
    ./config -env prod \
    -key APP_API_PATH -value "" \
    -key APP_API_CUSTOM -value api.mozey.co \
    -key APP_API_DOMAIN -value mozey.co \
    -update
    
    $(./config -env prod) 
    
    ./scripts/domain.sh && $(./config -env prod)
    
Script will print an error message if cert is still validating.
Wait for certificate validation to complete,
then run the script again to finish setup
    
Call API (DNS may take some time to propagate)

    http ${APP_API_CUSTOM_ENDPOINT}/foo?foo=foo
    
Deploy to update the lambda fn
    
    ./scripts/deploy.sh


# Delete lambda fn and API

    $(./config -env prod) && ./scripts/reset.sh


# Makefile

Install dependencies   

    brew install tmux
    
    brew install fswatch

Run with live reload    
    
    $(./config) && make dev
    
Build and deploy lambda fn

    $(./config -env prod) && make deploy
        
tmux workflow
    
    tmux new -d -s mozey-gateway '$(./config}) && make dev'
    
    tmux ls
    
    tmux a -t mozey-gateway # ctrl-b d
    
    tmux send-keys -t mozey-gateway C-c
    
fswatch all files except `dev.out`

    fswatch -or ${APP_DIR}/ -e dev.out


# Show processes listening on port

    lsof -nP -i4TCP:${APP_PORT} | grep LISTEN
    
    
# Docker container with live reload

Build container exe

    export APP_DIR=${GOPATH}/src/github.com/mozey/gateway
    
    $(./config})
    
    ./scripts/build.container.sh
    
Create container

    ./scripts/create.container.sh

Run container

    docker stop mozey-gateway
    docker run -it -d --rm --name mozey-gateway \
    -p ${APP_PORT}:${APP_PORT} \
    -v ./build:/mnt/build \
    mozey-gateway /mnt/build/container.out
    
Test

    http localhost:${APP_PORT}


# [apex/gateway](https://github.com/apex/gateway)

...provides a drop-in replacement for net/http's ListenAndServe 
for use in AWS Lambda & API Gateway, 
simply swap it out for gateway.ListenAndServe

Layout follows [go project layout](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)
and [github here](https://github.com/golang-standards/project-layout)

Use of gateway inspired by [aws-sam-golang-example](https://github.com/cpliakas/aws-sam-golang-example),
but this example does not use sam local

Use monolithic lambda fn (routing is done internal to the fn) by default, 
see [Serverless API with Go and AWS Lambda](https://github.com/mozey/aws-lambda-go/tree/master/examples/books-api).
Pass in an optional path prefix when creating the fn

Use a [shared library](https://stackoverflow.com/a/35060357/639133) 
if the lambda zip [size limit](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
becomes an issue: "Each Lambda function receives an additional 512MB of 


# Caller id

APIGatewayProxyRequestContext contains the information to identify the 
AWS account and resources invoking the Lambda function. 
It also includes Cognito identity information for the caller. 
See [requestContext.Authorizer](https://github.com/apex/gateway/blame/cdfe71df1421609687c01dda11f13ef068784e5b/Readme.md#L31)


# sam local

Alternative to deploy lambda functions and test them locally

Commands below are untested...

    GOOS=linux go build -o main ./cmd/gateway 
    
    # TODO Credentials store error?
    sam local start-api -p 8080

    # Package SAM template
    sam package --template-file ./template.yml --s3-bucket ${APP_BUCKET} \
    --output-template-file packaged.yaml
    
    # Deploy packaged SAM template
    sam deploy --template-file ./packaged.yaml --stack-name ${APP_STACK_NAME} \
    --capabilities CAPABILITY_IAM


