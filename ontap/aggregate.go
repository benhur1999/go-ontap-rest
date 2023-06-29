package ontap

import (
	"net/http"
)

type CloudStore struct {
	Resource
	MirrorDegraded            bool   `json:"mirror_degraded"`
	Availability              string `json:"availability"`
	Used                      int    `json:"used"`
	Primary                   bool   `json:"primary"`
	UnreclaimedSpaceThreshold int    `json:"unreclaimed_space_threshold"`
}

type CloudStorageTier struct {
	Used       int         `json:"used"`
	CloudStore *CloudStore `json:"stores,omitempty"`
}

type Aggregate struct {
	Resource
	SnaplockType string    `json:"snaplock_type,omitempty"`
	CreateTime   string    `json:"create_time,omitempty"`
	State        string    `json:"state,omitempty"`
	Node         *Resource `json:"node,omitempty"`
	HomeNode     *Resource `json:"home_node,omitempty"`
	DrHomeNode   *Resource `json:"dr_home_node,omitempty"`
	BlockStorage *struct {
		Mirror *struct {
			State   string `json:"state,omitempty"`
			Enabled bool   `json:"enabled"`
		} `json:"mirror,omitempty"`
		Primary *struct {
			RaidType      string `json:"raid_type"`
			DiskClass     string `json:"disk_class"`
			ChecksumStyle string `json:"checksum_style"`
			DiskCount     int    `json:"disk_count"`
			RaidSize      int    `json:"raid_size"`
		} `json:"primary,omitempty"`
		HybridCache *struct {
			Enabled   bool   `json:"enabled"`
			RaidType  string `json:"raid_type,omitempty"`
			Used      int    `json:"used,omitempty"`
			Size      int    `json:"size,,omitempty"`
			DiskCount int    `json:"disk_count,omitempty"`
		} `json:"hybrid_cache,omitempty"`
		Plexes *[]Resource `json:"plexes,omitempty"`
	} `json:"block_storage,omitempty"`
	CloudStorage *struct {
		AttachEligible           bool                `json:"attach_eligible"`
		TieringFullnessThreshold int                 `json:"tiering_fullness_threshold"`
		Stores                   *[]CloudStorageTier `json:"stores,omitempty"`
	} `json:"cloud_storage,omitempty"`
	DataEncryption *struct {
		SoftwareEncryptionEnabled bool `json:"software_encryption_enabled"`
		DriveProtectionEnabled    bool `json:"drive_protection_enabled"`
	} `json:"data_encryption,omitempty"`
	Space *struct {
		Footprint    int `json:"footprint"`
		BlockStorage struct {
			Available            int `json:"available"`
			Used                 int `json:"used"`
			Size                 int `json:"size"`
			InactiveUserData     int `json:"inactive_user_data"`
			FullThresholdPercent int `json:"full_threshold_percent"`
		} `json:"block_storage"`
		CloudStorage struct {
			Used int `json:"used"`
		} `json:"cloud_storage"`
		Efficiency struct {
			Ratio       float64 `json:"ratio"`
			LogicalUsed int     `json:"logical_used"`
			Savings     int     `json:"savings"`
		} `json:"efficiency"`
		EfficiencyWithoutSnapshots struct {
			Ratio       float64 `json:"ratio"`
			LogicalUsed int     `json:"logical_used"`
			Savings     int     `json:"savings"`
		} `json:"efficiency_without_snapshots"`
	} `json:"space,omitempty"`
}

type AggregateResponse struct {
	BaseResponse
	Aggregates []Aggregate `json:"records,omitempty"`
}

func (c *Client) AggregateGetIter(parameters []string) (aggregates []Aggregate, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/aggregates"
	reqParameters := parameters
	for {
		r := AggregateResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		aggregates = append(aggregates, r.Aggregates...)
		if r.IsPaginate() {
			path = r.GetNextRef()
			reqParameters = []string{}
		} else {
			break
		}
	}
	return
}

func (c *Client) AggregateGet(href string, parameters []string) (*Aggregate, *RestResponse, error) {
	r := Aggregate{}
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
