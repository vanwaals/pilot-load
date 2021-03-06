package simulation

var (
	serviceYml = `
apiVersion: v1
kind: Service
metadata:
  name: {{.App}}
  namespace:  {{.Namespace}}
spec:
  clusterIP: {{.IP}}
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: {{.App}}
  type: ClusterIP
`
)

type ServiceSpec struct {
	App       string
	Namespace string
	IP        string
}

type Service struct {
	Spec *ServiceSpec
}

func (s Service) Run(ctx Context) (err error) {
	return RunConfig(ctx, func() string { return render(serviceYml, s.Spec) })
}

var _ Simulation = &Service{}

func NewService(s ServiceSpec) *Service {
	return &Service{Spec: &s}
}
