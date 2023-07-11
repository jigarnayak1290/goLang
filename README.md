# Simple Vassel

Vassel detail program (Without DataBase feature added in source code)

## Usage

- To show all vassel detail
- To add new vassel detail
- To update current vassel detail by its NACCS code which is unique

## How to use

- Open Learn folder in visual code (Learn folder given in zip)
- Open console of vs code and enter below command
  go run main.go
- You can use any application which support api calling
  (I have used Postman application)
  
## endPoint information and calling sample

- To get all vassel detail
	open Postman and set below settings in Postman application
	calltype : Get
	URL : http://localhost:9090
	
	Press send button
	
- To add new vassel detail 
	open Postman and set below settings in Postman application
	calltype : Post
	URL : http://localhost:9090
	Body Type : raw
	Body Format : JSON
	Body : {"NACCS_Code":"NewNACCS_Code", "name":"New Vassel name", "Owner_ID":"Vassel owner name", "desc":"Vassel is white"}
	
	Press send button
	
- To update current vassel detail by its NACCS code which is unique
	open Postman and set below settings in Postman application
	calltype : Post
	URL : http://localhost:9090/?NACCS_Code=310P
	(Change NACCS_Code as per your original NACCS_Code)
	Body Type : raw
	Body Format : JSON
	Body : {"name":"New Vassel name", "Owner_ID":"Vassel owner name", "desc":"Vassel is white"}
	(Change Vassel detail as per your requierement)
	
	Press send button