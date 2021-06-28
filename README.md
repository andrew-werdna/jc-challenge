# This is a simple webserver with 4 endpoints

* `POST /hash`
* `GET /hash/{/[1-9]+/s}`
* `GET /stats`
* `GET /shutdown`

The `POST /hash` endpoint takes one form field of the form `password=your-password` and returns an integer key. It will hash the password, then base64 string encode it in a background process that waits 5 seconds before doing the work to simulate a long running process.

To access this hash, call the `GET /hash/` endpoint and put the integer key on the end of the url. So if you got back 3, you would make a GET request to `/hash/3` and will be returned the hash (so long as at least 5 seconds have gone by).

As passwords the POST endpoint gets called over and over some basic data is being kept track of. To access this information, make a GET request to the `/stats` endpoint. This will return how many POST requests have been made and the average time spent returning the integer in a JSON object.

Lastly, to shutdown the server, make a GET request to the `/shutdown` endpoint. This will immediately return a note saying the server is shutting down, however it will wait for all current passwords that are in the process of being hashed to finish before it finally exits.

## Caveats

The basic requirements for each endpoint have been satisfied, but as far as **Production Worthy** code, I would say this is more of a *quick and dirty* solution. I would want to finish with all tests, (both unit and e2e) as well as use godoc to generate documentation for this module to call it production worthy simply due to the incomplete tests.

### Room for Improvement

While I've used multithreaded code in Go regularly, none of us on my team have made regular use of channels. I will continue to experiment and work on this in the near future to see if I can't re-structure things to use channels instead of mutexes. I also would like to get better at testing multi-threaded scenarios and finish the unit and e2e tests.

### Other Notes and E2E Testing

I've included Postman collections for *manual* e2e testing in the postman directory. Just import the collection and environment file into your Postman workspace and anyone can run the server and then execute the requests while watching the terminal for logging. Postman provides a ver easy interface for writing e2e tests and if you install their CLI tool called **Newman** you can simple run `newman path/to/collection.json -e path/to/environment-file.json` to while the server is running to execute the e2e tests.
