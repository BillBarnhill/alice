// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"github.com/getamis/alice/example/config"
	"github.com/getamis/alice/example/service"
	"github.com/getamis/alice/example/utils"
	"github.com/getamis/sirius/log"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var dkgCmd = &cobra.Command{
	Use:   "dkg",
	Short: "dkg",
	Long:  `dkg`,
	RunE: func(cmd *cobra.Command, args []string) error {
		initService(cmd)

		config, err := config.ReadYamlFile(configFile)
		if err != nil {
			log.Crit("Failed to read config file", "configFile", configFile, "err", err)
		}

		// Make a host that listens on the given multiaddress.
		host, err := makeBasicHost(config.Port)
		if err != nil {
			log.Crit("Failed to create a basic host", "err", err)
		}

		// Create a new peer manager.
		pm := newPeerManager(utils.GetPeerIDFromPort(config.Port), host)
		err = pm.addPeers(config.Peers)
		if err != nil {
			log.Crit("Failed to add peers", "err", err)
		}

		// Create a new service.
		service, err := service.NewDKGService(config, pm)
		if err != nil {
			log.Crit("Failed to new service", "err", err)
		}
		// Set a stream handler on the host.
		host.SetStreamHandler(dkgProtocol, func(s network.Stream) {
			service.Handle(s)
		})

		// Ensure all peers are connected before starting DKG process.
		pm.EnsureAllConnected()

		// Start DKG process.
		service.Process()

		return nil
	},
}

func init() {
	dkgCmd.Flags().String("config", "", "")
}

func initService(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	configFile = viper.GetString("config")

	return nil
}
