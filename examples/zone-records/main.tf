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

module "dns_zone" {
  source = "git::https://github.com/launchbynttdata/tf-aws-module_primitive-dns_zone.git?ref=1.0.0"

  zones = var.zones
  tags  = var.tags
}

# This example demonstrates creating multiple DNS records in a single DNS zone.
module "dns_record" {
  source = "../.."

  zone_id = local.zone_id[0]
  records = var.records
}
