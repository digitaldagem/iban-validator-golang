# iban-validator-golang

This repository holds the source code for a golang chi web-service that validates an International Bank Account Number (IBAN).

It's ValidateIBAN(iban string) function checks to see:
* if an IBAN has a valid country code
* if an IBAN has a valid length when it has a valid country code
* if the modulo operation on the IBAN, in a piece-wise manner (9 pieces at a time), is equal to one

## When running the web-service locally

API endpoint for the validation:

`GET http://localhost:8080/api/validate_iban/{iban}`

A successful validation call will respond with a string message saying so in a status 200 response entity. An unsuccessful validation call will respond with a string message saying why it did not validate in a status 400 response entity.