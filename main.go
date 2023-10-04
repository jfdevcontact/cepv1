package main



import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)


func main(){

	type dola struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		Uf          string `json:"uf"`
		Ibge        string `json:"ibge"`
		Gia         string `json:"gia"`
		Ddd         string `json:"ddd"`
		Siafi       string `json:"siafi"`
	}

	req, err := http.Get("https://viacep.com.br/ws/01001000/json/")
	if err != nil{
		fmt.Println("ERRO", err)
	}
	

	res,_ := io.ReadAll(req.Body)

	var data dola

	json.Unmarshal(res, &data)
	

	fmt.Println(data.Cep)
}