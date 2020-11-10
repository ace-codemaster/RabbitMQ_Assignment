# RabbitMQ_Assignment
Assignment for Study Purpose

- EasternEnterprise folder contains required code

Folder Structure
 - Sender folder contains publisher code
 - Receiver folder contains subscriber code
 - Services folder contains services which will store data into database.
 - Model contains structs which will be used by GORM and to map transmitted json to a model.
 - helper folder contains reusebale libraries for
    1. Database connections
    2. Loggers
  
 
- Steps to execute program  
   1. Change Mysql credentials into file EasternEnterprise\helpers\dbhelper\dbhelper.go
   2. Change dsn string to your dsn
   3. Make sure RabbitMQ server is running
   4. Run receiver.go from EasternEnterprise\receiver
   5. Run sender.go from \EasternEnterprise\sender
   4. Check entries in MySQL

- Steps to execute test cases
  1. change working directory to \EasternEnterprise\services
  2. hotelService_test.go file contains test cases
  3. run following command --> "go test hotelService_test.go hotelService.go"
  
 
