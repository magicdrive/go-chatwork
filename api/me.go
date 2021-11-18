package api


type MeParam struct {

}

type MePram struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]string
}

type MeResult struct {
	Credential  string
	Method      string
	ResouceName string
	Params      map[string]string
}

func Me() {
}
