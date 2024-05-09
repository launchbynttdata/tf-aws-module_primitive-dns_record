package common

import (
	"github.com/launchbynttdata/lcaf-component-terratest/types"
)

type ThisTFModuleConfig struct {
	types.GenericTFModuleConfig
	Zones   map[string]Zones   `json:"zones"`
	Records map[string]Records `json:"records"`
}

type Zones struct {
	Domain_name   string            `json:"domain_name"`
	Comment       string            `json:"comment"`
	Force_destroy bool              `json:"force_destroy"`
	Tags          map[string]string `json:"tags"`
	Vpc           map[string]Vpc    `json:"vpc"`
}

type Vpc struct {
	Vpc_id     string `json:"vpc_id"`
	Vpc_region string `json:"vpc_region"`
}

type Records struct {
	Type                             string                     `json:"type"`
	Ttl                              int                        `json:"ttl"`
	Name                             string                     `json:"name"`
	Records                          []string                   `json:"records"`
	Set_identifier                   string                     `json:"set_identifier"`
	Health_check_id                  string                     `json:"health_check_id"`
	Alias                            Alias                      `json:"alias"`
	Cidr_routing_policy              Cidr_routing_policy        `json:"Cidr_routing_policy"`
	Failover_routing_policy          Failover_routing_policy    `json:"failover_routing_policy"`
	Geolocation_routing_policy       Geolocation_routing_policy `json:"geolocation_routing_policy"`
	Latency_routing_policy           Latency_routing_policy     `json:"latency_routing_policy"`
	Multivalue_answer_routing_policy bool                       `json:"multivalue_routing_policy"`
	Weighted_routing_policy          Weighted_routing_policy    `json:"weighted_routing_policy"`
	Allow_overwrite                  bool                       `json:"allow_overwrite"`
}

type Alias struct {
	Name                   string `json:"name"`
	Zone_id                string `json:"zone_id"`
	Evaluate_target_health bool   `json:"evaluate_target_health"`
}

type Cidr_routing_policy struct {
	Collection_id string `json:"collection_id"`
	Location_name string `json:"location_name"`
}

type Failover_routing_policy struct {
	Type string `json:"type"`
}

type Geolocation_routing_policy struct {
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	Subdivision string `json:"Subdivision"`
}

type Latency_routing_policy struct {
	Region string `json:"region"`
}

type Weighted_routing_policy struct {
	Weight int `json:"weight"`
}
