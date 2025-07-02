## Set up the use case of ACM certificate validation records for a SAN certificate.
records = {
  "basename.nttdata-test3.com" = {
    type            = "CNAME"
    name            = "_8932841022ami3i17anc27.nttdata-test3.com"
    ttl             = 300
    records         = ["_25ae4821b2248a788e9f20948d.acm-validation.aws."]
    allow_overwrite = true
  }
  "*.basename.nttdata-test3.com" = {
    type            = "CNAME"
    name            = "_8932841022ami3i17anc27.nttdata-test3.com"
    ttl             = 300
    records         = ["_25ae4821b2248a788e9f20948d.acm-validation.aws."]
    allow_overwrite = true
  }
}
zones = {
  "launch.nttdata-test3.com" = {
    domain_name = "launch.nttdata-test3.com"
    comment     = "Test to create a public hosted zone through terraform"
    tags        = {}
  }
}
