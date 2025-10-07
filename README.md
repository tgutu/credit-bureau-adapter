# Credit Bureau Adapter

The Credit Bureau Adapter is a service layer that provides a unified API to interact with multiple credit bureaus (e.g., Experian, Equifax, TransUnion).
It abstracts away bureau-specific integrations, formats, and authentication schemes — exposing a single, consistent interface for retrieving credit data.

## 🚀 Features

🔄 Unified Interface — One API to access multiple credit bureaus

🧠 Response Normalization — Consistent credit report structure across bureaus

⚙️ Pluggable Architecture — Easily add new bureau integrations

🧰 Error Normalization — Common error codes for predictable handling


## Architecture Overview
```
┌──────────────────────────┐
│  Loan Origination System │
└──────────┬───────────────┘
           │ REST/gRPC
           ▼
┌─────────────────────────────┐
│ Credit Bureau Adapter API   │
│  POST /v1/credit/report     │
│  GET /v1/credit/bureaus     │
└──────────┬──────────────────┘
           │
           ▼
┌──────────────────────────┐
│ Bureau Integrations      │
│  • EquifaxClient         │
│  • ExperianClient        │
│  • TransUnionClient      │
└──────────┬───────────────┘
           │
           ▼
┌──────────────────────────┐
│ External Bureau APIs     │
└──────────────────────────┘
```