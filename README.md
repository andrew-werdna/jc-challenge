# This is a simple webserver with 4 endpoints

* `POST /hash`
* `GET /hash/{/[1-9]+/s}`
* `GET /stats`
* `GET /shutdown`

## Caveats

The basic requirements for each endpoint have been satisfied, but as far as **Production Worthy** code, I would say this is more of a *quick and dirty* solution. I would want to finish with all tests, (both unit and e2e) as well as use godoc to generate documentation for this module to call it production worthy simply due to the incomplete tests.

### Room for Improvement

While I've used multithreaded code in Go regularly, none of us on my team have made regular use of channels. I will continue to experiment and work on this in the near future to see if I can't re-structure things to use channels instead of mutexes. I also would like to get better at testing multi-threaded scenarios and finish the unit and e2e tests.

### Other Notes and E2E Testing

I've included Postman collections for *manual* e2e testing in the postman directory. Just import the collection and environment file into your Postman workspace and anyone can run the server and then execute the requests while watching the terminal for logging. Postman provides a ver easy interface for writing e2e tests and if you install their CLI tool called **Newman** you can simple run `newman path/to/collection.json -e path/to/environment-file.json` to while the server is running to execute the e2e tests.
