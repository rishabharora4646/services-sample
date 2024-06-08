# Request and response from api

## POST /users/login
Get token for a user.
### Request
```javascript
{
    "email" : "test@test.com",
    "password" : "test@test"
}
```
### Response
```javascript
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjcGhqanJ0YWFzMWM3MTRtcmU3aCIsImV4cCI6MTcyNTYzNTE3OH0.VmFhyFc28dxrofOToD3_1ZpsRb3fdSi5aCnxIdFcx5A"
}
```

## GET /orgs?offset={offset}
List orgs for which user has access. Requires jwt token x-auth-token header. Limit is set to 25.
### Response
```javascript
[
    {
        "orgId": "cphjjrtaas1c714mre7i",
        "orgName": "Test Org 1",
        "createdAt": "2024-06-08T16:50:47Z"
    }
]
```

## GET /orgs/{orgId}/services?offset={offset}&name={serviceName}
List services for org. Requires jwt token x-auth-token header. Limit is set to 25. Services can be filtered by name using name parameter.
### Response
```javascript
[
    {
        "serviceId": "cphjjrtaas1c714mre7j",
        "serviceName": "Locate Us",
        "serviceDescription": "This is Locate Us service",
        "createdAt": "2024-06-08T16:51:47Z",
        "updatedAt": "2024-06-08T16:51:47Z",
        "versionCount": 2
    },
    {
        "serviceId": "cphjjrtaas1c714mre7k",
        "serviceName": "Contact Us",
        "serviceDescription": "Contact Us service description.",
        "createdAt": "2024-06-08T16:52:47Z",
        "updatedAt": "2024-06-08T16:52:47Z",
        "versionCount": 1
    }
]
```

## GET /orgs/services/{serviceId}
Details for a service. Requires jwt token x-auth-token header.
### Response
```javascript
{
    "serviceId": "cphjjrtaas1c714mre7j",
    "serviceName": "Locate Us",
    "serviceDescription": "This is Locate Us service",
    "createdAt": "2024-06-08T16:51:47Z",
    "updatedAt": "2024-06-08T16:51:47Z",
    "versionCount": 2,
    "versions": [
        {
            "versionId": "cphjjrtaas1c714mre7l",
            "versionName": "version 1",
            "serviceHost": "locate-v1.abc.com",
            "servicePort": 8080,
            "isActive": false,
            "createdAt": "2024-06-08T16:53:47Z",
            "updatedAt": "2024-06-08T16:53:47Z"
        },
        {
            "versionId": "cphjjrtaas1c714mre7m",
            "versionName": "version 2",
            "serviceHost": "locate-v2.abc.com",
            "servicePort": 8081,
            "isActive": true,
            "createdAt": "2024-06-08T16:54:47Z",
            "updatedAt": "2024-06-08T16:54:47Z"
        }
    ]
}
```