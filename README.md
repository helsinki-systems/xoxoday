# xoxoday API client

```go
// Generate a token and OAuth tokens here: https://stores.xoxoday.com/admin/accounts/platform-preferences/
const t = `{"success":1,"status":200,"access_token":"eyJ...`

var token xoxoday.Token
if err := json.Unmarshal([]byte(t), &token); err != nil {
    return fmt.Errorf("failed to parse token: %w", err)
}

api := xoxoday.New(
    context.Background(),
    xoxoday.EnvDevelopment,
    token,
    xoxoday.OAuthConfig{
        ClientID:     "As generated on the URL above",
        ClientSecret: "As generated on the URL above",
    },
)

res, err := api.Balance()
if err != nil {
    return fmt.Errorf("failed to get balance: %w", err)
}

fmt.Printf("Balance: %f%s\n", res.Points, res.Currency)
```
