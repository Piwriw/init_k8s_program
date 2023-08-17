package app

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func NewAgentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "agent",
		Long: "agent 与 agent-cloud 交互，共同完成云边协同任务。",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

// GracefulShutdown 退出前的操作，执行一些清空操作
// TODO： 是否应该处于主线程，工作线程处于协程
func GracefulShutdown() error {
	// signal - SIGINT(control + c 产生), SIGABRT(abort()函数产生)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP, syscall.SIGABRT)
	select {
	case s := <-c:
		switch s {
		// TODO: 限定时间内关闭所有打开的资源
		case syscall.SIGTERM:
			//tasksC, _, err := dao.GetAllNodeTasksC("")
			//if err != nil {
			//	return errors.Wrap(err, "get all node tasks failed, Error: ")
			//}
			//klog.Infoln("these container will be remove: ", tasksC)
			//var err1 error
			//for _, v := range tasksC {
			//	klog.Infof("stop container: %v and it will be remove by its groutine", v.ContainerID.String)
			//	//err = docker.DockerClient.RemoveContainer(nodeTask.ContainerID.String)
			//	// notice: here is not to adopt this method deleting record by callbak
			//	//		   there is no enough time
			//	// TODO: consider how to optimize this piece of code
			//	err1 = docker.DockerClient.StopContainer(v)
			//	//err1 = docker.DockerClient.RemoveContainer(v.ContainerID.String)
			//	if err1 != nil {
			//		err1 = errors.Wrap(err1, "stop container failed")
			//		continue
			//	}
		}
		//klog.Infoln("wait to termination program....")
		//time.Sleep(15 * time.Second)
		//return err1
		//
		//default:
		//	klog.Infof("Get os signal %v", s.String())
		//	//_ = docker.DockerClient.RemoveContainer(containerID)
		//	// 这里进行一些清理的操作，比如关闭 docker...
		return nil
		//}
	}
}
