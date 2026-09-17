
Cấu trúc thư mục
```

│   README.md
│
├───0.HTTP_2
│       HTTP2 Fundamentals.md          # gồm: Connection, Stream, Stream ID, Multiplexing
│       Binary framing & Header compression.md   
│
├───1.Protocol Buffers
│       Protobuf Basics.md             # gồm: Introduction, Syntax, Scalar types, Enum
│       Protobuf Advanced Types.md     # gồm: Nested message, Repeated/Map, Oneof, Default/Optional
│       Service Definition & Compiler.md  # gồm: rpc keyword, protoc, generate code
│
├───2.gRPC Fundamentals
│       gRPC Overview.md               # gồm: What is gRPC, Why gRPC vs REST, Architecture
│       Setup & First Project.md       # gồm: cài môi trường, generate code, cấu trúc project
│
├───3.4 Types of RPC
│       Unary RPC.md
│       Streaming RPC.md               # gồm cả 3 loại: Server/Client/Bidirectional streaming
│                                         (vì cùng bản chất "stream", học chung dễ so sánh)
│
├───4.Context - Metadata - Error
│       Context, Deadline & Cancellation.md   
│       Metadata & Status Codes.md            # gồm: metadata + error handling
│
├───5.Interceptor
│       Interceptor Overview.md        # gồm: Unary/Stream, Client/Server-side, ví dụ logging
│
└───6.Authentication & Authorization
        Auth Overview.md               # gồm: TLS/SSL, mTLS, JWT, per-RPC credentials
```

--- 
Phần nội dung thực hành [Golang + gRPC]([https://github.com/GiaBao0510/ExperienceLearnedAndGathered/tree/main/Go/GO_web/RPC_gRPC)
