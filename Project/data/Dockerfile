FROM mongo

COPY data.json /data.json
CMD mongoimport --host mongo --db StockSpace --collection apartments --type json --file /data.json --jsonArray