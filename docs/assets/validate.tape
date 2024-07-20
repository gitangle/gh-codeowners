Output validate.gif

Set Margin 5
Set PlaybackSpeed 0.25
Set Padding 5
Set FontSize 20
Set BorderRadius 15
Set CursorBlink false
Set WindowBar Colorful

Set Width 1420
Set Height 1120

Env PROMPT "  ▲ "

Set Shell zsh
Sleep 1s
Type@100ms "gh codeowners validate --owner asyncapi --all"
Enter
Sleep 0.5s
Type "# Stay tuned for more! 🚀"
Sleep 5s

