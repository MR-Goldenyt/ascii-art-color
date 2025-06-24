# 🎨 ascii-art-color

## 💡 Description

`ascii-art-color` is a CLI tool written in Go that prints **ASCII art** from a `standard.txt` font file, with support for **coloring** specific substrings or the entire string using a `--color=<color>` flag.

---

## 🚀 Usage

```bash
go run . --color=<color> [substring] "your string here"
```

- If you provide a substring, only matching parts will be colored.
- If no substring is provided, the entire string will be colored.



---

## ✨ Examples

```bash
go run . --color=red "hello world"
```
⟶ Colors the entire ASCII output in **red**.

```bash
go run . --color=orange kit "a king kitten has a kit"
```
⟶ Colors every occurrence of the substring `kit` in **orange** only.

```bash
go run . --color=underline "important"
```
⟶ Applies the underline style to the entire string.

---

## 🎨 Supported Colors & Styles

### 🎨 Colors Table

| Basic Colors  | Light Colors    | Bright Colors    | Extended Colors |
|---------------|------------------|-------------------|-----------------|
| `black`       | `lightred`       | `brightblack`     | `orange`        |
| `red`         | `lightgreen`     | `brightred`       | `purple`        |
| `green`       | `lightyellow`    | `brightgreen`     | `pink`          |
| `yellow`      | `lightblue`      | `brightyellow`    | `gray`          |
| `blue`        | `lightmagenta`   | `brightblue`      |                 |
| `magenta`     | `lightcyan`      | `brightmagenta`   |                 |
| `cyan`        | `lightwhite`     | `brightcyan`      |                 |
| `white`       |                  | `brightwhite`     |                 |

### ✏️ Text Styles

| Style            | Description             |
|------------------|-------------------------|
| `bold`           | Bold text               |
| `dim`            | Dim/faint text          |
| `italic`         | Italic text             |
| `underline`      | Underlined text         |
| `blink`          | Blinking text           |
| `reverse`        | Reverse video effect    |
| `hidden`         | Hidden text             |
| `strikethrough`  | Strike-through text     |

### 🧼 Style Reset Flags

| Flag               | Effect                |
|--------------------|-----------------------|
| `reset`            | Reset all styles      |
| `nobold`           | Cancel bold           |
| `noitalic`         | Cancel italic         |
| `nounderline`      | Cancel underline      |
| `noblink`          | Cancel blinking       |
| `noreverse`        | Cancel reverse effect |
| `nohidden`         | Cancel hidden         |
| `nostrikethrough`  | Cancel strike         |

🧱 Extra Color Categories

🎨 foreground + background (fgbk)

| Color Combo       | Code Example Usage             |
| ----------------- | ------------------------------ |
| `redfgbk`         | Red text with red background   |
| `orangefgbk`      | Orange foreground + background |
| `lightcyanfgbk`   | Light cyan both fg & bg        |
| `brightwhitefgbk` | Bright white both fg & bg      |
| `grayfgbk`        | Gray both fg & bg              |


🎨 Foreground Variants (fg)

| Foreground Color |
| ---------------- |
| `redfg`          |
| `bluefg`         |
| `orangefg`       |
| `purplefg`       |
| `pinkfg`         |
| `grayfg`         |
| `lightredfg`     |
| `brightyellowfg` |

🎨 Background Variants (bg)

| Background Color |
| ---------------- |
| `redbg`          |
| `bluebg`         |
| `orangebg`       |
| `purplebg`       |
| `pinkbg`         |
| `graybg`         |
| `lightgreenbg`   |
| `brightcyanbg`   |

🖼️ Extra Reset/Cancel Styles

| Reset Flag | Description           |
| ---------- | --------------------- |
| `normal`   | Reset to normal style |
| `nobright` | Cancel bold/bright    |
| `nofaint`  | Cancel dim/faint      |




---

## ⚠️ Invalid Usage

If the usage is incorrect, the program shows:

```bash
Usage: go run . --color=<color> [substring] "something"
```

---

## 👥 Team Members

- 👑 **HUSSAIN ALI** – Team Leader [`@hussainali7`](https://learn.reboot01.com/git/root/hussainali7)
- 🛠️ **Ahmed Alsafseef** – Developer [`@aalsafse`](https://learn.reboot01.com/git/root/aalsafse)
- 🧩 **Abdulla Ashoor** – Developer [`@abashoor`](https://learn.reboot01.com/git/root/abashoor)

---

## 📁 Project Structure

```
ascii-art-color/
├── main.go
├── standard.txt
├── test.sh
├── README.md
├── helpers/
│   ├── color.go
│   └── SplitSegments.go
```

---

## ✅ Notes

- Only standard Go packages are used (no external libraries).
- The project is fully compatible with ANSI terminals.
