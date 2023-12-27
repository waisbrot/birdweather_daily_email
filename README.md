# birdweather_daily_email

A personal project for sending a daily summary of the birds whose calls were identified by my stations on [BirdWeather](https://app.birdweather.com). And also learning some basic Go.

## Usage

Fetch data for 3 stations and send a daily digest to two email addresses:

```
birdweather_daily_email send --email=me@example.com,friend@example.com --station=123,456,789
```

Currently expects that `countEmail.tmpl` (the email template) and `email_password`

## How to update the Birdweather API

1. In the `birdweather` directory
2. Fetch the schema: `npx gql-sdl https://app.birdweather.com/graphql --sdl -o schema.graphql`
3. Regenerate client: `go run github.com/Khan/genqlient`

I got some errors reported in the schema, but making the obvious-seeming fixes by hand appeared to resolve things.

Genqlient depends on `schema.graphql` being the schema, `genqlient.graphql` containing the query-functions we want to be able to run, and `genqlient.yaml` being its config.

## Updating email template

I used [Stripo](https://my.stripo.email) to design the email template. Might be useful for redesign.
