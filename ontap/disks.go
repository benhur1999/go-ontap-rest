package ontap

import (
	"net/http"
)

type Disk struct {
	Resource
	Node                 *Resource `json:"node,omitempty"`
	HomeNode             *Resource `json:"home_node,omitempty"`
	DrHomeNode           *Resource `json:"dr_home_node,omitempty"`
	Uid                  string    `json:"uid"`
	SerialNumber         string    `json:"serial_number,omitempty"`
	Model                string    `json:"model,omitempty"`
	Vendor               string    `json:"vendor,omitempty"`
	FirmwareVersion      string    `json:"firmware_version,omitempty"`
	UseableSize          int       `json:"usable_size,omitempty"`
	PhysicalSize         int       `json:"physical_size,omitempty"`
	SectorCount          int       `json:"sector_count,omitempty"`
	RightSizeSectorCount int       `json:"right_size_sector_count,omitempty"`
	BytesPerSector       int       `json:"bytes_per_sector,omitempty"`
	Rpm                  int       `json:"rpm,omitempty"`
	Type                 string    `json:"type,omitempty"`
	Class                string    `json:"class,omitempty"`
	ContainerType        string    `json:"container_type,omitempty"`
	Pool                 string    `json:"pool,omitempty"`
	Bay                  int       `json:"bay,omitempty"`
	Drawer               *struct {
		Id   int `json:"id"`
		Slot int `json:"slot"`
	} `json:"drawer,omitempty"`
	Shelf                *Resource   `json:"shelf,omitempty"`
	State                string      `json:"state,omitempty"`
	Aggregates           *[]Resource `json:"aggregates,omitempty"`
	RatedLifeUsedPercent int         `json:"rated_life_used_percent,omitempty"`
}

type DiskResponse struct {
	BaseResponse
	Disks []Disk `json:"records,omitempty"`
}

func (c *Client) DiskGetIter(parameters []string) (disks []Disk, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/disks"
	reqParameters := parameters
	for {
		r := DiskResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		disks = append(disks, r.Disks...)
		if r.IsPaginate() {
			path = r.GetNextRef()
			reqParameters = []string{}
		} else {
			break
		}
	}
	return
}

func (c *Client) DiskGet(href string, parameters []string) (*Disk, *RestResponse, error) {
	r := Disk{}
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
