
### DEMO USE

`go build main.go`


`./main`

fill the domain, mailgun api and recipient in main.go


### Example

```bash
./main \
-mailgun-api-key="xxxx" \
-domain="xxx.xx.com" \
-noreply="Tim <robot@xxx.xx.com>" \
-recipient="target@xxx.xx.com" \
-attachment-path="/xx/xx" \
-subject="Hello, I am sender" \
-body="You are good boy. right?"
```
