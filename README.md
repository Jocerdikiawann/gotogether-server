### Server Share Trip

Installation :
```bash
docker compose up -d
```
then open the terminal, enter into docker with the command:
```bash
docker exec -it server_share_trip-cassandra-1 bash
```
Then :
```bash
cqlsh cassandra -u your_username -p your_password
```
in CQL mode you have to create keyspace first :
```
cassandra@cqlsh> CREATE KEYSPACE your_keyspace_name WITH REPLICATION = {'class' : 'NetworkTopologyStrategy','datacenter1' : 1} AND DURABLE_WRITES = true;
```
Don't forget to check whether the keyspace has been created or not with the command:
```
cassandra@cqlsh> describe keyspace
```
after that :
```
cassandra@cqlsh> use your_keyspace_name
```
done, please create a COLUMN database.
You can use [Retool](https://retool.com/) to monitor your database.