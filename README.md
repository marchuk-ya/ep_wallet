# ep_wallet
# Description
- Database - Cassandra
- Configuration file with limit in conf.json file. 

## Start REST API server
  1. Run the command


      docker-compose build
  2. Run the command


      docker-compose up
      
## How To Check
- Enrollment: 


    `curl -X POST -H "Content-Type: application/json"
    -d '{"User_Id":"2", "Amount":"1000.0", "Currency": "EUR", "Description": "Food"}'
    localhost:10000/api/v1/post/enrollment`
- Outlay:


    `curl -X POST -H "Content-Type: application/json"
    -d '{"User_Id":"2", "Amount":"-500.0", "Currency": "EUR", "Description": "Food"}'
    localhost:10000/api/v1/post/enrollment`
    
    
  Response:
  
  
    in TODO

- Get balance for `<user_id>` in `<currency>`:
  
  
    `curl localhost:10000/api/v1/balance/{id}/{currency}`
  
  
  For example:
  
  
    `curl localhost:10000/api/v1/balance/1/EUR`
  
  
    `curl localhost:10000/api/v1/balance/2/EUR`
  
  
  Response:
  
  
    `["2450"]`
  
  
    `["500"]`
    
- Get a list of <user_id> operations between `<beginTime>` and `<endTime>`:
  
  
    `curl localhost:10000/api/v1/time/balance/{id}/{beginTime}/{endTime}`
    
    
  For example:
  
  
    `curl localhost:10000/api/v1/time/balance/2/2021-07-17/2021-07-19`
    
    
  Response:
  
  
    `[{"user_id":"1","amount":"1000","currency":"EUR","description":"Food"}]`
    
    
    
## TODO list
1. Implementation of authentication.
2. Implementation of logging with rotation of log-files. 
3. Generation of responses for each request with incorrect data. 
4. Implementation of unit-tests.
5. Generate HTML files from docstrings.  
