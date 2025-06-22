# ascii-art-color

## 🎯 Description

`ascii-art-color` is a CLI tool in Go that prints colored ASCII art using a banner (`standard.txt`).  
You can color either the **whole string** or a specific **substring** using a `--color=<color>` flag.

---

## 🚀 Usage

```bash
go run . --color=<color> <substring> "Your text here"
```

### Examples

```bash
go run . --color=red kit "a king kitten has a kit"
```

⟶ This will print the ASCII art and color every `kit` in red.

```bash
go run . --color=cyan "hello world"
```

⟶ This colors the whole string if no substring is provided.

---

## 🎨 Supported Colors

- black
- red
- green
- yellow
- blue
- magenta
- cyan
- white

---

## ⚠️ Invalid Usage

If the command is written incorrectly, this message will be shown:

```bash
Usage: go run . --color=<color> <substring> "something"
```

---

## 👥 Team Members

- 👑 **HUSSAIN ALI** – Team Leader [`@hussainali7`](https://learn.reboot01.com/git/root/hussainali7)
- 🛠️ **Ahmed Alsafseef** – Developer [`@aalsafse`](https://learn.reboot01.com/git/root/aalsafse)
- 🧩 **Abdulla Ashoor** – Developer [`@abashoor`](https://learn.reboot01.com/git/root/abashoor)

---

## 📂 Project Structure

```
ascii-art-color/
├── main.go
├── standard.txt
├── helpers/
│   ├── color.go
│   └── SplitSegments.go
```