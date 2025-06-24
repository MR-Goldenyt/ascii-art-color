#!/bin/bash

echo "🧪 Testing all cases for ascii-art-color project..."
echo "======================================================"

run_test() {
    echo -e "\\n🔹 Test: $1"
    echo "Command: $2"
    eval $2
    echo "------------------------------------------------------"
}

# ===========================
echo -e "\\n🎨 Testing BASIC COLORS..."
# ===========================
run_test "Color whole string in red" 'go run . --color=red "hello world"'
run_test "Color numbers in green" 'go run . --color=green "1 + 1 = 2"'
run_test "Color special characters in yellow" 'go run . --color=yellow "(%&) ??"'
run_test "Color substring 'ello worl'" 'go run . --color=blue "ello worl" "hello world"'
run_test "Color letter 'e'" 'go run . --color=cyan e "hello world"'
run_test "Color 2 letters 'Gu'" 'go run . --color=magenta Gu "Hey GuYs"'
run_test "Case-sensitive match: 'GuYs'" 'go run . --color=orange GuYs "HeY GuYs"'
run_test "Color letter B in 'RGB()'" 'go run . --color=blue B "RGB()"'
run_test "Random string - lower/upper/case/number" 'go run . --color=green "H3llO W0rld 123"'
run_test "Random string - special characters + one letter" 'go run . --color=yellow A "@!$#A&*("'

# ===========================
echo -e "\\n🌈 Testing EXTENDED COLORS..."
# ===========================
run_test "Test orange" 'go run . --color=orange "orange"'
run_test "Test purple" 'go run . --color=purple "purple"'
run_test "Test pink" 'go run . --color=pink "pink"'
run_test "Test gray" 'go run . --color=gray "gray"'

# ===========================
echo -e "\\n💡 Testing LIGHT COLORS..."
# ===========================
run_test "Test lightred" 'go run . --color=lightred "light red"'
run_test "Test lightgreen" 'go run . --color=lightgreen "light green"'
run_test "Test lightyellow" 'go run . --color=lightyellow "light yellow"'
run_test "Test lightblue" 'go run . --color=lightblue "light blue"'
run_test "Test lightmagenta" 'go run . --color=lightmagenta "light magenta"'
run_test "Test lightcyan" 'go run . --color=lightcyan "light cyan"'
run_test "Test lightwhite" 'go run . --color=lightwhite "light white"'

# ===========================
echo -e "\\n✨ Testing BRIGHT COLORS..."
# ===========================
run_test "Test brightblack" 'go run . --color=brightblack "bright black"'
run_test "Test brightred" 'go run . --color=brightred "bright red"'
run_test "Test brightgreen" 'go run . --color=brightgreen "bright green"'
run_test "Test brightyellow" 'go run . --color=brightyellow "bright yellow"'
run_test "Test brightblue" 'go run . --color=brightblue "bright blue"'
run_test "Test brightmagenta" 'go run . --color=brightmagenta "bright magenta"'
run_test "Test brightcyan" 'go run . --color=brightcyan "bright cyan"'
run_test "Test brightwhite" 'go run . --color=brightwhite "bright white"'

# ===========================
echo -e "\\n🖋️ Testing TEXT STYLES..."
# ===========================
run_test "Test bold style" 'go run . --color=bold "bold"'
run_test "Test underline style" 'go run . --color=underline "underline"'
run_test "Test italic style" 'go run . --color=italic "italic"'
run_test "Test strikethrough style" 'go run . --color=strikethrough "strike"'
run_test "Test blink style" 'go run . --color=blink "blinking text"'
