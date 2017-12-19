package main
import(
	"fmt"
	"strings"
	"archive/zip"
	"log"
	"io/ioutil"
	"regexp"
	//"go/types"
)
func leitura(filename string) (arquivo string){
	if !strings.HasSuffix(filename,"docx"){
		fmt.Println("Arquivo inválido")
		return
		}
	r,err := zip.OpenReader(filename)
	if err !=nil{
		log.Fatal(err)
	}
	defer r.Close()
	var f *zip.File
	var contem bool
	for _,f = range r.File{
		if strings.EqualFold(f.Name,"word/document.xml"){
			contem=true
			break
		}
	}
	if !contem{
		log.Fatal("documento inválido")
		return
	}
	rc,err := f.Open()
	if err !=nil{
		log.Fatal(err)
	}
	data,err := ioutil.ReadAll(rc)
	if err!=nil{
		log.Fatal(err)
	}
	defer rc.Close()
	arquivo = string(data)
	return
}
func normalizeQuotes(in rune) rune {
	switch in {
	case '“', '”':
		return '"'
	case '‘', '’':
		return '\''
	}
	return in
}
func normalizeAll(text string) string {
	brakets := regexp.MustCompile("<.*?>")
	quotes := regexp.MustCompile("&quot;")
	text = brakets.ReplaceAllString(text, "")
	text = quotes.ReplaceAllString(text, "\"")
	return strings.Map(normalizeQuotes, text)
}
func cleanText(text string) string {
	braketFinder := regexp.MustCompile("{{.*?}}")
	return braketFinder.ReplaceAllStringFunc(text, normalizeAll)
}

func cVariaveis(texto string)[]string{
	var y = strings.Split(texto,"[[")
	var x = make([]string,len(y))
	for i:=1;i<len(y)-1 ;i++  {
		var z = strings.Split(y[i],"]]")
		x[i]=z[0]
		//fmt.Println(x[i])

	}
	return x
}
type variavel struct{
	nome string
	sub []string
}
/*var lista =make([]variavel,5)
func seleciona(texto string) (lista []variavel){
	var x = strings.SplitN(texto,".",2)
	fmt.Println("valor = ",x)
	for i:=0;i<len(lista);i++  {
		if lista[i].nome==x[0]{
			lista[i].sub[len(lista[i].sub)]=x[1]
			return
		}
	}
	lista[len(lista)].nome=x[0]
	lista[len(lista)].sub[0]=x[1]
	return
}

func contem(v []string) bool{

}*/
func edVari(s string,v []string)( [] string){
	var x = make([]string,len(v))
	for i:=0;i<len(v) ;i++  {
		x[i]=strings.SplitN(v[i],".",2)[0]

	}
}
func main (){
	var teste = leitura("AgendaKids_Contrato_de_Locacao_de_Software_Educacional.docx")
	var x = normalizeAll(teste)

	var y = cVariaveis(x)
	/*var lista =make([]string,len(y))
	for i:=0; i<=len(y);i++  {

			 strings.SplitN(y[i],".",2)[0])

		}




	for i:=0;i<len(y) ;i++  {
		seleciona(y[i])
	}
	fmt.Println(len(lista))
	var v = make([]string,len(y))
	for i:=0;i<len(y) ;i++  {

		var aux string
		aux = string(y[i])
		var z = strings.Split(aux,"]]")
		v[i]=z[0]
		fmt.Println(v[i])
	}*/
	fmt.Println(y[0])
}