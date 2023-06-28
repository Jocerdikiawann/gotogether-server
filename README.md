### Server Share Trip

Setup :
```bash
touch .env
nano .env
```
set .env :
```bash
MONGO_PORT=
MONGO_PORT_SECOND=
MONGO_PORT_THIRD=
MONGO_USERNAME=
MONGO_PASSWORD=
MONGO_HOST=
MONGO_DB_NAME=
SECRET_KEY=
TELEGRAM_URL=
TELEGRAM_TOKEN=
GROUP_ID=
```

Installation :
```bash
make gen
```

Run :
```bash
make build && make run
```
