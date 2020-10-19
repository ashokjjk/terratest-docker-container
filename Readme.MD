## Test your Docker image using Terratest
<p> Terratest is a testing tool form Gruntwork with Go library which uses your Dockerfile and runs the test and performs validation on the result of the container.</p>

<b>Prerequisite</b>

1. Go compiler 1.14 or higher
2. Docker for desktop

<b>How it works</b>

* Code directory contains Dockerfile and a NodeJS application
* The Node application has exactly one endpoint "/"
* Endpoint returns <code> Hello World from NodeJS!! #Happy coding# </code> as a result on GET call
* Test directory contains the Go code for performing the test.
* Notice the name of the go file. Should end with _test.go
* Once all the files are placed open shell or cmd navigate to test directory.
```
go mod init test
go test -v
```
<b>Expected result</b>

```
=== RUN   TestDockerContainerEndpoint
TestDockerContainerEndpoint 2020-10-19T19:31:50+05:30 logger.go:66: Running 'docker build' in ../code/
TestDockerContainerEndpoint 2020-10-19T19:31:50+05:30 logger.go:66: Running command docker with args [build --tag testnode/terratest ../code/]
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Sending build context to Docker daemon   25.6kB
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Step 1/7 : FROM node:12
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> 28faf336034d
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Step 2/7 : WORKDIR /usr/src/app
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> Using cache
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> d94d230cb713
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Step 3/7 : COPY package*.json ./
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> Using cache
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> 7eb90de5ffb6
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Step 4/7 : RUN npm install
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> Using cache
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> 4fa47bd1b49b
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Step 5/7 : COPY . .
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> Using cache
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> d7ee2db4abe0
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Step 6/7 : EXPOSE 8080
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> Using cache
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> b55a245cb102
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Step 7/7 : CMD [ "node", "server.js" ]
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> Using cache
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66:  ---> dfb5f65832f0
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Successfully built dfb5f65832f0
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Successfully tagged testnode/terratest:latest
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is recommended to double check and reset permissions for sensitive files and directories.
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Running 'docker run' on image 'testnode/terratest'
TestDockerContainerEndpoint 2020-10-19T19:31:51+05:30 logger.go:66: Running command docker with args [run --detach -p=8081:8080 testnode/terratest]
TestDockerContainerEndpoint 2020-10-19T19:31:52+05:30 logger.go:66: f980b196ea6e1a4dad72f5be27a3c0a60f553dc3dc1fe44b09da67d8c0f066f1
TestDockerContainerEndpoint 2020-10-19T19:31:53+05:30 retry.go:72: HTTP GET to URL http://localhost:8081
TestDockerContainerEndpoint 2020-10-19T19:31:53+05:30 http_helper.go:32: Making an HTTP GET call to URL http://localhost:8081
TestDockerContainerEndpoint 2020-10-19T19:31:53+05:30 retry.go:84: HTTP GET to URL http://localhost:8081 returned an error: Get "http://localhost:8081": EOF. Sleeping for 3s and will try again.
TestDockerContainerEndpoint 2020-10-19T19:31:56+05:30 retry.go:72: HTTP GET to URL http://localhost:8081
TestDockerContainerEndpoint 2020-10-19T19:31:56+05:30 http_helper.go:32: Making an HTTP GET call to URL http://localhost:8081
f980b196ea6e1a4dad72f5be27a3c0a60f553dc3dc1fe44b09da67d8c0f066f1
TestDockerContainerEndpoint 2020-10-19T19:31:56+05:30 logger.go:66: Running 'docker stop' on containers '[f980b196ea6e1a4dad72f5be27a3c0a60f553dc3dc1fe44b09da67d8c0f066f1]'
TestDockerContainerEndpoint 2020-10-19T19:31:56+05:30 logger.go:66: Running command docker with args [stop --time 1 f980b196ea6e1a4dad72f5be27a3c0a60f553dc3dc1fe44b09da67d8c0f066f1]
TestDockerContainerEndpoint 2020-10-19T19:31:58+05:30 logger.go:66: f980b196ea6e1a4dad72f5be27a3c0a60f553dc3dc1fe44b09da67d8c0f066f1
--- PASS: TestDockerContainerEndpoint (8.21s)
PASS
ok      test    8.651s
```

<p>On successful execution terratest Go library builds the application using the Dockerfile,  runs the image and publish on selected port and can also inject environment if needed, performs the validation and destroys the environment.</p>