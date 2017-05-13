package main

import (
	"gosible"
	"gosible-kubernetes"
)

type hosts GosibleHost {
	Hosts: GosibleHosts{
		GosibleHost {
			Location: "server1.keknet.net",
			Key: "~/.ssh/id_rsa",
			Port: 2222,
		}
		...,
	},
}

// type restartApacheHandler GosibleHandler{
// 	Name:
// }

type installPackagesTask gosible.Apt{
	Upgrade: true,
	Packages: gosible.Packages{
		"kubernetes@1.2.3",
		"docker@234567$%#@!#",
		"dritt",
	},
	State: gosible.state.AVAILABLE
}

type copyCrapTask gosible.Copy{
	DependsOn: [tasks..],
	Files: gosible.Files{ "**/*.dritt", "**/*.sh" }
	Source: "./files",
	Dest: "/var/www",
	Callback: func(task gosible.Task) gosible.Status {
		if task.Output == "shit" {
			return gosible.Error {
				Message: "This went to shit"
			}
		}
	},
}

type firstGroup = gosible.TaskGroup{
	Name: "Best group",
	Tasks: gosible.Tasks{
		installPackages,
		...,
	}
}

type installKubernetes GosibleRole {
	Name: "Install Kubernetes",
	Root: true,
	Tasks: gosible.Tasks{
		firstGroup,
		installPackagesTask,
	},
}

type config GosibleConf {
	Hosts: gosible.Hosts{""},
	Roles: GosibleRoles{
		installKubernetes,
		.....,
		.....,
	}
}

func main() {
	Gosible.init(config)
}
