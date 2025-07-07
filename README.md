# Week 3 - Real-time Chat App (Go + WebSocket + Redis)

Ứng dụng chat thời gian thực sử dụng WebSocket và Redis.

## ✅ Tính năng
- Nhắn tin realtime giữa nhiều client
- Hiển thị user đang online
- Lưu lịch sử tin nhắn bằng Redis
- Giới hạn tốc độ gửi để chống spam

## 🧰 Công nghệ
- [Gin](https://github.com/gin-gonic/gin): router nhẹ
- [gorilla/websocket](https://github.com/gorilla/websocket): giao tiếp realtime
- [Redis](https://redis.io): lưu online & message history
- [rate limiter](https://pkg.go.dev/golang.org/x/time/rate): giới hạn spam

## 📡 WebSocket
- Kết nối WebSocket:
ws://localhost:8080/ws/{username}
- Gửi tin nhắn:
```json
{ "content": "hello" }
