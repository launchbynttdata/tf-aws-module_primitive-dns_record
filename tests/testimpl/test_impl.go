package common

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/lib/aws/dns"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestDoesDNSZoneExist(t *testing.T, ctx types.TestContext) {
	t.Run("TestIsZoneExist", func(t *testing.T) {
		zoneIds := terraform.OutputMap(t, ctx.TerratestTerraformOptions(), "route53_zone_zone_ids")
		for zoneName := range zoneIds {
			expectedZoneName := ctx.TestConfig().(*ThisTFModuleConfig).Zones[zoneName].Domain_name
			zone := dns.GetHostedZoneById(t, zoneIds[zoneName])
			assert.Equal(t, *zone.HostedZone.Name, dns.NameNormalize(expectedZoneName))
		}
	})
}

func TestDoesDNSZoneRecordExist(t *testing.T, ctx types.TestContext) {
	t.Run("TestIsRecordExists", func(t *testing.T) {
		if !testDataHaveDNSRecords(t, ctx) {
			zoneIds := terraform.OutputMap(t, ctx.TerratestTerraformOptions(), "route53_zone_zone_ids")
			for zoneName := range zoneIds {
				for _, rec := range ctx.TestConfig().(*ThisTFModuleConfig).Records {
					fullQualifiedRecordName := rec.Name + "." + dns.NameNormalize(ctx.TestConfig().(*ThisTFModuleConfig).Zones[zoneName].Domain_name)
					_, err := dns.LookupDNSRecordInPublicRoute53ZoneByDNSProtocol(t, zoneIds[zoneName], fullQualifiedRecordName, rec.Type)
					assert.NoError(t, err, "can not find expected DNS record in AWS DNS: "+fullQualifiedRecordName)
				}
			}
		}

	})
}

func testDataHaveDNSRecords(t *testing.T, ctx types.TestContext) bool {
	return len(ctx.TestConfig().(*ThisTFModuleConfig).Records) > 0
}
