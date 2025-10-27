## ⚙️ **Finished Features**

* [X] **Log levels**: `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`
* [X] **Configurable log level** (e.g. only show INFO)
* [X] **Log formatting system** (text, JSON, custom template)
* [X] **Timestamp in each log**
* [X] **Include filename and line number of caller**
* [X] **Console output**
* [X] **File output** (basic)
* [X] **Custom field injection** (like request_id, trace_id)
* [X] **Colored console output**
* [X] **Structured fields** (key-value pairs)
* [X] **Context-based logging**
* [X] **Custom hooks** (before/after log handlers)
* [X] **Multiple outputs** (console + file + remote)
* [X] **Config-driven setup**

## 🧱 **UnderWork Features**

* [ ] **Asynchronous logging** (background writer with channel buffer)
* [ ] **File rotation** by size or date
* [ ] **Old log cleanup** (max days or max size)
* [ ] **Multiple outputs** (console + file + remote)
* [ ] **Custom log format per output**
* [ ] **Error handling in logging itself** (e.g., write errors)
* [ ] **Log filtering** (by module name, tag, or keyword)
* [ ] **Graceful shutdown** (flush buffered logs before exit)

---

## 🚀 **Future Features**

* [ ] **Dynamic configuration reload** (change log level/output at runtime)
* [ ] Thread-safe writes (mutex or channel)
* [ ] **Remote log streaming** (HTTP, syslog, gRPC, etc.)
* [ ] **Compression for rotated files (gzip)**
* [ ] **Log sampling / rate limiting** (drop spammy logs)
* [ ] **Buffered writer with size/time flush intervals**
* [ ] **Log message deduplication**
* [ ] **Log redaction** (mask sensitive data)
* [ ] **Custom field injection** (like request_id, trace_id)
* [ ] **Integration with tracing (OpenTelemetry)**
* [ ] **Integration with metrics (Prometheus)**
* [ ] **Panic recovery logger**
* [ ] **Cross-platform path support (Windows/Linux/Mac)**
* [ ] **Binary log format** (compact, ultra-fast)
* [ ] **Pluggable backends:**

  * File
  * ElasticSearch
  * Loki
  * Kafka
  * FluentBit / Vector
  * SQLite or custom DB
* [ ] **Structured JSON with schemas (for machine parsing)**
* [ ] **Log viewer CLI** (filter, tail, search logs)
* [ ] **Log viewer web dashboard**
* [ ] **Encryption for logs at rest**
* [ ] **Replay / audit trail mode (append-only, tamper-proof)**
* [ ] **Correlation IDs** for distributed tracing
* [ ] **Plugin system for custom formatters or writers**
* [ ] **Performance benchmarking mode**
* [ ] **Profiling hooks (measure log throughput, latency)**
