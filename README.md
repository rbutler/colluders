# colluders

## running the application
`make build` will build and create a `target/colluders` application

`GROUP_ID=<GROUPID> TOKEN=<TOKEN> make run` will build and run

## example query
curl GET -H "Content-Type: application/json" https://api.groupme.com/v3/groups/YOUR_GROUP/messages\?token\=YOUR_TOKEN
