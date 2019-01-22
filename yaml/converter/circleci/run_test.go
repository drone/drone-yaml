package circleci

const testRun = `
- run:
    name: test
    command: go test
`

const testRunShort = `
- run: go test
`
