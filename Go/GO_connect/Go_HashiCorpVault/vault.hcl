# vault.hcl – Cấu hình Vault server cho môi trường development

# storage: dùng file system (chỉ cho DEV, không dùng cho PRODUCTION) 
storage "file" {
    path = "/vault/data"
}

#listener: vault lắng nghe trên port 8200
listener "tcp" {
    address = "127.0.0.1:8200"
    tls_disable = true # tắt TLS - chỉ chấp nhận trong .env
}

#UI web console 
ui = true

# API address (cho cluster communication)
api_address = "https://127.0.0.1:8200"

# log level: trace, debug, info, warn, error
log_level = "info"
