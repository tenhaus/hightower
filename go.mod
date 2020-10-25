module hightower

go 1.15

replace (
	github.com/containerd/containerd => github.com/containerd/containerd v1.3.4
	github.com/docker/docker => github.com/docker/docker v1.4.2-0.20200221181110-62bd5a33f707
)

require (
	github.com/GoogleContainerTools/skaffold v1.15.0 // indirect
	github.com/gruntwork-io/terratest v0.30.13 // indirect
	github.com/spf13/cobra v1.1.1
)
