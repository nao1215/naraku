[![Gosec](https://github.com/go-spectest/naraku/actions/workflows/gosec.yml/badge.svg)](https://github.com/go-spectest/naraku/actions/workflows/gosec.yml)
[![LinuxUnitTest](https://github.com/go-spectest/naraku/actions/workflows/linux_test.yml/badge.svg)](https://github.com/go-spectest/naraku/actions/workflows/linux_test.yml)
[![MacUnitTest](https://github.com/go-spectest/naraku/actions/workflows/mac_test.yml/badge.svg)](https://github.com/go-spectest/naraku/actions/workflows/mac_test.yml)
[![WindowsUnitTest](https://github.com/go-spectest/naraku/actions/workflows/windows_test.yml/badge.svg)](https://github.com/go-spectest/naraku/actions/workflows/windows_test.yml)
[![reviewdog](https://github.com/go-spectest/naraku/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/go-spectest/naraku/actions/workflows/reviewdog.yml)
![Coverage](https://raw.githubusercontent.com/go-spectest/octocovs-central-repo/main/badges/go-spectest/naraku/coverage.svg)
# naraku (奈落) 
"Naraku" is a repository created to assess the quality of End-to-End (E2E) tests using [go-spectest/spectest](https://github.com/go-spectest/spectest). Spectest is an E2E testing framework, and to confirm its ease of use, it is essential to create test cases for APIs.

I chose to build APIs to test a wide range of API types. If a useful API is created, it will be deployed using Lambda + API Gateway.

## API Reference
Please see the [API Reference](https://go-spectest.github.io/naraku/).

### GET /v1/health
```bash
curl -X GET "http://localhost:8080/v1/health" -H  "accept: application/json" | jq .
{
  "name": "naraku",
  "version": "v0.0.1",
  "revision": "79564c979263a1fa893f7d6f2505fb0c26197b4c"
}
```

### GET /v1/os/:distribution
Currently, the android, debian distributions are supported.
```bash
$ curl -X GET http://localhost:8080/v1/os/android | jq .
{
  "items": [
    {
      "name": "Android",
      "version": "1.0",
      "release": "2008-09-23T00:00:00Z",
      "api_level": [
        1
      ]
    },
    {
      "name": "Android",
      "version": "1.1",
      "release": "2009-02-09T00:00:00Z",
      "api_level": [
        2
      ]
    },
    {
      "name": "Android",
      "version": "1.5",
      "code_name": "Cupcake",
      "release": "2009-04-27T00:00:00Z",
      "api_level": [
        3
      ]
    },
    ...
}
```

### GET /v1/ulid
`GET /v1/ulid` endpoint was implemented by [@aqyuki](https://github.com/aqyuki).
```bash
curl -X GET http://localhost:8080/v1/ulid | jq .
{
  "ulid": "01HD3107RHMXD5K0R02TVDPZHB"
}
```

## Contribution (How to develop)
Thank you for expressing your willingness to contribute. Contribution is not limited to code modifications alone. If you can provide a simple API specification, there is a possibility of implementing its features. Also, your star on the project repository serves as a source of motivation for our development efforts!

### Prerequisites
You need to install the following tools to develop this project.
| Category | Technology |
|:---|:---|
| DB Accessor | [sqlc](https://github.com/sqlc-dev/sqlc) |
| Swagger generator | [echo-swagger](https://github.com/swaggo/swag) |
| ER diagram| [tbls](https://github.com/k1LoW/tbls)|
| Test framework | [ginkgo](https://github.com/onsi/ginkgo)|
| Dependency Injection |[wire](https://github.com/google/wire)|
|Build Tools| [make](https://www.gnu.org/software/make/)|

Please install make and Golang using package managers, and other Golang-based tools can be installed using the following command.
```bash
make install-tools 
```

### How to execute server
```bash
make server
```

### How to execute test
```bash
make test
```

### How to add API
1. Choose a proposal for the API you want to create from the Issues section. If there isn't a proposal for the API you want to create, please create a new Issue.
2. Fork the go-spectest/naraku repository.
3. Create a new branch in the forked repository.
4. Implement the API. Writing unit tests is optional; @nao1215 will add unit tests later if they are not provided.
5. Create a Pull Request.

At this stage, please refrain from creating APIs that connect to databases or AWS. Initially, @nao1215 will implement sample code for connecting to DB and AWS. 

I expect the APIs you create to be small in scale. You can write the API entry points in the api directory and the corresponding logic within the `api` package. The APIs should be small enough that it is acceptable to write the logic within the `api`` package

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## The origin of the name
The name "Naraku" was borrowed from a character in the series InuYaSha. Naraku is a being formed by the accumulation of numerous demons around a human. The inspiration for the name comes from the way multiple APIs gather in the spectest framework. Naraku is a half-demon with excellent defensive abilities. I also hope that these characteristics will manifest in spectest.

Moreover, "Naraku" also means falling into hell. Don't you think creating and testing numerous APIs can be challenging?

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=go-spectest/naraku&type=Date)](https://star-history.com/#go-spectest/naraku&Date)
