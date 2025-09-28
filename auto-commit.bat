@echo off
:: 自动提交代码到GitHub的Windows批处理脚本
:: 使用方法: auto-commit.bat "提交信息"

:: 检查是否提供了提交信息
if "%1"=="" (
    echo 未提供提交信息，正在生成默认提交信息...
    
    :: 获取当前时间作为默认提交信息的一部分
    for /f "tokens=2 delims==" %%a in ('wmic OS Get localdatetime /value') do set "dt=%%a"
    set TIMESTAMP=%dt:~0,4%-%dt:~4,2%-%dt:~6,2% %dt:~8,2%:%dt:~10,2%:%dt:~12,2%
    
    :: 生成默认提交信息
    set COMMIT_MESSAGE=Auto commit at %TIMESTAMP%
    
    echo 使用默认提交信息: %COMMIT_MESSAGE%
) else (
    :: 使用提供的提交信息
    set COMMIT_MESSAGE=%1
)

echo 开始自动提交流程...
echo 提交信息: %COMMIT_MESSAGE%

:: 添加所有更改到暂存区（排除auto-commit.bat和block-emulator目录）
echo 正在添加文件到暂存区...
git add --all
git reset -- auto-commit.bat block-emulator

:: 检查是否有文件被添加
git diff --cached --quiet
if %ERRORLEVEL% == 0 (
    echo 没有文件需要提交
    exit /b 0
)

:: 提交更改
echo 正在提交更改...
git commit -m "%COMMIT_MESSAGE%"

:: 推送到远程仓库
echo 正在推送到远程仓库...
git push origin main

echo 代码已成功提交并推送到GitHub!

:: 显示提交历史
echo 最新的提交历史:
git log --oneline -5

pause