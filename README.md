# This is a simple webserver with 4 endpoints

* `POST /hash`
* `GET /hash/{/[1-9]+/s}`
* `GET /stats`
* `GET /shutdown`

## Caveats

The basic requirements for each endpoint have been satisfied, but as far as **Production Worthy** code, I would say this is more of a *quick and dirty* solution. I would want to finish with all tests, (both unit and e2e) as well as use godoc to generate documentation for this module to call it production worthy.
