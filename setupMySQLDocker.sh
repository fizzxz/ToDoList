#change password and username to something more secure
#n.b. this is just for testing purposes atm
docker run --name test-mysql -p 3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
# get ip, port# dbname and password of the database and automate
# inputting into json file.

#use the following to produce a table for the db
#create table toDoListTest
#(
#	taskID int not null
#		primary key,
#	taskPriority int null,
#	taskCheck tinyint(1) null,
#	taskDescription char null,
#	taskCategory text null,
#	taskStartDate datetime null,
#	taskDueDate datetime null
#);

