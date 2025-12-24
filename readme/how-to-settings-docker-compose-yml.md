### How to Settings docker-compose.yml

Based on the docker-compose.yml example above:

1. Port 7200 is used for the web admin dashboard. You can change the port in the hermawan_monitora_webadmin PORT environment variable.

    **NOTE:**  
    If you change the PORT to 9000, go to http://localhost:9000 to open the web admin dashboard.

2. The hermawan_monitora_redis service will create a Redis service in your local. You can remove it if you already have a Redis server, but make sure to adjust environment variables below correctly to connect to your Redis server.

    [All Docker Environment Variables information](./docker-environment-variables.md)

3. Make sure the network mode in the hermawan_monitora_worker_mp is host, because the service needs to bind directly to the host network, with no network isolation to monitor services. Actually you can change it to another network mode than host if you want to, but make sure you have knowledge about Docker and computer networking.
4. In the volume map, there are two important things you need to map:
    a. SQLite database file in every Hermawan Monitora service, map your local SQLite database file to `/opt/hermawan-monitora/data/sqlite`.
    b. Log files folder will contain log files for troubleshooting.

        * Map log folder for hermawan_monitora_bgadmin to `/var/log/hermawan-monitora/bgadmin`.
        * Map log folder for hermawan_monitora_webadmin to `/var/log/hermawan-monitora/webadmin`.
        * Map log folder for hermawan_monitora_worker_mp to `/var/log/hermawan-monitora/worker-mp`.

    c. Pic files folder in hermawan_monitora_webadmin, folder to keep uploaded pictures, map to `/opt/hermawan-monitora/data/pic`.
    d. Certificate files to access Redis if needed.
