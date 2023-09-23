package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	appsV1 "k8s.io/api/apps/v1"
	batchV1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

type workload struct {
	Namespace  string
	Name       string
	Containers []coreV1.Container
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "manifest file path")
	flag.Parse()

	if filePath != "" {
		flag.Usage()
		os.Exit(1)
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	manifests := string(b)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, spec := range strings.Split(manifests, "---") {
		if len(spec) == 0 {
			continue
		}
		obj, _, err := decode([]byte(spec), nil, nil)
		if err != nil {
			continue
		}

		// ワークロードリソース一覧: https://kubernetes.io/ja/docs/concepts/workloads/controllers/
		var workloads []workload
		switch obj.GetObjectKind().GroupVersionKind().Kind {
		case "Pod":
			p := obj.(*coreV1.Pod)
			workloads = append(workloads, workload{
				Namespace:  p.GetNamespace(),
				Name:       p.GetName(),
				Containers: p.Spec.Containers,
			})
		case "ReplicaSet":
			rs := obj.(*appsV1.ReplicaSet)
			workloads = append(workloads, workload{
				Namespace:  rs.GetNamespace(),
				Name:       rs.GetName(),
				Containers: rs.Spec.Template.Spec.Containers,
			})
		case "Deployment":
			d := obj.(*appsV1.Deployment)
			workloads = append(workloads, workload{
				Namespace:  d.GetNamespace(),
				Name:       d.GetName(),
				Containers: d.Spec.Template.Spec.Containers,
			})
		case "StatefulSet":
			ss := obj.(*appsV1.StatefulSet)
			workloads = append(workloads, workload{
				Namespace:  ss.GetNamespace(),
				Name:       ss.GetName(),
				Containers: ss.Spec.Template.Spec.Containers,
			})
		case "DaemonSet":
			ds := obj.(*appsV1.DaemonSet)
			workloads = append(workloads, workload{
				Namespace:  ds.GetNamespace(),
				Name:       ds.GetName(),
				Containers: ds.Spec.Template.Spec.Containers,
			})
		case "Job":
			j := obj.(*batchV1.Job)
			workloads = append(workloads, workload{
				Namespace:  j.GetNamespace(),
				Name:       j.GetName(),
				Containers: j.Spec.Template.Spec.Containers,
			})
		}
		if len(workloads) > 0 {
			b, err := json.Marshal(workloads)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := os.Stdout.Write(b); err != nil {
				log.Fatal(err)
			}
		}
	}
}
