package app

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string

func NewUnhealthyCmd(args []string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "unhealthy",
		Short: "debug unhealthy pod network",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	flags := cmd.PersistentFlags()
	flags.StringVar(&kubeconfig, "kubeconfig", "", "kube config")
	flags.Parse(args)
	return cmd
}

func run() {
	var config *rest.Config
	var err error
	if kubeconfig != "" {
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	c := New(clientset)
	stop := make(chan struct{})
	defer close(stop)
	go c.Run(1, stop)
	// Wait forever
	select {}
}
