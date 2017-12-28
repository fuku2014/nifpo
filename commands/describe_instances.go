package commands

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fuku2014/nifpo/nifcloud"
)

func describeInstances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe-instances",
		Short: "Describe computing instances",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := runDescribeInstances()

			if err != nil {
				cmd.Printf(err.Error())
				os.Exit(1)
			}

			cmd.Printf(res)
		},
	}

	flags := cmd.Flags()
	flags.StringSliceP("instance-ids", "", nil, "One or more instance IDs. (optional)")
	flags.StringSliceP("tenancys", "", nil, "One or more Tenancys. (optional)")

	viper.BindPFlag("computing.describe-instances.instance-ids", flags.Lookup("instance-ids"))
	viper.BindPFlag("computing.describe-instances.tenancys", flags.Lookup("tenancys"))

	return cmd
}

func runDescribeInstances() (res string, err error) {
	client := nifcloud.NewClient(viper.GetString("region"), viper.GetString("access-key"), viper.GetString("secret-key"), viper.GetBool("debug"))
	params := map[string]string{"Action": "DescribeInstances"}

	instanceIds := viper.GetStringSlice("computing.describe-instances.instance-ids")
	tenancys := viper.GetStringSlice("computing.describe-instances.tenancys")

	for n, instanceId := range instanceIds {
		params["InstanceId."+strconv.Itoa(n+1)] = instanceId
	}

	for n, tenancy := range tenancys {
		params["Tenancy."+strconv.Itoa(n+1)] = tenancy
	}

	res, err = client.CallComputingAPI(params)
	return
}
