package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	NACCS_Code int     `json:"NACCS_Code"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Owner_ID   string  `json:"Owner_ID"`
	CreatedOn  string  `json:"-"`
	UpdatedOn  string  `json:"-"`
	DeletedOn  string  `json:"-"`
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
	p.NACCS_Code = getNextID()
	ProductList = append(ProductList, p)
}

func UpdateProduct(NACCS_Code int, p *Product) error {
	_, pos, err := findProduct(NACCS_Code)
	if err != nil {
		return err
	}

	p.NACCS_Code = NACCS_Code
	ProductList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(NACCS_Code int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.NACCS_Code == NACCS_Code {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.NACCS_Code + 1
}

var ProductList = []*Product{
	&Product{
		NACCS_Code: 1,
		Name:       "102 SUNG SHIN",
		Desc:       "Vassel 102",
		Price:      2.45,
		Owner_ID:   "D70P",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
	&Product{
		NACCS_Code: 2,
		Name:       "HYODONG CHEMI",
		Desc:       "Vassel 107",
		Price:      1.99,
		Owner_ID:   "DSRG6",
		CreatedOn:  time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}
