How to run the service


Setup the Go server

1. Make sure you have Go installed in the system
2. Use 'go get' to install the dependencies
3. Run 'go run main.go' to get the Go server running

Setup the NodeJS server

1. Make sure you have node and npm installed
2. Use 'npm i' to install the dependecies
3. Run 'node server.js' to start the node server


Three endpoints are defined to access NodeJS server which in turn update the state in Go service for a and b values
1. http://127.0.0.1:8000/server/sendA 
Body : 
{"a":7}

This will update the value for 'a' in running Go program


2. http://127.0.0.1:8000/server/sendB
Body :
{"b":7}

This will update the value for 'b' in running Go program

3. http://127.0.0.1:8000/server/sum

This will return the sum after taking values from Go program

In case the node server is restarted, Go program will still maintain the values of 'a' and 'b' and return the sum based on last values.
