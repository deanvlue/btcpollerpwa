# BTC Poller
This is a simple service that performs to things:
1. It polls for the biffing price on a supported crypto coin symbol (currently supported BTC-MXN and LTC-MXN)
2. It has a server that provides:
    * An API for the last 5 prices of those symbols
    * A Web server to display a PWA with the current 5 prices to chech the changes

## Backend

The backend consist on:
* Sqlite databse with all the historical prices fot both symbols
* An API that queries this DB for the last 5 prices
* A poller that queries an exchange API (this case Bitso) and saves the information to SQLite every X minutes (2 minutes)

## Frontend (in progress)
The PWA is a VueJs applications that displays the last 5 prices of any symbol supported and the times that this prices were changed

## Installations and execution
The btcpoller binary and the database files is enough for this to run:

`./btcpoller`

Will run the app.
TODO: Install it as a Service: [Digital Ocean: Serve a Golang Service thropugh nginx](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04)

## Useful Links.

[Http Request Parameters](https://golangcode.com/get-a-url-parameter-from-a-request/)
[API JWT: TODO](https://golangcode.com/api-auth-with-jwt/)
[SQLite in Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.3.html)
[JSON Response](https://medium.com/@vivek_syngh/http-response-in-golang-4ca1b3688d6)
[Channels and GoRoutines](https://golangbot.com/channels/)
