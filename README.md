# Backend
1. REST API for managing application

# Frontend
1. Single page react app with dependency on REST API

# TODO
Authentication :
If user logs in with user/pass that exists, create struct 

`
auth_token (
    expirationDate string
)
`
and encrypt it with secret.

Send that to front end, requests in future decrypt it check if its still valid
