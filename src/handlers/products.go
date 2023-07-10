package handlers

import (
	"log"
	"net/http"

	//data "github.com/jigarnayak1290/goLang/src/Data"
	data "github.com/jigarnayak1290/goLang/tree/master/Learn/src/Data"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "mysecretpassword"
// 	dbname   = "postgres"
// )

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// //fmt.Println("Successfully connected!")

	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// put via id
	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL)

		//Check for ID in URI
		//reg := regexp.MustCompile(`/([0-9]+)`)
		// reg := regexp.MustCompile(`NACCS_Code=(\d+)`)
		// g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		queryParams := r.URL.Query()
		// Get the query parameters
		//queryParams := parsedURL.Query()

		if len(queryParams) != 1 {
			p.l.Println("Must have 1 parameter to be updated in URL, your parameter count is ", len(queryParams))
			http.Error(rw, "Parameter count is mismatch", http.StatusBadRequest)
			return
		}

		// Check if a parameter exists
		var receivedParams []string
		for param := range queryParams {
			receivedParams = append(receivedParams, param)
		}

		firstParam := receivedParams[0]
		if firstParam != "NACCS_Code" {
			p.l.Println("Must have NACCS_Code in parameter, your parameter name is ", receivedParams[0])
			http.Error(rw, "Incorrect parameter name ", http.StatusBadRequest)
			return
		}

		// Access individual parameter values
		NACCS_Code := queryParams.Get("NACCS_Code")
		//NACCS_Code, _ := strconv.Atoi(param1)
		// if len(g) != 1 {
		// 	p.l.Println("Invalid URL ID is not one -> ", len(g))
		// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
		// 	return
		// }

		// if len(g[0]) != 2 {
		// 	p.l.Println("Invalid URL more than capture group")
		// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
		// 	return
		// }

		// idString := g[0][1]
		// id, err := strconv.Atoi(idString)
		// if err != nil {
		// 	p.l.Println("Invalid URL unable to convert to number", idString)
		// 	http.Error(rw, "Invalid URI", http.StatusBadRequest)
		// 	return
		// }

		//p.updateProducts(id, rw, r)
		p.UpdateProductByNACCSCode(NACCS_Code, rw, r)
		return
		//p.l.Println("Got Id", id)
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

func (p *Products) UpdateProductByNACCSCode(NACCSID string, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProductByNACCSCode(NACCSID, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	//rw.Write(d)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
