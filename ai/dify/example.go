/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2025-06-01 20:47:38
 * @LastEditTime: 2025-06-01 20:47:43
 * @FilePath: /k-infra/ai/dify/example.go
 * @description: Dify workflow API 使用示例
 */
package dify

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Example 演示如何使用 Dify workflow API
func Example() {
	// 1. 创建客户端
	config := WorkflowConfig{
		BaseURL: "https://api.dify.ai", // 替换为你的 Dify API 地址
		APIKey:  "app-xxx",             // 替换为你的 API 密钥
		Timeout: 60 * time.Second,      // 设置超时时间
	}

	client := NewWorkflowClient(config)

	// 2. 执行 workflow
	ctx := context.Background()
	executeReq := ExecuteWorkflowRequest{
		Inputs: map[string]interface{}{
			"query":    "帮我写一篇关于人工智能的文章",
			"language": "zh-CN",
		},
		ResponseMode: ResponseModeBlocking, // 使用阻塞模式
		User:         "user-123",
		Files: []WorkflowFile{
			{
				Type:           "image",
				TransferMethod: "remote_url",
				URL:            "https://example.com/image.jpg",
			},
		},
	}

	fmt.Println("正在执行 workflow...")
	resp, err := client.ExecuteWorkflow(ctx, executeReq)
	if err != nil {
		log.Fatalf("执行 workflow 失败: %v", err)
	}

	fmt.Printf("Workflow 执行结果:\n")
	fmt.Printf("- Workflow Run ID: %s\n", resp.WorkflowRunID)
	fmt.Printf("- Task ID: %s\n", resp.TaskID)
	fmt.Printf("- 状态: %s\n", resp.Data.Status)

	// 3. 如果是异步执行，可以轮询获取结果
	if !resp.Data.IsCompleted() {
		fmt.Println("Workflow 正在运行中，等待完成...")
		workflowRunID := resp.WorkflowRunID

		// 轮询直到完成
		for {
			detail, err := client.GetWorkflowRunDetail(ctx, workflowRunID)
			if err != nil {
				log.Printf("获取运行详情失败: %v", err)
				break
			}

			fmt.Printf("当前状态: %s\n", detail.Status)

			if detail.IsCompleted() {
				if detail.IsSuccessful() {
					fmt.Println("Workflow 执行成功!")
					fmt.Printf("输出结果: %+v\n", detail.Outputs)
					fmt.Printf("耗时: %.2f 秒\n", detail.ElapsedTime)
					fmt.Printf("使用 tokens: %d\n", detail.TotalTokens)
				} else {
					fmt.Printf("Workflow 执行失败: %s\n", detail.Error)
				}
				break
			}

			// 等待一段时间再查询
			time.Sleep(2 * time.Second)
		}

		// 4. 获取详细日志
		fmt.Println("\n获取执行日志...")
		logs, err := client.GetWorkflowLogs(ctx, workflowRunID)
		if err != nil {
			log.Printf("获取日志失败: %v", err)
		} else {
			fmt.Printf("共有 %d 条日志记录:\n", len(logs.Data))
			for i, entry := range logs.Data {
				fmt.Printf("步骤 %d: %s (%s) - 状态: %s\n",
					i+1, entry.Title, entry.NodeType, entry.Status)
				if entry.Error != "" {
					fmt.Printf("  错误: %s\n", entry.Error)
				}
			}
		}
	}
}

// ExampleStreamingMode 演示流式模式的使用
func ExampleStreamingMode() {
	config := WorkflowConfig{
		BaseURL: "https://api.dify.ai",
		APIKey:  "app-xxx",
		Timeout: 60 * time.Second,
	}

	client := NewWorkflowClient(config)

	// 使用流式模式执行
	executeReq := ExecuteWorkflowRequest{
		Inputs: map[string]interface{}{
			"prompt": "生成一个创意故事",
		},
		ResponseMode: ResponseModeStreaming, // 使用流式模式
		User:         "user-123",
	}

	ctx := context.Background()
	resp, err := client.ExecuteWorkflow(ctx, executeReq)
	if err != nil {
		log.Fatalf("执行 workflow 失败: %v", err)
	}

	fmt.Printf("开始流式执行，Task ID: %s\n", resp.TaskID)

	// 在流式模式下，你需要监听 SSE 事件或定期轮询状态
	// 这里简化为定期轮询
	workflowRunID := resp.WorkflowRunID
	for {
		detail, err := client.GetWorkflowRunDetail(ctx, workflowRunID)
		if err != nil {
			log.Printf("获取状态失败: %v", err)
			break
		}

		fmt.Printf("当前状态: %s\n", detail.Status)

		if detail.IsCompleted() {
			fmt.Println("流式执行完成")
			break
		}

		time.Sleep(1 * time.Second)
	}
}

// ExampleErrorHandling 演示错误处理
func ExampleErrorHandling() {
	config := WorkflowConfig{
		BaseURL: "https://api.dify.ai",
		APIKey:  "invalid-key", // 故意使用无效的 API 密钥
		Timeout: 10 * time.Second,
	}

	client := NewWorkflowClient(config)

	executeReq := ExecuteWorkflowRequest{
		Inputs: map[string]interface{}{
			"test": "error handling",
		},
		ResponseMode: ResponseModeBlocking,
		User:         "user-123",
	}

	ctx := context.Background()
	_, err := client.ExecuteWorkflow(ctx, executeReq)
	if err != nil {
		// 检查是否是 API 错误
		if apiErr, ok := err.(*APIError); ok {
			fmt.Printf("API 错误: 代码=%s, 消息=%s, 状态=%d\n",
				apiErr.Code, apiErr.Message, apiErr.Status)
		} else {
			fmt.Printf("其他错误: %v\n", err)
		}
	}
}

// ExampleTaskManagement 演示任务管理功能
func ExampleTaskManagement() {
	config := WorkflowConfig{
		BaseURL: "https://api.dify.ai",
		APIKey:  "app-xxx",
		Timeout: 60 * time.Second,
	}

	client := NewWorkflowClient(config)

	// 启动一个长时间运行的任务
	executeReq := ExecuteWorkflowRequest{
		Inputs: map[string]interface{}{
			"task": "long running task",
		},
		ResponseMode: ResponseModeStreaming,
		User:         "user-123",
	}

	ctx := context.Background()
	resp, err := client.ExecuteWorkflow(ctx, executeReq)
	if err != nil {
		log.Fatalf("执行 workflow 失败: %v", err)
	}

	taskID := resp.TaskID
	fmt.Printf("任务已启动，Task ID: %s\n", taskID)

	// 等待一段时间后停止任务
	time.Sleep(5 * time.Second)

	fmt.Println("停止任务...")
	if err := client.StopWorkflowTask(ctx, taskID); err != nil {
		log.Printf("停止任务失败: %v", err)
	} else {
		fmt.Println("任务已停止")
	}

	// 验证任务状态
	detail, err := client.GetWorkflowRunDetail(ctx, resp.WorkflowRunID)
	if err != nil {
		log.Printf("获取任务状态失败: %v", err)
	} else {
		fmt.Printf("最终状态: %s\n", detail.Status)
	}
}
