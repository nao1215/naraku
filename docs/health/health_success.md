## GET /v1/health
![Badge](https://img.shields.io/badge/200-green)
  
```mermaid
sequenceDiagram
    autonumber
    cli->>sut: GET /v1/health
    sut-->>cli: 200
```
  
## Event log
#### Event 1
  
GET /v1/health HTTP/1.1  
Host: sut  
  

  
---
  
#### Event 2
  
HTTP/1.1 200 OK  
Connection: close  
Content-Type: application/json; charset=UTF-8  
  

  
```json
{
    "name": "naraku",
    "version": "v0.0.1",
    "revision": "revision-0.0.1"
}

```
  
---
  