# naraku (奈落) 
"Naraku" is a repository created to assess the quality of End-to-End (E2E) tests using [go-spectest/spectest](https://github.com/go-spectest/naraku). Spectest is an E2E testing framework, and to confirm its ease of use, it is essential to create test cases for APIs.

I chose to build APIs to test a wide range of API types. If a useful API is created, it will be deployed using Lambda + API Gateway.

## API Reference
Please see the [API Reference](https://go-spectest.github.io/naraku/).

### /v1/health
```bash
curl -X GET "http://localhost:8080/v1/health" -H  "accept: application/json" | jq .
{
  "name": "naraku",
  "version": "unknown",
  "revision": "unknown"
}
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## The origin of the name
The name "Naraku" was borrowed from a character in the series InuYaSha. Naraku is a being formed by the accumulation of numerous demons around a human. The inspiration for the name comes from the way multiple APIs gather in the spectest framework. Naraku is a half-demon with excellent defensive abilities. I also hope that these characteristics will manifest in spectest.

Moreover, "Naraku" also means falling into hell. Don't you think creating and testing numerous APIs can be challenging?