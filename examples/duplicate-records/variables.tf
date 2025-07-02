// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

## DNS record related variables
variable "records" {
  type = map(object({
    type            = string
    ttl             = optional(number)
    name            = string
    records         = optional(list(string))
    set_identifier  = optional(string)
    health_check_id = optional(string)
    alias = optional(object({
      name                   = string
      zone_id                = string
      evaluate_target_health = bool
    }))
    cidr_routing_policy = optional(object({
      collection_id = string
      location_name = string
    }))
    failover_routing_policy = optional(object({
      type = string
    }))
    geolocation_routing_policy = optional(object({
      continent   = string
      country     = string
      subdivision = optional(string)
    }))
    latency_routing_policy = optional(object({
      region = string
    }))
    multivalue_answer_routing_policy = optional(bool)
    weighted_routing_policy = optional(object({
      weight = number
    }))
    allow_overwrite = optional(bool)
  }))

  description = "Records and their properties"
}

## DNS zone related variables
variable "tags" {
  description = "Tags added to all zones. Will take precedence over tags from the 'zones' variable"
  type        = map(string)
  default     = {}
}

variable "zones" {
  description = "Map of Route53 zone parameters"
  type        = map(any)
  default     = {}
}
