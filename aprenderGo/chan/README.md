# Go Channels Example

## 📌 Description
This project demonstrates the use of **channels in Golang** to safely communicate data between goroutines.

It includes an example of a **buffered channel**, showing how values can be sent and received without blocking immediately.

> Go philosophy:  
> **“Do not communicate by sharing memory; share memory by communicating.”**

---

## 🚀 Technologies Used
- Golang
- Goroutines
- Channels
- fmt package

---

## 📄 Code Explanation

### 🔗 Channels
Channels allow goroutines to **send and receive data safely**.

This example uses a **buffered channel** of size 2:
```go
ch := make(chan string, 2)
