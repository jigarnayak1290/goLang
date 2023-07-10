package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	NACCS_Code string `json:"NACCS_Code"`
	Owner_ID   string `json:"Owner_ID"`
	CreatedOn  string `json:"-"`
	UpdatedOn  string `json:"-"`
	DeletedOn  string `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return ProductList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func UpdateProduct(ID int, p *Product) error {
	_, pos, err := findProduct(ID)
	if err != nil {
		return err
	}

	p.ID = ID
	ProductList[pos] = p
	return nil
}

func UpdateProductByNACCSCode(NACCS_Code string, p *Product) error {
	_, pos, err := findProductByNACCSCode(NACCS_Code)
	if err != nil {
		return err
	}

	p.NACCS_Code = NACCS_Code
	ProductList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(ID int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.ID == ID {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func findProductByNACCSCode(NACCS_Code string) (*Product, int, error) {
	for i, p := range ProductList {
		if p.NACCS_Code == NACCS_Code {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.ID + 1
}

var ProductList = []*Product{
	&Product{
		ID:         1,
		Name:       "102 SUNG SHIN",
		Desc:       "Vassel 102",
		NACCS_Code: "13FZ",
		Owner_ID:   "D70P",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:         2,
		Name:       "HYODONG CHEMI",
		Desc:       "Vassel 107",
		NACCS_Code: "310P",
		Owner_ID:   "DSRG6",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}
