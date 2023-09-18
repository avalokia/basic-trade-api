## Final Project - Basic Trade API
The final project is to create an application called BasicTrade. The admin (seller) can create a product, storing the product's image, and create variant for the product.

The project is deployed on: 
### https://basic-trade-api-production-b019.up.railway.app


## Environment
In my local environment, I used **PostgreSQL** for the database. Thus, the package that I imported in `db.go` is `"gorm.io/driver/postgres"`. In cases where other database is used, please change accordingly.

For an example of the .env file, there's an `env.example` file in the repository.

**The .env file is ignored from being submitted into the repository, so please create a new .env file before executing the program**

## Endpoints
The postman collection is provided in `Base Trade API - Devi Tara Avalokita.postman_collection` file.
The `url` is set as an environment variable that can be changed accordingly.

| Method | URL | Description |
| ------ | --- | ----------- |
| POST | {{url}}/auth/register | Register new user as an admin |
| POST | {{url}}/auth/login | Login as an admin |
| GET | {{url}}/products | Get all products |
| POST | {{url}}/products | Create new product |
| PUT | {{url}}/products:productUUID | Update product |
| DELETE | {{url}}/products:productUUID | Delete product |
| GET | {{url}}/products:productUUID | Get product by UUID |
| GET | {{url}}/products/variants | Get all variants |
| POST | {{url}}/products/variants | Create new variant |
| PUT | {{url}}/products/variants/:variantUUID | Update variant |
| DELETE | {{url}}/products/variants/:variantUUID | Delete variant |
| GET | {{url}}/products/variants/:variantUUID | Get variant by UUID |