use book_lib;
CREATE TABLE IF NOT EXISTS book(
  id int(12) AUTO_INCREMENT,
  title varchar(255),
  author varchar(65),
  year int(4),
  PRIMARY KEY (id)
)ENGINE=INNODB;
insert into book(id,title,author,year) values (1,"Head First Go","Jay McGavren ",2015);
insert into book(id,title,author,year) values (2,"Go in Action","William Kennedy, Brian Ketelsen, Erik St. Martin",2011);
insert into book(id,title,author,year) values (3,"Get Programming with Go","Nathan Youngman",2017);
insert into book(id,title,author,year) values (4,"Go Web Programming","Sau Sheong Chang",2013);