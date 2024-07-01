## HNG Stage 1 Task

My submission for the [HNG](https://hng.tecinternshiph) stage 1 task for the backend track.

This repo contains code for an API endpoint that returns a JSON response of this type:

```sh
{
  "client_ip": "127.0.0.1",
  "location": "New York",
  "greeting": "Hello, Mark!, the temperature is 11 degrees Celcius in New York"
}
```

### Sample Query 

To test this run:
```sh
curl https://c2dgj8vh-3000.euw.devtunnels.ms/hello?visitor_name="Mark"
```

### Installation

To run the code on your machine follow the following steps

* Create a `.env` 