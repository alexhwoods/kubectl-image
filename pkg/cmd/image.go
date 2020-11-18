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
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd/api"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	namespaceExample = `
	# view the pod of the image foo
	%[1]s image foo
`

	errNoContext = fmt.Errorf("no context is currently set, use %q to select a new one", "kubectl config use-context <context>")
)

// ImageOptions provides information required to update
// the current context on a user's KUBECONFIG
type ImageOptions struct {
	configFlags      *genericclioptions.ConfigFlags
	resultingContext *api.Context

	rawConfig api.Config

	genericclioptions.IOStreams
}

// NewImageOptions provides an instance of ImageOptions with default values
func NewImageOptions(streams genericclioptions.IOStreams) *ImageOptions {
	return &ImageOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

// NewCmdImage provides a cobra command wrapping ImageOptions
func NewCmdImage(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewImageOptions(streams)

	cmd := &cobra.Command{
		Use:          "image [pod-name] [flags]",
		Short:        "View the image for some pod",
		Example:      fmt.Sprintf(namespaceExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Complete sets all information required for updating the current context
func (o *ImageOptions) Complete(cmd *cobra.Command, args []string) error {
	o.resultingContext.Namespace, _ = cmd.Flags().GetString("namespace")

	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *ImageOptions) Validate() error {
	return nil
}

// Runs the command
func (o *ImageOptions) Run() error {
	fmt.Print(o)
	// o.
	// fmt.Print(&(o.configFlags.KubeConfig))

	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }

	return nil
}
