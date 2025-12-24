### Env Files

1. Hermawan Monitora Background-Server:

    * REDIS_HOST
    Redis server host.
    * REDIS_PORT
    Redis server port.
    * REDIS_DB
    Database to be selected after connecting to the server.
    * REDIS_PASSWORD
    Redis password.
    * REDIS_MAX_RETRIES
    Maximum number of retries before giving up.
    * REDIS_MIN_RETRY_BACKOFF
    Minimum backoff between each retry.
    * REDIS_MAX_RETRY_BACKOFF
    Maximum backoff between each retry.
    * REDIS_DIAL_TIMEOUT
    Dial timeout for establishing new connections.
    * REDIS_READ_TIMEOUT
    Timeout for socket reads. If reached, commands will fail with a timeout instead of blocking.
    * REDIS_WRITE_TIMEOUT
    Timeout for socket writes. If reached, commands will fail with a timeout instead of blocking.
    * REDIS_POOL_SIZE
    Maximum number of socket connections.
    * REDIS_MIN_IDLE_CONNS
    Minimum number of idle connections which is useful when establishing new connection is slow.
    * REDIS_MAX_CONN_AGE
    Connection age at which client retires (closes) the connection.
    * REDIS_POOL_TIMEOUT
    Amount of time client waits for connection if all connections are busy before returning an error.
    * REDIS_IDLE_TIMEOUT
    Amount of time after which client closes idle connections.
    Should be less than server's timeout.
    * REDIS_IDLE_CHECK_FREQUENCY
    Frequency of idle checks made by idle connections reaper, but idle connections are still discarded by the client if Idle timeout is set.
    * REDIS_CLIENT_CRT_FILEPATH
    Crt certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * REDIS_CLIENT_KEY_FILEPATH
    Key certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * REDIS_CLIENT_CA_FILEPATH
    CA certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * LOG_DIRPATH
    Log folder.
    * SQLITE_DB_FILEPATH
    SQLite database file path.

2. Hermawan Monitora Web-Admin:

    * JWT_SIGNATURE_KEY
    Secret key for web admin, it’s really important for the web admin security, use unique key combination which unpredictable and keep it secret.
    * PORT
    Web admin port.
    * PIC_DIRPATH
    Uploaded pictures folder path.
    * REDIS_HOST
    Redis server host.
    * REDIS_PORT
    Redis server port.
    * REDIS_DB
    Database to be selected after connecting to the server.
    * REDIS_PASSWORD
    Redis password.
    * REDIS_MAX_RETRIES
    Maximum number of retries before giving up.
    * REDIS_MIN_RETRY_BACKOFF
    Minimum backoff between each retry.
    * REDIS_MAX_RETRY_BACKOFF
    Maximum backoff between each retry.
    * REDIS_DIAL_TIMEOUT
    Dial timeout for establishing new connections.
    * REDIS_READ_TIMEOUT
    Timeout for socket reads. If reached, commands will fail with a timeout instead of blocking.
    * REDIS_WRITE_TIMEOUT
    Timeout for socket writes. If reached, commands will fail with a timeout instead of blocking.
    * REDIS_POOL_SIZE
    Maximum number of socket connections.
    * REDIS_MIN_IDLE_CONNS
    Minimum number of idle connections which is useful when establishing new connection is slow.
    * REDIS_MAX_CONN_AGE
    Connection age at which client retires (closes) the connection.
    * REDIS_POOL_TIMEOUT
    Amount of time client waits for connection if all connections are busy before returning an error.
    * REDIS_IDLE_TIMEOUT
    Amount of time after which client closes idle connections.
    Should be less than server's timeout.
    * REDIS_IDLE_CHECK_FREQUENCY
    Frequency of idle checks made by idle connections reaper, but idle connections are still discarded by the client if Idle timeout is set.
    * REDIS_CLIENT_CRT_FILEPATH
    Crt certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * REDIS_CLIENT_KEY_FILEPATH
    Key certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * REDIS_CLIENT_CA_FILEPATH
    CA certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * LOG_DIRPATH
    Log folder.
    * SQLITE_DB_FILEPATH
    SQLite database file path.

3. Hermawan Monitora Worker-MP:

    * REDIS_HOST
    Redis server host.
    * REDIS_PORT
    Redis server port.
    * REDIS_DB
    Database to be selected after connecting to the server.
    * REDIS_PASSWORD
    Redis password.
    * REDIS_MAX_RETRIES
    Maximum number of retries before giving up.
    * REDIS_MIN_RETRY_BACKOFF
    Minimum backoff between each retry.
    * REDIS_MAX_RETRY_BACKOFF
    Maximum backoff between each retry.
    * REDIS_DIAL_TIMEOUT
    Dial timeout for establishing new connections.
    * REDIS_READ_TIMEOUT
    Timeout for socket reads. If reached, commands will fail with a timeout instead of blocking.
    * REDIS_WRITE_TIMEOUT
    Timeout for socket writes. If reached, commands will fail with a timeout instead of blocking.
    * REDIS_POOL_SIZE
    Maximum number of socket connections.
    * REDIS_MIN_IDLE_CONNS
    Minimum number of idle connections which is useful when establishing new connection is slow.
    * REDIS_MAX_CONN_AGE
    Connection age at which client retires (closes) the connection.
    * REDIS_POOL_TIMEOUT
    Amount of time client waits for connection if all connections are busy before returning an error.
    * REDIS_IDLE_TIMEOUT
    Amount of time after which client closes idle connections.
    Should be less than server's timeout.
    * REDIS_IDLE_CHECK_FREQUENCY
    Frequency of idle checks made by idle connections reaper, but idle connections are still discarded by the client if Idle timeout is set.
    * REDIS_CLIENT_CRT_FILEPATH
    Crt certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * REDIS_CLIENT_KEY_FILEPATH
    Key certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * REDIS_CLIENT_CA_FILEPATH
    CA certificate file path, if your Redis needs SSL to access.  
    NOTE: Make sure to map the file to Docker volume map.
    * LOG_DIRPATH
    Log folder.
