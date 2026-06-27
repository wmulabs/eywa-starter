# Generate the README demo GIF with VHS: https://github.com/charmbracelet/vhs
#
#   brew install vhs          # or: go install github.com/charmbracelet/vhs@latest
#   cp .env.example .env      # set a real GEMINI_API_KEY (the demo makes a live LLM call)
#   vhs assets/demo.tape      # writes assets/demo.gif
#
# Then add it to README.md:  <p align="center"><img src="assets/demo.gif" width="100%"/></p>

Output assets/demo.gif

Set Shell "bash"
Set FontSize 18
Set Width 1200
Set Height 660
Set Padding 24
Set Theme "Dracula"

Type "make up"             Enter  Sleep 4s
Type "make run &"          Enter  Sleep 7s
Type "make chat"           Enter  Sleep 8s
