@echo off
setlocal enabledelayedexpansion

set "prefix=prefix_to_remove_"

for %%F in (*%prefix%*) do (
    set "filename=%%~nF"
    set "newfilename=!filename:%prefix%=!"
    ren "%%F" "!newfilename!%%~xF"
)

echo Prefix removal completed.
pause
