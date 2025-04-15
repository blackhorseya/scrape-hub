# GitHub Copilot Custom Instructions

本專案以 Go 語言實作，架構採用 Clean Architecture 與 Domain-Driven Design（DDD），撰寫風格遵循官方文件 [Effective Go](https://go.dev/doc/effective_go)。以下是團隊約定，請 Copilot 回覆符合這些慣例的建議與程式碼。

---

## 語言與撰寫風格

- 使用 Go 1.21+
- 所有程式碼必須使用 `gofmt` 格式化，縮排使用 tab
- 命名使用 `MixedCaps`，避免使用底線（如 `myVar` 而非 `my_var`）
- Getter 方法不加 `Get` 前綴：使用 `User()` 而非 `GetUser()`
- Interface 命名慣例：`-er` 結尾（如 `Fetcher`, `Storer`, `Writer`）
- 多回傳值中的錯誤應命名為 `err`，早期返回（early return）
- 善用 `_` 忽略不需要的回傳值
- 控制結構 (`if`, `for`, `switch`) 不加括號
- 測試檔案應命名為 `_test.go`

---

## 架構設計原則

- 採用 Clean Architecture：
  - `domain` 層：純商業邏輯、entity、value object、interface，**不得依賴外部套件**
  - `application` 層：執行 use case，調用 domain 及 infrastructure
  - `infrastructure` 層：處理 GORM、Redis、API 等外部資源
  - `interfaces` 層：HTTP handler、gRPC adapter 等對外介面
- 採用 Domain-Driven Design（DDD）：
  - 明確聚合邊界，實體（entity）封裝狀態與行為
  - 避免貧血模型，不將邏輯散落於 handler 或 repository 實作

---

## 檔案與命名慣例

- 實作 struct 後綴使用 `Impl`（如 `UserRepositoryImpl`）
- 實作檔案使用 `_impl.go` 命名（如 `user_repository_impl.go`）
- 所有 public symbol 必須有 doc comment，以 `XXX ...` 開頭

---

## Context 使用規範

- 定義 function 時，請傳入原生 `context.Context`，變數名稱固定為：`c`
- 在 function 實作中，若需 logger 或擴展 context 功能，起手式一律為：

```go
ctx := contextx.WithContext(c)
```

- 只有擴展過的 context（如 `contextx.Contextx`）才命名為 `ctx`
- 所有 logging 一律使用 `ctx.Logger`（來自 `contextx`）

---

## 依賴注入與 Mock 工具

- **依賴注入（DI）**：使用 [`google/wire`](https://github.com/google/wire)
  - 每個功能模組需建立 wire provider
- **Mock 工具**：僅使用 [`uber-go/mock`](https://github.com/uber-go/mock)
  - **禁用 golang/mock（已棄用）**
  - mock interface 應定義於 `mock_*.go` 中，並由 wire 注入

---

## ORM 使用原則

- 可使用 GORM，但**僅限 infrastructure 層**
- domain 層禁止出現 GORM tag 或 struct，僅定義 interface
- infra 層的實作可包含 GORM model struct 與 tag
- 優先使用 interface + 明確 SQL 實作，而非過度依賴 ORM 魔法

---

## MCP（Model Context Protocol）工具支援

- 專案已整合 MCP server，提供 metadata 查詢、工具呼叫與自動轉換功能
- Copilot 回覆中，若涉及以下狀況請主動建議使用 MCP：
  - 跨模組協作
  - 需要查詢結構、mapping 或資料轉換
  - 需封裝成工具或 helper

---

## Copilot 回應格式建議

請 Copilot 回應時遵循以下風格：

- 使用 idiomatic Go，符合 [Effective Go](https://go.dev/doc/effective_go)
- 函式參數中的 context 請命名為 `c`，並以 `ctx := contextx.WithContext(c)` 起手
- 結構清晰、可測試，優先提供 interface 與起手實作
- 若設計為 infrastructure 實作，請將 struct 命名為 `XXXImpl`，檔名為 `xxx_impl.go`
- 提及擴展功能或服務時，建議使用 MCP 工具，並說明使用情境
