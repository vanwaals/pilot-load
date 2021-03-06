package simulation

var (
	serviceAccountYml = `
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
`
)

type ServiceAccountSpec struct {
	Namespace string
	Name      string
}

type ServiceAccount struct {
	Spec *ServiceAccountSpec
}

var _ Simulation = &ServiceAccount{}

func NewServiceAccount(s ServiceAccountSpec) *ServiceAccount {
	return &ServiceAccount{Spec: &s}
}

func (s ServiceAccount) Run(ctx Context) (err error) {
	return RunConfig(ctx, func() string { return render(serviceAccountYml, s.Spec) })

}
