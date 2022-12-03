# Splunk On

Basic utility to query the Splunk management REST interface with valid credentials.

Users can pull back a bit of data but probably want `admin` permissions ([see here](https://docs.splunk.com/Documentation/InfraApp/2.2.5/Admin/AdminUserAccounts)) to get interesting data, and the management port 8089 needs to be exposed.

If you're on a local instance, you may need to specify `localhost:8089` instead of `cloudinstance.splunkcloud.com:8089`.

## Usage
`./splunkon https://instance.splunkcloud.com:8089/ username password`

JSON file with the entire (very verbose) Splunk responses logged to `splunkon.json`

```
$ ./splunkon https://localhost:8089 admin password123
2022/12/01 20:30:42 Beginning recon on https://localhost:8089

CURRENT USER:
+---+----------+----------------------+---------------+------------+--------+----+---------+
| # | USERNAME | EMAIL                | REAL NAME     | LAST LOGIN | LOCKED | TZ | ROLES   |
+---+----------+----------------------+---------------+------------+--------+----+---------+
| 0 | admin    | changeme@example.com | Administrator | 1670039682 | false  |    | [admin] |
+---+----------+----------------------+---------------+------------+--------+----+---------+

USERS:
+---+----------+----------------------+---------------+------------+--------+-----+---------+--------+
| # | USERNAME | EMAIL                | REAL NAME     | LAST LOGIN | LOCKED | TZ  | ROLES   | TYPE   |
+---+----------+----------------------+---------------+------------+--------+-----+---------+--------+
| 0 | admin    | changeme@example.com | Administrator | 1670039682 | false  |     | [admin] | Splunk |
| 1 | j_doe    |                      | John Doe      | 1670038706 | false  | GMT | [user]  | Splunk |
+---+----------+----------------------+---------------+------------+--------+-----+---------+--------+

ROLES:
+---+--------------------+--------------------+--------------+-----------------------+--------------------+
| # | NAME               | DEFAULT SEARCH IDX | CAPABILITIES | IMPORTED CAPABILITIES | ADMIN CAPABILITIES |
+---+--------------------+--------------------+--------------+-----------------------+--------------------+
| 0 | admin              | [main os]          |          104 |                    39 | true               |
| 1 | can_delete         | []                 |            6 |                     0 | false              |
| 2 | power              | [main]             |           10 |                    29 | false              |
| 3 | splunk-system-role | []                 |            0 |                   143 | true               |
| 4 | user               | [main]             |           29 |                     0 | false              |
+---+--------------------+--------------------+--------------+-----------------------+--------------------+
```

## Endpoints

Replace splunkcloud.com with localhost or your target instance

There are a ton of endpoints, feel free to add more
- https://docs.splunk.com/Documentation/Splunk/latest/RESTREF/RESTlist
- https://docs.splunk.com/Documentation/Splunk/9.0.2/RESTUM/RESTusing#Authentication_and_authorization

```
https://<cloudinstance>.splunkcloud.com:8089/services/authentication/users
https://<cloudinstance>.splunkcloud.com:8089/services/admin/system-info
https://<cloudinstance>.splunkcloud.com:8089/services/admin/server-info
https://<cloudinstance>.splunkcloud.com:8089/services/server/settings
https://<cloudinstance>.splunkcloud.com:8089/services/properties/authentication/saml
https://<cloudinstance>.splunkcloud.com:8089/services/properties/authentication/userToRoleMap_SAML
https://<cloudinstance>.splunkcloud.com:8089/services/admin/script
https://<cloudinstance>.splunkcloud.com:8089/services/admin/tokens
https://<cloudinstance>.splunkcloud.com:8089/services/authorization/roles
https://<cloudinstance>.splunkcloud.com:8089/services/properties/indexes
https://<cloudinstance>.splunkcloud.com:8089/services/properties/sourcetypes
https://<cloudinstance>.splunkcloud.com:8089/services/apps/local
https://<cloudinstance>.splunkcloud.com:8089/services/admin/savedsearch
https://<cloudinstance>.splunkcloud.com:8089/services/alerts/fired_alerts
https://<cloudinstance>.splunkcloud.com:8089/services/storage/passwords
```

## Detection
Splunk logs all the things.

```
index=_internal root=services
| stats values(uri_path) values(useragent) values(clientip) by user
```
