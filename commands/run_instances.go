package commands

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fuku2014/nifpo/nifcloud"
)

func runInstances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-instances",
		Short: "Create computing instance",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := runRunInstances()

			if err != nil {
				cmd.Printf(err.Error())
				os.Exit(1)
			}

			cmd.Printf(res)
		},
	}

	// TODO 起動スクリプト
	flags := cmd.Flags()
	flags.StringP("image-id", "", "", "OS image ID.")
	flags.StringP("key-name", "", "", "SSH Key Name.")
	flags.StringP("security-group", "", "", "FW Group Name.")
	flags.StringP("instance-type", "", "", "Server Type.")
	flags.StringP("availability-zone", "", "", "Zone.")
	flags.BoolP("disable-api-termination", "", false, "Whether to delete servers from the API.")
	flags.StringP("accounting-type", "", "2", "Price type.")
	flags.StringP("instance-id", "", "", "Server Name.")
	flags.StringP("admin", "", "", "Admin Account Name for windows.")
	flags.StringP("password", "", "", "Admin Account PW for windows.")
	flags.StringP("ip-type", "", "static", "IP address type.")
	flags.StringP("public-ip", "", "", "Elastic public IP address.")
	flags.BoolP("agreement", "", false, "Red Hat Enterprise Linux 5.8 64bit / 6.3 64bit, or consent if SPLA server is specified.")
	flags.StringP("description", "", "", "Description.")

	viper.BindPFlag("computing.describe-instances.instance-ids", flags.Lookup("instance-ids"))
	viper.BindPFlag("computing.describe-instances.tenancys", flags.Lookup("tenancys"))

	return cmd
}

func runRunInstances() (res string, err error) {
	client := nifcloud.NewClient(viper.GetString("region"), viper.GetString("access-key"), viper.GetString("secret-key"), viper.GetBool("debug"))
	params := map[string]string{"Action": "RunInstances"}

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
