module github.com/joe-sonrichard/argo-kube-notifier

go 1.13

require (
	github.com/antchfx/jsonquery v0.0.0-20180821084212-a2896be8c82b
	github.com/argoproj-labs/argo-kube-notifier v0.0.0-20200310171754-b5cdadcd10fe
	github.com/nlopes/slack v0.5.1-0.20190515005541-e2954b1409b0
	github.com/onsi/gomega v1.5.0
	github.com/prometheus/common v0.4.0
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	gopkg.in/gomail.v2 v2.0.0-20150902115704-41f357289737
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.0
)

replace github.com/argoproj-labs/argo-kube-notifier => ./
