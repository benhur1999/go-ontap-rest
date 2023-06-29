package ontap

import (
	"net/http"
)

type IscsiService struct {
	Resource
	Enabled *bool `json:"enabled,omitempty"`
	Metric  *struct {
		Resource
		Duration string `json:"duration,omitempty"`
		Iops     struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"iops"`
		Latency struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"latency"`
		Status     string `json:"status,omitempty"`
		Throughput struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"throughput"`
		Timestamp string `json:"timestamp,omitempty"`
	} `json:"metric,omitempty"`
	Statistics *struct {
		IopsRaw struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"iops_raw"`
		LatencyRaw struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"latency_raw"`
		Status        string `json:"status,omitempty"`
		ThroughputRaw struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"throughput_raw"`
		Timestamp string `json:"timestamp,omitempty"`
	} `json:"statistics,omitempty"`
	Svm    *Resource `json:"svm,omitempty"`
	Target *struct {
		Alias string `json:"alias,omitempty"`
		Name  string `json:"name"`
	} `json:"target,omitempty"`
}

type IscsiServiceResponse struct {
	BaseResponse
	IscsiServices []IscsiService `json:"records,omitempty"`
}

func (c *Client) IscsiServiceGetIter(parameters []string) (iscsiServices []IscsiService, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/protocols/san/iscsi/services"
	reqParameters := parameters
	for {
		r := IscsiServiceResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		iscsiServices = append(iscsiServices, r.IscsiServices...)
		if r.IsPaginate() {
			path = r.GetNextRef()
			reqParameters = []string{}
		} else {
			break
		}
	}
	return
}

func (c *Client) IscsiServiceGet(href string, parameters []string) (*IscsiService, *RestResponse, error) {
	r := IscsiService{}
	req, err := c.NewRequest("GET", href, parameters, nil)
	if err != nil {
		return nil, nil, err
	}
	res, err := c.Do(req, &r)
	if err != nil {
		return nil, nil, err
	}
	return &r, res, nil
}
