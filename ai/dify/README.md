# Dify Workflow API 客户端

这是一个用 Go 语言实现的 Dify Workflow API 客户端，提供了完整的 workflow 执行和管理功能。

## 功能特性

- ✅ **执行 Workflow**: 支持阻塞和流式两种模式
- ✅ **获取运行详情**: 实时查询 workflow 执行状态
- ✅ **停止任务**: 支持停止正在运行的 workflow 任务
- ✅ **获取日志**: 查看详细的执行日志和节点状态
- ✅ **错误处理**: 完善的错误处理和类型安全
- ✅ **文件支持**: 支持图片文件上传和处理
- ✅ **超时控制**: 可配置的请求超时时间

## 安装

```bash
go get github.com/wnz27/k-infra/ai/dify
```

## 快速开始

### 1. 创建客户端

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/wnz27/k-infra/ai/dify"
)

func main() {
    // 配置客户端
    config := dify.WorkflowConfig{
        BaseURL: "https://api.dify.ai",  // 你的 Dify API 地址
        APIKey:  "app-xxxxxxxxxx",      // 你的 API 密钥
        Timeout: 60 * time.Second,      // 请求超时时间
    }

    client := dify.NewWorkflowClient(config)
}
```

### 2. 执行 Workflow（阻塞模式）

```go
func executeWorkflow(client *dify.WorkflowClient) {
    ctx := context.Background()

    req := dify.ExecuteWorkflowRequest{
        Inputs: map[string]interface{}{
            "query":    "帮我写一篇关于人工智能的文章",
            "language": "zh-CN",
        },
        ResponseMode: dify.ResponseModeBlocking, // 阻塞模式
        User:         "user-123",
    }

    resp, err := client.ExecuteWorkflow(ctx, req)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Workflow 执行完成!\n")
    fmt.Printf("Run ID: %s\n", resp.WorkflowRunID)
    fmt.Printf("状态: %s\n", resp.Data.Status)
    fmt.Printf("输出: %+v\n", resp.Data.Outputs)
}
```

### 3. 执行 Workflow（流式模式）

```go
func executeWorkflowStreaming(client *dify.WorkflowClient) {
    ctx := context.Background()

    req := dify.ExecuteWorkflowRequest{
        Inputs: map[string]interface{}{
            "prompt": "生成一个创意故事",
        },
        ResponseMode: dify.ResponseModeStreaming, // 流式模式
        User:         "user-123",
    }

    resp, err := client.ExecuteWorkflow(ctx, req)
    if err != nil {
        panic(err)
    }

    // 轮询获取结果
    workflowRunID := resp.WorkflowRunID
    for {
        detail, err := client.GetWorkflowRunDetail(ctx, workflowRunID)
        if err != nil {
            break
        }

        fmt.Printf("当前状态: %s\n", detail.Status)

        if detail.IsCompleted() {
            if detail.IsSuccessful() {
                fmt.Println("执行成功!")
                fmt.Printf("结果: %+v\n", detail.Outputs)
            } else {
                fmt.Printf("执行失败: %s\n", detail.Error)
            }
            break
        }

        time.Sleep(2 * time.Second)
    }
}
```

### 4. 支持文件上传

```go
func executeWithFiles(client *dify.WorkflowClient) {
    ctx := context.Background()

    req := dify.ExecuteWorkflowRequest{
        Inputs: map[string]interface{}{
            "description": "分析这张图片",
        },
        ResponseMode: dify.ResponseModeBlocking,
        User:         "user-123",
        Files: []dify.WorkflowFile{
            {
                Type:           "image",
                TransferMethod: "remote_url",
                URL:            "https://example.com/image.jpg",
            },
            // 或者使用本地文件
            {
                Type:           "image",
                TransferMethod: "local_file",
                UploadFileID:   "uploaded-file-id",
            },
        },
    }

    resp, err := client.ExecuteWorkflow(ctx, req)
    if err != nil {
        panic(err)
    }

    fmt.Printf("文件处理完成: %+v\n", resp.Data.Outputs)
}
```

### 5. 获取执行日志

```go
func getWorkflowLogs(client *dify.WorkflowClient, workflowRunID string) {
    ctx := context.Background()

    logs, err := client.GetWorkflowLogs(ctx, workflowRunID)
    if err != nil {
        panic(err)
    }

    fmt.Printf("共有 %d 条日志记录:\n", len(logs.Data))
    for i, entry := range logs.Data {
        fmt.Printf("步骤 %d: %s (%s) - 状态: %s\n",
            i+1, entry.Title, entry.NodeType, entry.Status)

        if entry.Error != "" {
            fmt.Printf("  错误: %s\n", entry.Error)
        }

        fmt.Printf("  耗时: %.2f 秒\n", entry.ElapsedTime)
    }
}
```

### 6. 停止运行中的任务

```go
func stopTask(client *dify.WorkflowClient, taskID string) {
    ctx := context.Background()

    err := client.StopWorkflowTask(ctx, taskID)
    if err != nil {
        fmt.Printf("停止任务失败: %v\n", err)
        return
    }

    fmt.Println("任务已成功停止")
}
```

## 错误处理

客户端提供了完善的错误处理机制：

```go
func handleErrors(client *dify.WorkflowClient) {
    ctx := context.Background()

    req := dify.ExecuteWorkflowRequest{
        Inputs:       map[string]interface{}{"test": "data"},
        ResponseMode: dify.ResponseModeBlocking,
        User:         "user-123",
    }

    _, err := client.ExecuteWorkflow(ctx, req)
    if err != nil {
        // 检查是否是 API 错误
        if apiErr, ok := err.(*dify.APIError); ok {
            fmt.Printf("API 错误:\n")
            fmt.Printf("  代码: %s\n", apiErr.Code)
            fmt.Printf("  消息: %s\n", apiErr.Message)
            fmt.Printf("  状态: %d\n", apiErr.Status)
        } else {
            fmt.Printf("其他错误: %v\n", err)
        }
    }
}
```

## 状态检查

`WorkflowRunDetail` 提供了便捷的状态检查方法：

```go
func checkStatus(detail *dify.WorkflowRunDetail) {
    if detail.IsRunning() {
        fmt.Println("Workflow 正在运行中...")
    }

    if detail.IsCompleted() {
        if detail.IsSuccessful() {
            fmt.Println("Workflow 执行成功!")
        } else {
            fmt.Println("Workflow 执行失败")
        }
    }

    if detail.HasError() {
        fmt.Printf("发现错误: %s\n", detail.Error)
    }
}
```

## API 参考

### WorkflowClient 方法

| 方法                               | 描述          |
| ---------------------------------- | ------------- |
| `ExecuteWorkflow(ctx, req)`        | 执行 workflow |
| `GetWorkflowRunDetail(ctx, runID)` | 获取运行详情  |
| `StopWorkflowTask(ctx, taskID)`    | 停止任务      |
| `GetWorkflowLogs(ctx, runID)`      | 获取执行日志  |

### 响应模式

| 模式                    | 描述     | 使用场景                 |
| ----------------------- | -------- | ------------------------ |
| `ResponseModeBlocking`  | 阻塞模式 | 简单任务，等待完整结果   |
| `ResponseModeStreaming` | 流式模式 | 长时间任务，需要实时状态 |

### 执行状态

| 状态              | 描述     |
| ----------------- | -------- |
| `StatusRunning`   | 运行中   |
| `StatusSucceeded` | 成功完成 |
| `StatusFailed`    | 执行失败 |
| `StatusStopped`   | 已停止   |

## 注意事项

1. **API 密钥**: 确保使用正确的 Dify API 密钥
2. **超时设置**: 根据 workflow 复杂度设置合适的超时时间
3. **轮询频率**: 在流式模式下，建议设置合理的轮询间隔（1-5 秒）
4. **错误重试**: 对于网络错误，建议实现重试机制
5. **并发控制**: 注意 API 的并发限制

## 示例代码

查看 [example.go](example.go) 文件获取更多完整的使用示例。

## 许可证

本项目基于 MIT 许可证开源。

## 参考链接

- [Dify 官方文档](https://docs.dify.ai/)
- [Dify Workflow API 文档](https://docs.dify.ai/api-reference/workflow-execution/)
- [Dify GitHub 仓库](https://github.com/langgenius/dify)
