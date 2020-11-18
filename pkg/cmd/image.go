/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	// "flag"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"context"
)

var (
	namespaceExample = `
	# view the pod of the image foo
	%[1]s image foo
`

	errNoContext = fmt.Errorf("no context is currently set, use %q to select a new one", "kubectl config use-context <context>")
)

type ImageOptions struct {
	configFlags *genericclioptions.ConfigFlags
}

// NewImageOptions provides an instance of ImageOptions with default values
func NewImageOptions(streams genericclioptions.IOStreams) *ImageOptions {
	return &ImageOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
	}
}

// NewCmdImage provides a cobra command wrapping ImageOptions
func NewCmdImage(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewImageOptions(streams)

	cmd := &cobra.Command{
		Use:          "image [pod-name] [flags]",
		Short:        "View the image for some pod",
		Example:      fmt.Sprintf(namespaceExample, "kubectl"),
		Args: 				cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Run(args[0]); err != nil {
				return err
			}

			return nil
		},
	}

	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Runs the command
func (o *ImageOptions) Run(podName string) error {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		if (*o.configFlags.KubeConfig == "") {
			kubeconfig = filepath.Join(home, ".kube", "config")
		} else {
			kubeconfig = *o.configFlags.KubeConfig
		}
	}

	ns := *o.configFlags.Namespace
	if (ns == "") {
		ns = "default"
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	pod, err := clientset.CoreV1().Pods(ns).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Pod not found in the %s namespace\n", ns)
	} else {
		for _, container := range pod.Spec.Containers {
			fmt.Println(container.Image)
		}
	}

	return nil
}
