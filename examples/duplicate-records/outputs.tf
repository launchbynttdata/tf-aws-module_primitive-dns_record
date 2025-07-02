# // Licensed under the Apache License, Version 2.0 (the "License");
# // you may not use this file except in compliance with the License.
# // You may obtain a copy of the License at
# //
# //     http://www.apache.org/licenses/LICENSE-2.0
# //
# // Unless required by applicable law or agreed to in writing, software
# // distributed under the License is distributed on an "AS IS" BASIS,
# // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# // See the License for the specific language governing permissions and
# // limitations under the License.

output "record_fqdns" {
  description = "FQDNs built using the zone domain and name."
  value       = module.dns_record.record_fqdns
}

output "route53_zone_zone_ids" {
  description = "Zone IDs of Route53 zone"
  value       = module.dns_zone.route53_zone_zone_ids
}
