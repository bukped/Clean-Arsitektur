# go-crud-mysql-sederhana
Create , Read, Edit, Delete data Mysql dengan bahasa Go


##Instalasi##

- Buat database ber-nama **go-crud** dan sebuah tabel di database mysql dengan nama tabel **users**

```
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `first_name` varchar(200) NOT NULL,
  `last_name` varchar(200) NOT NULL,
  `password` varchar(120) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

```

##