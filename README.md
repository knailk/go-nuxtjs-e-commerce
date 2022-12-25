<h1 align="center">Design Ecommerce - Online Shopping System</h1>

**We'll cover the following:**

* [System Requirements](#system-requirements)
* [Use Case Diagram](#use-case-diagram)
* [Class Diagram](#class-diagram)
* [Entity Relationship Diagram](#entity-relationship-diagram)
* [Implementation](#implementation)

[Shopee e-commerce trading floor](https://shopee.vn/) is owned by SEA Group in Singapore. It is present in many countries such as Vietnam, Malaysia, Thailand, the Philippines,and Indonesia. So, I will focus on their online retail business where users can sell/buy their products.

### System Requirements

I will be designing a system with the following requirements:

1. Admin should be able to add new products to sell.
2. Admin can update/delete their products.
3. Users should be able to search for products by their name or category.
4. Customers can search and view all the products, but they will have to become a registered member to buy a product.
5. Customers should be able to add/remove/modify product items in their shopping cart.
6. Customers can check out and buy items in the shopping cart.
7. Customers should be able to specify a shipping address.
8. Customers can cancel an order if it has not shipped.
9. Customers should be able to pay through credit cards or cash.
10. Customers should be able to track their shipment to see the current state of their order.

* **Note:** Use case with red color was deleted, but it can be developed in the the future.
### Use Case Diagram

As described, We will have 2 main actors: 
* **Seller:**
* **Customer:** 

Here is the use case diagram of our Online Shopping System:

<p align="center">
    <img src="/doc/use-case.png" alt="Online Shoopping System Use Case Diagram">
    <br />
    Use Case Diagram
</p>

### Class Diagram

Here are the descriptions of the different classes of our Online Shopping System:
* **User:** There are three types of registered accounts in the system: one will be an Admin, who is responsible for adding new product categories and manage members; two is a Member, who can buy products and the last, a Seller, who can sell product.
* **Profile:** Save personal information of user, which can be changed.
* **Cart:** Customer will add product items that they intend to buy to the shopping cart.
* **Cart Item:** This class will encapsulate a product item that the users will be buying or placing in the shopping cart.
* **Product:** This class will encapsulate the entity that the users of our system will be buying and selling. Each Product will belong to a ProductCategory.
* **ProductCategory:** This will encapsulate the different categories of products, such as books, electronics, etc.
* **Order**: This will encapsulate a buying order to buy everything in the shopping cart.
* **Transaction**: This will keep status of payment and payment method.
* **OrderLog:** Will keep a track of the status of orders, such as unshipped, pending, complete, canceled, etc.
* **Payment:** This class will encapsulate the payment for an order. Members can pay through credit card or electronic bank transfer.
* **CreditCart**: This will encapsulate information of credit card.

* **Note:** Class with red color was deleted, but it can be developed in the the future.

<p align="center">
    <img src="/doc/class-diagram.png" alt="Online Shoopping System Class Diagram">
    <br />
    Class Diagram
</p>

### Entity Relationship Diagram

Following is the entity relationship diagram:

<p align="center">
    <img src="/doc/entity-relationship-diagram.png" alt="Entity relationship Diagram">
    <br />
    Entity relationship Diagram
</p>

### Implementation

After reading some post about Clean Architecture, I knew this arch is suitable for Go.

This project has 4 Domain layer :
* Models Layer(entity).
* Repository Layer.
* Usecase Layer.
* Delivery Layer.

The diagram (you can read more [here](https://github.com/bxcodec/go-clean-arch)):
<p align="center">
    <img src="/doc/clean-arch.png" alt="Clean Architecture">
    <br />
    Clean Architecture Diagram
</p>
