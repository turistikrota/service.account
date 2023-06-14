# Run Project

## With Docker

```bash
docker build -t api.turistikrota.com/account .        
docker run -p 6014:6014 --env-file .env --network="turistikrota" api.turistikrota.com/account
```
