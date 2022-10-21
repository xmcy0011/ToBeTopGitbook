# jwt

## login

```bash
curl -X POST http://localhost:8080/login \
-d '{"username":"admin","password":"123","id":3}'
```

```json
{
    "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjMzMWI3YTJjLTY2NWEtNDFmOS05YjUwLWQzYzAwNjE2ZTdlMiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY2MTE1NjkwNiwidXNlcl9pZCI6MX0.ddwa_Bbx-ueGqYQcv4NWgt0R_k5JCvyMQmfcBIUN7xM","refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE3NjA4MDYsInJlZnJlc2hfdXVpZCI6IjhkMDhlMmU1LTIxMzMtNDM2NS1hOGE0LWQ1ZTNhMDk4OWMwOSIsInVzZXJfaWQiOjF9.MS2slR6IoF2FX7gtG69k-It31dVGtB5rZSArbAq_C0E"
}
```

## todo

```bash
curl -X POST http://localhost:8080/todo \
-d '{"user_id":1,"title":"this is todo"}' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImE5MGEwN2Y5LTE0OTgtNDc3YS1hNDdhLWMyY2FiNWQ3NzE3NiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY2MTMxNDgwMSwidXNlcl9pZCI6MX0.6aImjI_ATsuZ9ASdoDetFCBMQpL3wIpnjxeQcJhyexU'
```
