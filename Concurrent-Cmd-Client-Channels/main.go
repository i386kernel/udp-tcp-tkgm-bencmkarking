package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os/exec"
	"strings"
)

func main() {
	c := make(chan string)
	// gpn := getPodName()
	//fmt.Println(gpn)
	//for _, j := range gpn {
	//	fmt.Println(j)
	//	go runcmd(j, c)
	//}
	podn := getPodNameClientGo()
	getPodNameClientGo()
	for _, j := range podn {
		fmt.Println(j)
		go runcmd(j, c)
	}

	fmt.Println(len(c))
	for j := range c {
		fmt.Println(j)
	}
	// fmt.Println(<-c)
	close(c)
	fmt.Println("done.....")
}

func getPodName() []string {
	cmd := exec.Command("kubectl", "get", "pods", "--no-headers")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fs := strings.Fields(string(out))
	var strcon []string
	for _, j := range fs {
		if strings.Contains(j, "redis-dep") {
			strcon = append(strcon, j)
		}
	}
	return strcon
}

func getPodNameClientGo() []string {

	kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "location to your kubeconfig file")
	flag.Parse()

	kconf := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{})

	currnamespace, _, err := kconf.Namespace()
	if err != nil {
		fmt.Println(err)
	}
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	var podnames []string
	pods, err := clientset.CoreV1().Pods(currnamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
		podnames = append(podnames, pod.Name)
	}
	return podnames
}

func runcmd(podname string, pc chan string) {
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-h", "172.16.9.76", "-n", "10000", "-c", "100")
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "set,lpush", "-n", "100000", "-q", "-h", "172.16.9.76")
	cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "set,lpush", "-n", "100000", "-q", "-h", "redis-dep.redis-deployment.svc.cluster.local")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	pc <- "==========================" + podname + "============================" + "\n" + string(out)
}
