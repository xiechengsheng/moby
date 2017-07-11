package daemon

import (
	"github.com/docker/docker/engine"
)

func (daemon *Daemon) ContainerUpdate(job *engine.Job) engine.Status {
	if len(job.Args) != 1 {
		return job.Errorf("Usage: %s CONTAINER", job.Name)
	}
	name := job.Args[0]
	container, err := daemon.Get(name)
	container.LogEvent("update")
	if err != nil {
		return job.Error(err)
	}
	if job.EnvExists("Shares") {
		var Shares = job.GetenvInt64("Shares")
		if err := daemon.execDriver.Update(container.command, Shares); err != nil {
			return job.Errorf("Cannot update container %s: %s", name, err)
		}
	} else {
		return job.Errorf("Shares env does not exist %s: %s", name, err)
	}
	return engine.StatusOK
}