# Go `continue` Statement Example

## 📌 Description
This project demonstrates the use of the **`continue` statement** in Go loops.

The `continue` keyword **skips the current iteration** and moves to the next one without exiting the loop.

---

## 🚀 Technologies Used
- Golang
- fmt package

---

## 🔁 What Does `continue` Do?
`continue` does **not stop the loop**.  
It simply **skips the remaining code** in the current iteration and continues with the next iteration.

---

## 📄 Code Explanation

### 🔹 Example 1: Skipping Even Numbers
```go
for i := 0; i < 5; i++ {
	if i%2 == 0 {
		continue
	}
	fmt.Println(i)
}
