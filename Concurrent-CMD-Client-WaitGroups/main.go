package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var lps []string

func main() {
	wg := new(sync.WaitGroup)
	gpn := getPodName()
	wg.Add(len(gpn))
	for _, j := range gpn {
		go runcmd(j, wg)
	}
	wg.Wait()
	fmt.Println("======All Get Operation=====")
	fmt.Println(lps)
	var alss []string
	for _, j := range lps {
		ss := strings.Split(j, " ")
		alss = append(alss, ss[1])
	}
	fmt.Println("=====ALL Get Values=====")
	fmt.Println(alss)
	var alssfl float64
	for _, j := range alss {
		sp, err := strconv.ParseFloat(j, 64)
		if err != nil {
			fmt.Println(err)
		}
		alssfl += sp
	}
	fmt.Println("Cumulative Operations in one second for Get Operations would be: ", alssfl)
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

func getPodName() []string {
	cmd := exec.Command("kubectl", "get", "pods", "--no-headers")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fs := strings.Fields(string(out))
	var strcon []string
	for _, j := range fs {
		if strings.Contains(j, "redis-client") { // use "redis-client" for other cluster client.
			strcon = append(strcon, j)
		}
	}
	return strcon
}

func runcmd(podname string, wg *sync.WaitGroup) {
	defer wg.Done()
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-h", "172.16.9.76", "-n", "10000", "-c", "100")
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "set,lpush", "-n", "100000", "-q", "-h", "172.16.9.76")
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "set,lpush", "-n", "10000", "-q", "-h", "redis-dep.redis-deployment.svc.cluster.local")
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "set,lpush", "-n", "10000", "-q", "-h", "172.16.9.76")
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "get", "-n", "100000", "-q", "-h", "172.16.9.76")
	// cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "get", "-n", "2000", "-q", "-h", "redis-dep.redis-deployment.svc.cluster.local")
	cmd := exec.Command("kubectl", "exec", podname, "--", "redis-benchmark", "-t", "get", "-n", "2000", "-P", "16", "-h", "redis-dep.redis-deployment.svc.cluster.local", "-q")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	//cl := cumulativeLPUSH(string(out), podname)
	//lps = append(lps, cl)

	cg := cumulativeGET(string(out), podname)
	lps = append(lps, cg)
}

//LPUSH: 20000.00
func cumulativeLPUSH(os, podname string) string {
	re := regexp.MustCompile(`LPUSH: [0-9].......`)
	fs := re.FindString(os)
	fmt.Printf("==================" + podname + "==================" + "\n")
	fmt.Println(os)
	return fs
}

func cumulativeGET(os, podname string) string {
	re := regexp.MustCompile(`GET: [0-9].......`)
	fs := re.FindString(os)
	fmt.Printf("==================" + podname + "==================" + "\n")
	fmt.Println(os)
	return fs
}
