package arquitetura

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "am64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	t.Fail()
}
