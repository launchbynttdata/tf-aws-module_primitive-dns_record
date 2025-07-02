locals {
  zone_id = [for k, v in var.zones : module.dns_zone.route53_zone_zone_ids[k]]
}
