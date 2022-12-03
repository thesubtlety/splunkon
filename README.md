# Splunk On

Basic REST reader allows you to query a Splunk management interface with valid credentials.

Users can pull back a bit of data but probably want `admin` permissions ([see here](https://docs.splunk.com/Documentation/InfraApp/2.2.5/Admin/AdminUserAccounts)) to get interesting data, and the management port 8089 needs to be exposed.

If you're on a local instance, you may need to specify `localhost:8089` instead of `cloudinstance.splunkcloud.com:8089`.

## Usage
`./splunkon https://instance.splunkcloud.com:8089/ username password`

JSON file with the entire (very verbose) Splunk responses logged to `splunkon.json`

## Paths

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

There are a ton of endpoints, feel free to add more
- https://docs.splunk.com/Documentation/Splunk/latest/RESTREF/RESTlist
- https://docs.splunk.com/Documentation/Splunk/9.0.2/RESTUM/RESTusing#Authentication_and_authorization

## Detection
Splunk logs all the things.

```
index=_internal root=services
| stats values(uri_path) values(useragent) values(clientip) by user
```
