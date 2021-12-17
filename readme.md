# MailToSlack

## Install

```sh
go install github.com/TKMAX777/MailToSlack@latest
sudo apt install postfix
```

## Configuration

### AddEnvironmentVariables

```
SLACK_BOT_TOKEN="xoxb-xxxx"
SLACK_CHANNEL="Cxxxxxxx"
SLACK_ICON_EMOJI=":+1:"
```

### AddMailSettings

```
# /etc/aliases
USERNAME:   "|/path/to/binary"
```

## Example
### Cron logging

- Make shell script to excute binary

```sh
# /usr/local/MailToSlack/start.sh

#!/usr/bin/bash
export SLACK_BOT_TOKEN="xoxb-***"
export SLACK_ICON_EMOJI=":mailbox:"
export SLACK_CHANNEL="C******"

/usr/local/MailToSlack/MailToSlack

```

- Change aliases file

```
# /etc/aliases
cronuser: root, "|/usr/local/MailToSlack/start.sh"
```

- Change cron settings

```
# crontab
MAILTO=cronuser
```

