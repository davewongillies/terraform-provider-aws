package aws

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSEc2ResourceTag_basic(t *testing.T) {
	var tag ec2.TagDescription

	testCheck := func(*terraform.State) error {
		key := aws.StringValue(tag.Key)
		if key != "Name" {
			return fmt.Errorf("Expected Key to be 'Name'; got '%s'", key)
		}

		value := aws.StringValue(tag.Value)
		if value != "Hello World" {
			return fmt.Errorf("Expected Value to be 'Hello World'; got '%s'", value)
		}

		return nil
	}

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEc2ResourceTagConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEc2ResourceTagExists(
						"aws_ec2_tag.test", &tag),
					testCheck,
				),
			},
		},
	})
}

func TestAccAWSEc2ResourceTag_subnet(t *testing.T) {
	var tag ec2.TagDescription

	testCheck := func(*terraform.State) error {
		key := aws.StringValue(tag.Key)
		if key != "Name" {
			return fmt.Errorf("Expected Key to be 'Name'; got '%s'", key)
		}

		value := aws.StringValue(tag.Value)
		if value != "Hello World" {
			return fmt.Errorf("Expected Value to be 'Hello World'; got '%s'", value)
		}

		return nil
	}

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEc2ResourceTagSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEc2ResourceTagExists(
						"aws_ec2_tag.test", &tag),
					testCheck,
				),
			},
		},
	})
}

func testAccCheckEc2ResourceTagExists(n string, tag *ec2.TagDescription) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		parts := strings.Split(rs.Primary.ID, ":")
		id := parts[0]
		key := parts[1]
		conn := testAccProvider.Meta().(*AWSClient).ec2conn
		resp, err := conn.DescribeTags(&ec2.DescribeTagsInput{
			Filters: []*ec2.Filter{
				{
					Name:   aws.String("resource-id"),
					Values: []*string{aws.String(id)},
				},
				{
					Name:   aws.String("key"),
					Values: []*string{aws.String(key)},
				},
			},
		})

		if err != nil {
			return err
		}

		if len(resp.Tags) == 0 {
			return fmt.Errorf("No tags found")
		}

		*tag = *resp.Tags[0]
		//		return fmt.Errorf("Tag found %s => %s", aws.StringValue(tag.Key), aws.StringValue(tag.Value))

		return nil
	}
}

const testAccEc2ResourceTagConfig = `
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_ec2_tag" "test" {
  resource_id = "${aws_vpc.test.id}"
  key         = "Name"
  value       = "Hello World"
}
`

const testAccEc2ResourceTagSubnetConfig = `
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "test" {
  vpc_id     = "${aws_vpc.test.id}"
  cidr_block = "10.0.1.0/24"
}

resource "aws_ec2_tag" "test" {
  resource_id = "${aws_subnet.test.id}"
  key         = "Name"
  value       = "Hello World"
}
`
