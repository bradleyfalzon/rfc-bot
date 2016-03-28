# RFC-Bot
A bot to autopost recently published RFCs to social media.

Configure Twitter configuration via a .env file.

```
cp .env.example .env
vim .env
```

## Docker

I run the bot as a docker container, see the included `Dockerfile`

```
docker build -t rfc-bot .
docker run -it --rm --name rfc-bot rfc-bot
```

## Systemd

Included is an example systemd unit file for running as a docker container
