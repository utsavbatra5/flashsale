The aim of the project is to handle race condition for 10 products of X in a flash sale where millions of users can cocurrently checkout for the same.

**Algorithm for race condition**

The race condition has been handled using timed semaphores. Semaphores have been created using channels in golang and at max 10 semaphores can be allocated. Only when a payment fails, 1 semaphore resource is released and another user is allowed to try payment.

**Project Configuration**

- The product name and stock have been configured in the config file. 
- The database type and configuration of various microservices are configured.
- By default, we are using Inmemory database only for the simulation purposes. Wrappers are provided in each microservice to implement SQL functions.

**Microservices details**

There are predominantly 3 microservices. All the APIs need a basic auth to run as per the configuration maintained.

1. Microservice Stock: This takes care of maintaining product stock. It comes with APIs to get current product stock as well update it.
- It is running on port 8001
- The APIs are as follows:
a) GET http://localhost:8001/api/{prodID}
   This gets the product stock. 
   Example: http://localhost:8001/api/X
  
b) POST http://localhost:8001/api/updatestock
   Sample Body:
    {
    "product": "X",
    "quantity": -1
    }
    Here we are reducing stock of product X by quantity 1.

2. Microservice Buy: This takes care of checkout of product from user cart. It permits at most 10 users to checkout and try payment using timed semaphores. 
- Only when a payment fails or gets timed out within 3 minutes, the next user queued in is permitted to try making a payment.
- The APIs are as follows:
a) http://localhost:8002/api/buy
  This API triggers simulation of user checkout and payment.
- User cart generator randomly generates user data and queues the requests on FIFO basis.
- Payment gateway simulation is done using a random logic to ensure only 1/45 payment succeeds and we can simulate high load traffic.

3. Microservice Add to cart: This allows user to add product to cart as well fetch all pending carts that are queued for checkout.
- The APIs are as follows:
a) POST http://localhost:8003/api/addtocart
  Sample body is:
  {
  "user": 1,
  "product": "x",
  "stock": 1
  }
  
b) GET http://localhost:8003/api/getcarts
This gets all pending carts for dequeue purpose.

**Unit testing**

All functions are covered by unit test cases.
    
