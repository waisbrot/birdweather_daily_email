# birdweather_daily_email

I got a brace of [BirdWeather](https://app.birdweather.com) PUCs for myself and family members. Rather than checking stations regularly, I thought it'd be nice to get a daily digest email of what birds appeared:

![example email](/img/birds-email-1.png "Example email sent by the program")

The email goes out to all the family bird-listeners summarizing what birds were active yesterday at the different stations. I can see my parents home in the mountains trail behind my sea-level place in winter bird activity, due to colder temperatures.

This project also serves as a vehicle for teaching myself programming in Go. That's the reason for the language selection.

## Usage

### Configuration file

You must use a configuration file, to tell the program which stations to check, while emails to send, and how to connect to an email server. Here's an example configuration:

```yaml
---
email:
  recipients:
    - nathaniel@github.co.uk
    - nate@gmail.org
  sender: My BirdWeather PUC <puc@gmail.edu>
  smtp:
    host: smtp.fastmail.com
    port: 465
    user: me@vanity.name
    pass: setecastronomy
  template: "/etc/birdweather/countEmail.tmpl"
stations:
    - 1985
    - 2122
influx:
  url: https://influxdb.local
  token: 0KPCajYMUIg5k6RpkfW5qvJalN8xWNJYPMdIphS29JauuWgjO6tvKmTAsT-aWsC5NR5IGcDoEGJAzW6J1GuG5w==
  org: 18292773589b0208
  bucket: ebc4fb5a458af26e
```

If you're not running an InfluxDB server (or don't want to collect metrics) just leave that section out. All other configuration elements are required.

Consult your email provider for what you need to send email by SMTP. I thought that Fastmail's instructions were easy but others might be more difficult.

### Running

By default, the configuration file is expected to be at `/etc/birdweather/config.yaml`. You can override this by setting the `BIRDWEATHER_CONFIG_FILE` environment variable:

```sh
BIRDWEATHER_CONFIG_FILE=$PWD/test.config ./birdweather_daily_email send
```

Note the `send` argument at the end. There's some other testing-related sub-commands, but `send` is the one to fetch data and email it.

### Email template

You need to provide an HTML email template file in the `email.template` configuration key. Only HTML emails are currently supported.

I used [Stripo](https://my.stripo.email) to design the provided email template (`countEmail.tmpl`), but I'm not very good at HTML design.

You can create your own, if you like. The template can use the following components, nested:

* `{{ .Day }}`: The weekday-number of the day being fetched (_e.g._ Monday will replace this with "1")
* `{{ range .Stations }} ... {{ end }}`: this will iterate over each station defined in the config file
 * `{{ .Name }}`: The name of the station currently being iterated over. This is whatever the station-owner named it.
 * `{{ .Id }}`: The ID number of the current station
 * `{{ range .Counts }} ... {{ end }}`: this goes inside the `.Stations` loop and iterates over each bird-count for a station
  * `{{ .Name }}`: The common name of the bird
  * `{{ .SciName }}`: The scientific name of the bird
  * `{{ .ImageURL }}`: The BirdWeather-hosted URL for an image of the bird
  * `{{ .ImageCredit }}`: The credit for the bird image. It's scraped data and IME is often garbled.
  * `{{ .Count }}`: The number of times this bird was heard on this day

Fetch data for 3 stations and send a daily digest to two email addresses:

```
birdweather_daily_email send --email=me@example.com,friend@example.com --station=123,456,789
```

Currently expects that `countEmail.tmpl` (the email template) and `email_password`

## Hacking

### Code structure

Cobra is used for subcommands. I could remove that since I'm not using any command-line flags. `cmd/send.go` is where the main logic lives. I tried to split functionality out by module:
* `birdweather` - talking to BirdWeather via GraphQL
* `email` - talking to mail-provider over SMTP (the file says "fastmail" but it's actually agnostic)
* `metrics` - InfluxDB stuff
* `structs` - Go `struct`s that I moved here to break dependency-cycles

### How to update the BirdWeather API

1. In the `birdweather` directory
2. Fetch the schema: `npx gql-sdl https://app.birdweather.com/graphql --sdl -o schema.graphql`
3. Regenerate client: `go run github.com/Khan/genqlient`

I got some errors reported in the schema, but making the obvious-seeming fixes by hand appeared to resolve things.

Genqlient depends on `schema.graphql` being the schema, `genqlient.graphql` containing the query-functions we want to be able to run, and `genqlient.yaml` being its config.
