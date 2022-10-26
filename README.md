
# Apple TV rebooter

Simple application to reboot an Apple TV Homekit hub connected with a smart plug.

Will reboot on the defined schedule or when the Apple TV is not responding.

## Limitations

Only works with smart plugs that support the [TP-Link Tapo](https://www.tapo.com/) API.

Has to be disabled when the Apple TV is updating.

## Environment variables

ATV_ADDRESS: address of the Apple TV\
PING_TIME: time in seconds to run ping command\
CRON: Cron schedule expression\
TAPO_IP: IP of your Tapo device\
TAPO_USERNAME: Your Tapo username\
TAPO_PASSWORD: Your Tapo password

## Docker reference

 * docker build -t go-apt-reboot . --no-cache
 * docker run --env-file ./.env go-apt-reboot
 * docker rm -f $(docker ps -aq)
 * docker rmi $(docker images -q)

## References

 * [P100-go Library](https://github.com/artemvang/p100-go)
 * [Cron Expressions](https://crontab.guru/every-day-at-2am)
 * [Smallest Golang Docker Image](https://klotzandrew.com/blog/smallest-golang-docker-image)
 
## Disclaimer

GitHub Copilot was used to write part of the code and documentation.
