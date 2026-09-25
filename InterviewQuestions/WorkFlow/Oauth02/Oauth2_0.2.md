
```
                 Google
                   │
                   │ OAuth2
                   ▼
Client ────────► Backend
                   │
                   │ Authorization Code
                   ▼
                 Google
                   │
                   │ Access Token / UserInfo
                   ▼
              Backend
                   │
                   ▼
          OAuthUserInfo
                   │
                   ▼
       Find OAuth Account
                   │
          ┌────────┴────────┐
          │                 │
        Exists            Not exists
          │                 │
          ▼                 ▼
       Login          Create User
                            │
                            ▼
                     Create OAuth Account
                            │
                            ▼
                       Login User
                            │
                            ▼
                    Access/Refresh Token
```