## GET /v1/ip
![Badge](https://img.shields.io/badge/200-green)
  
```mermaid
sequenceDiagram
    autonumber
    cli->>sut: GET /v1/ip
    sut-->>cli: 200
```
  
## Event log
#### Event 1
  
GET /v1/ip HTTP/1.1  
Host: sut  
X-Real-Ip: 127.0.0.1  
  

  
---
  
#### Event 2
  
HTTP/1.1 200 OK  
Connection: close  
Content-Type: application/json; charset=UTF-8  
  

  
```json
{
    "global_ip": "127.0.0.1"
}

```
  
---
  