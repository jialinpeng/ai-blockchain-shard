#!/bin/bash

# 自动提交代码到GitHub的脚本
# 使用方法: ./auto-commit.sh "提交信息"

# 检查是否提供了提交信息
if [ $# -eq 0 ]; then
    echo "未提供提交信息，正在生成默认提交信息..."
    
    # 获取当前时间作为默认提交信息的一部分
    TIMESTAMP=$(date +"%Y-%m-%d %H:%M:%S")
    
    # 获取更改的文件数量
    CHANGED_FILES=$(git diff --name-only HEAD | wc -l)
    
    # 生成默认提交信息
    COMMIT_MESSAGE="Auto commit: $CHANGED_FILES files changed at $TIMESTAMP"
    
    echo "使用默认提交信息: $COMMIT_MESSAGE"
else
    # 使用提供的提交信息
    COMMIT_MESSAGE="$1"
fi

# 检查脚本是否有执行权限
if [ ! -x "$0" ]; then
    echo "警告: 脚本没有执行权限，正在添加执行权限..."
    chmod +x "$0"
    echo "已添加执行权限，请重新运行脚本"
    exit 0
fi

echo "开始自动提交流程..."
echo "提交信息: $COMMIT_MESSAGE"

# 添加所有更改到暂存区（排除auto-commit.sh和block-emulator目录）
echo "正在添加文件到暂存区..."
git add --all
git reset -- auto-commit.sh auto-commit.bat block-emulator

# 检查是否有文件被添加
if git diff --cached --quiet; then
    echo "没有文件需要提交"
    exit 0
fi

# 提交更改
echo "正在提交更改..."
git commit -m "$COMMIT_MESSAGE"

# 推送到远程仓库
echo "正在推送到远程仓库..."
git push origin main

echo "代码已成功提交并推送到GitHub!"

# 显示提交历史
echo "最新的提交历史:"
git log --oneline -5