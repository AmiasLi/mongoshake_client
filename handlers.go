package mongoshake

import api "github.com/AmiasLi/mongoshake_client/models"

func (c *Client) GetConfig() (*api.Conf, error) {
	resp, err := c.request("/conf", &api.Conf{})
	if err != nil {
		return nil, err
	}
	return resp.(*api.Conf), nil
}

func (c *Client) GetExecutor() (*[]api.Executor, error) {
	executors, err := c.request("/executor", &[]api.Executor{})
	if err != nil {
		return nil, err
	}
	return executors.(*[]api.Executor), nil
}

func (c *Client) GetRepl() (*api.Repl, error) {
	repl, err := c.request("/repl", &api.Repl{})
	if err != nil {
		return nil, err
	}
	return repl.(*api.Repl), nil
}

func (c *Client) GetQueue() (*api.Queue, error) {
	queue, err := c.request("/queue", &api.Queue{})
	if err != nil {
		return nil, err
	}
	return queue.(*api.Queue), nil
}

func (c *Client) GetWorker() (*[]api.Worker, error) {
	workers, err := c.request("/worker", &[]api.Worker{})
	if err != nil {
		return nil, err
	}
	return workers.(*[]api.Worker), nil
}

func (c *Client) GetSentinel() (*api.Sentinel, error) {
	sentinel, err := c.request("/sentinel", &api.Sentinel{})
	if err != nil {
		return nil, err
	}
	return sentinel.(*api.Sentinel), nil
}
