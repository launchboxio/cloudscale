package main

import (
	"context"
	"fmt"
	"github.com/launchboxio/cloudscale/internal/integration"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"os"
	"path/filepath"
)

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "CloudScale operator for Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CloudScale Operator!")
		fmt.Println("This functionality is not yet implemented")

		cloudscaleAddr, _ := cmd.Flags().GetString("cloudscale-addr")
		client, err := loadClientSet()
		if err != nil {
			log.Fatal(err)
		}

		if err := integration.New(client, cloudscaleAddr).Run(context.Background()); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	operatorCmd.Flags().String("cloudscale-addr", "127.0.0.1:9001", "The CloudScale API endpoint")
	rootCmd.AddCommand(operatorCmd)
}

func loadClientSet() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if _, inCluster := os.LookupEnv("KUBERNETE_SERVICE_HOST"); inCluster {
		config, err = rest.InClusterConfig()
	} else {
		var kubeconfig string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		} else {
			kubeconfig = os.Getenv("KUBECONFIG")
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		return nil, err
	}
	// creates the clientset
	return kubernetes.NewForConfig(config)
}
