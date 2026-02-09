/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2025-06-01 20:47:38
 * @LastEditTime: 2025-06-01 21:56:24
 * @FilePath: /k-infra/ai/dify/workflow.go
 * @description: Dify workflow API 客户端实现
 */
package dify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// WorkflowClient Dify workflow API 客户端
type WorkflowClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// WorkflowConfig 客户端配置
type WorkflowConfig struct {
	BaseURL string        // API 基础 URL
	APIKey  string        // API 密钥
	Timeout time.Duration // 请求超时时间
}

// ResponseMode 响应模式
type ResponseMode string

const (
	ResponseModeBlocking  ResponseMode = "blocking"  // 阻塞模式
	ResponseModeStreaming ResponseMode = "streaming" // 流式模式
)

// WorkflowStatus workflow 执行状态
type WorkflowStatus string

const (
	StatusRunning   WorkflowStatus = "running"   // 运行中
	StatusSucceeded WorkflowStatus = "succeeded" // 成功
	StatusFailed    WorkflowStatus = "failed"    // 失败
	StatusStopped   WorkflowStatus = "stopped"   // 已停止
)

// ExecuteWorkflowRequest 执行 workflow 请求
type ExecuteWorkflowRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`          // 输入变量
	ResponseMode ResponseMode           `json:"response_mode"`   // 响应模式
	User         string                 `json:"user"`            // 用户标识
	Files        []WorkflowFile         `json:"files,omitempty"` // 文件列表（可选）
}

// WorkflowFile 文件结构
type WorkflowFile struct {
	Type           string `json:"type"`                     // 文件类型，目前只支持 "image"
	TransferMethod string `json:"transfer_method"`          // 传输方式: "remote_url" 或 "local_file"
	URL            string `json:"url,omitempty"`            // 图片 URL（remote_url 模式）
	UploadFileID   string `json:"upload_file_id,omitempty"` // 上传文件 ID（local_file 模式）
}

// WorkflowExecutionResponse workflow 执行响应
type WorkflowExecutionResponse struct {
	WorkflowRunID string            `json:"workflow_run_id"` // workflow 运行 ID
	TaskID        string            `json:"task_id"`         // 任务 ID
	Data          WorkflowRunDetail `json:"data"`            // 执行详情
}

// WorkflowRunDetail workflow 运行详情
type WorkflowRunDetail struct {
	ID          string                 `json:"id"`           // workflow 运行 ID
	WorkflowID  string                 `json:"workflow_id"`  // 关联的 workflow ID
	Status      WorkflowStatus         `json:"status"`       // 执行状态
	Inputs      string                 `json:"inputs"`       // 输入内容
	Outputs     map[string]interface{} `json:"outputs"`      // 输出内容
	Error       string                 `json:"error"`        // 错误信息
	TotalSteps  int                    `json:"total_steps"`  // 总步数
	TotalTokens int                    `json:"total_tokens"` // 使用的总 token 数
	CreatedAt   int64                  `json:"created_at"`   // 创建时间戳
	FinishedAt  int64                  `json:"finished_at"`  // 完成时间戳
	ElapsedTime float64                `json:"elapsed_time"` // 执行耗时（秒）
}

// WorkflowLogEntry workflow 日志条目
type WorkflowLogEntry struct {
	ID                string                   `json:"id"`                            // 日志 ID
	WorkflowID        string                   `json:"workflow_id"`                   // workflow ID
	NodeID            string                   `json:"node_id"`                       // 节点 ID
	NodeType          string                   `json:"node_type"`                     // 节点类型
	Title             string                   `json:"title"`                         // 节点标题
	Index             int                      `json:"index"`                         // 执行序号
	PredecessorNodeID string                   `json:"predecessor_node_id,omitempty"` // 前置节点 ID
	Inputs            []map[string]interface{} `json:"inputs"`                        // 输入数据
	ProcessData       map[string]interface{}   `json:"process_data,omitempty"`        // 处理数据
	Outputs           map[string]interface{}   `json:"outputs,omitempty"`             // 输出数据
	Status            WorkflowStatus           `json:"status"`                        // 执行状态
	Error             string                   `json:"error,omitempty"`               // 错误信息
	ElapsedTime       float64                  `json:"elapsed_time"`                  // 执行耗时
	ExecutionMetadata map[string]interface{}   `json:"execution_metadata,omitempty"`  // 执行元数据
	CreatedAt         int64                    `json:"created_at"`                    // 创建时间
	FinishedAt        int64                    `json:"finished_at,omitempty"`         // 完成时间
}

// WorkflowLogsResponse workflow 日志响应
type WorkflowLogsResponse struct {
	Data []WorkflowLogEntry `json:"data"` // 日志列表
}

// APIError API 错误响应
type APIError struct {
	Code    string `json:"code"`    // 错误代码
	Message string `json:"message"` // 错误消息
	Status  int    `json:"status"`  // HTTP 状态码
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Dify API error [%d]: %s (%s)", e.Status, e.Message, e.Code)
}

// NewWorkflowClient 创建新的 workflow 客户端
func NewWorkflowClient(config WorkflowConfig) *WorkflowClient {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second // 默认超时时间
	}

	return &WorkflowClient{
		baseURL: config.BaseURL,
		apiKey:  config.APIKey,
		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// ExecuteWorkflow 执行 workflow
func (c *WorkflowClient) ExecuteWorkflow(ctx context.Context, req ExecuteWorkflowRequest) (*WorkflowExecutionResponse, error) {
	url := fmt.Sprintf("%s/v1/workflows/run", c.baseURL)

	// 如果没有指定响应模式，默认使用阻塞模式
	if req.ResponseMode == "" {
		req.ResponseMode = ResponseModeBlocking
	}

	// 序列化请求体
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	c.setHeaders(httpReq)

	// 发送请求
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 处理响应
	return c.handleWorkflowResponse(resp)
}

// GetWorkflowRunDetail 获取 workflow 运行详情
func (c *WorkflowClient) GetWorkflowRunDetail(ctx context.Context, workflowRunID string) (*WorkflowRunDetail, error) {
	if workflowRunID == "" {
		return nil, fmt.Errorf("workflow run ID 不能为空")
	}

	url := fmt.Sprintf("%s/v1/workflows/run/%s", c.baseURL, workflowRunID)

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	c.setHeaders(httpReq)

	// 发送请求
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, c.handleErrorResponse(resp)
	}

	// 解析响应
	var detail WorkflowRunDetail
	if err := json.NewDecoder(resp.Body).Decode(&detail); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &detail, nil
}

// StopWorkflowTask 停止 workflow 任务
func (c *WorkflowClient) StopWorkflowTask(ctx context.Context, taskID string) error {
	if taskID == "" {
		return fmt.Errorf("task ID 不能为空")
	}

	url := fmt.Sprintf("%s/v1/workflows/tasks/%s/stop", c.baseURL, taskID)

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	c.setHeaders(httpReq)

	// 发送请求
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return c.handleErrorResponse(resp)
	}

	return nil
}

// GetWorkflowLogs 获取 workflow 日志
func (c *WorkflowClient) GetWorkflowLogs(ctx context.Context, workflowRunID string) (*WorkflowLogsResponse, error) {
	if workflowRunID == "" {
		return nil, fmt.Errorf("workflow run ID 不能为空")
	}

	url := fmt.Sprintf("%s/v1/workflows/run/%s/logs", c.baseURL, workflowRunID)

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	c.setHeaders(httpReq)

	// 发送请求
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, c.handleErrorResponse(resp)
	}

	// 解析响应
	var logs WorkflowLogsResponse
	if err := json.NewDecoder(resp.Body).Decode(&logs); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &logs, nil
}

// setHeaders 设置请求头
func (c *WorkflowClient) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "k-infra-dify-client/1.0")
}

// handleWorkflowResponse 处理 workflow 执行响应
func (c *WorkflowClient) handleWorkflowResponse(resp *http.Response) (*WorkflowExecutionResponse, error) {
	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, c.handleErrorResponse(resp)
	}

	// 解析响应
	var execResp WorkflowExecutionResponse
	if err := json.NewDecoder(resp.Body).Decode(&execResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &execResp, nil
}

// handleErrorResponse 处理错误响应
func (c *WorkflowClient) handleErrorResponse(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取错误响应失败: %w", err)
	}

	// 尝试解析为 API 错误格式
	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err == nil && apiErr.Code != "" {
		apiErr.Status = resp.StatusCode
		return &apiErr
	}

	// 如果解析失败，返回通用错误
	return fmt.Errorf("API 请求失败 [%d]: %s", resp.StatusCode, string(body))
}

// IsRunning 检查 workflow 是否正在运行
func (w *WorkflowRunDetail) IsRunning() bool {
	return w.Status == StatusRunning
}

// IsCompleted 检查 workflow 是否已完成（成功或失败）
func (w *WorkflowRunDetail) IsCompleted() bool {
	return w.Status == StatusSucceeded || w.Status == StatusFailed
}

// IsSuccessful 检查 workflow 是否成功完成
func (w *WorkflowRunDetail) IsSuccessful() bool {
	return w.Status == StatusSucceeded
}

// HasError 检查 workflow 是否有错误
func (w *WorkflowRunDetail) HasError() bool {
	return w.Error != "" || w.Status == StatusFailed
}
