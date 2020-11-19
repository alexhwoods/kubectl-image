# kubectl-image

This repository implements a single kubectl plugin for quickly getting the images for any containers running in a pod.

It makes use of the genericclioptions in [k8s.io/cli-runtime](https://github.com/kubernetes/cli-runtime)
to generate a set of configuration flags which are in turn used to generate a raw representation of
the user's KUBECONFIG, as well as to obtain configuration which can be used with RESTClients when sending
requests to a kubernetes api server.

## Details

This plugin uses the [client-go library](https://github.com/kubernetes/client-go/tree/master/tools/clientcmd) to fetch a pod given the pod's name, and get its containers, and print the images of those containers.

Nothing is modified. Context is not switched, etc.

Built using Go version go1.15.5 darwin/amd64

## Example

```sh
$ kubectl image foo --namespace prod
foo-image
```

## Running

```sh
# assumes you have a working KUBECONFIG
$ GO111MODULE="on" go build cmd/kubectl-image.go
# place the built binary somewhere in your PATH
$ cp ./kubectl-image /usr/local/bin


$ kubectl image foo -n prod
```

## Cleanup

You can "uninstall" this plugin from kubectl by simply removing it from your PATH:

    $ rm /usr/local/bin/kubectl-image
