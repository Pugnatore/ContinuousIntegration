### Coding Challenge Guidelines
We are running a growing company and have recently noticed that we keep on introducing new errors and flaws into our code basis. To keep up quality without reducing development speed too much, we want to optimize integration and review of code in our repository.

Your task is to automate the testing of our code in a pull request (e.g. on Bitbucket, GitHub, GitLab or any other such platform). Whenever we merge code we want to be sure that we have not introduced new errors or reduced code quality.

Please add a simple qualification test by building and running our application.

### Evaluation Criteria
- We want a smooth experience for our developers
- Make it reproducible any time anywhere, i.e. on a CI system or on the computer of any developer
- Please assume that we will have long-running tests and prepare

### The Commands You Will Need
1. `go mod vendor` will download dependencies
2. `go build` will build the application
3. `go test ./...` will run all unit tests
4. `go clean -testcache ./...` will make sure tests are run even though there is a cached result

Please note to put this project under `$GOPATH/src` to avoid issues with package lookup

### Remarks on Technologies to Use
1. Feel free to fork this code into any code repository that suits your needs. Please make sure to push the final result to our repository in the end!
2. You may host your CI system locally, remotely or use any SaaS

### Useful Links
- On go test coverage: [https://tip.golang.org/doc/go1.2#cover]

### CodeSubmit

Please organize, design, test, and document your code as if it were
going into production - then push your changes to the master branch.

Have fun coding! 🚀
