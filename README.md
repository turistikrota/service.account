# Run Project

## With Docker

```bash
docker build --build-arg GITHUB_USER=<USER_NAME> --build-arg GITHUB_TOKEN=<ACCESS_TOKEN> -t api.turistikrota.com/account .   
docker run --name account.api.turistikrota.com -p 6014:6014 --env-file .env --network="turistikrota" api.turistikrota.com/account:latest
```
