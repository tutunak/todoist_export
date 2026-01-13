package todoist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tutunak/todoist_export/model"
)

const baseURL = "https://api.todoist.com/rest/v2"

type Client struct {
	Token      string
	HttpClient *http.Client
	BaseURL    string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		BaseURL: "https://api.todoist.com/rest/v2",
	}
}

func (c *Client) doRequest(method, endpoint string, params map[string]string) (*http.Response, error) {
	reqURL, err := url.Parse(c.BaseURL + endpoint)
	if err != nil {
		return nil, err
	}

	if params != nil {
		q := reqURL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		reqURL.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		resp.Body.Close()
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	return resp, nil
}

func (c *Client) GetProjects() ([]model.Project, error) {
	resp, err := c.doRequest("GET", "/projects", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var projects []model.Project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (c *Client) GetTasks(filter string) ([]model.Task, error) {
	params := make(map[string]string)
	if filter != "" {
		params["filter"] = filter
	}

	resp, err := c.doRequest("GET", "/tasks", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tasks []model.Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (c *Client) GetLabels() ([]model.Label, error) {
	resp, err := c.doRequest("GET", "/labels", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var labels []model.Label
	if err := json.NewDecoder(resp.Body).Decode(&labels); err != nil {
		return nil, err
	}
	return labels, nil
}


