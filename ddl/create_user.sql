-- mysql -u root
-- show databases
-- CREATE DATABASE batch_coincheck
-- ALTER DATABASE batch_coincheck DEFAULT CHARACTER SET=utf8;
-- GRANT ALL ON batch_coincheck.* to coin_app@localhost;
-- SET PASSWORD FOR coin_app@localhost=password('test');
GRANT ALL PRIVILEGES ON batch_coincheck.* TO 'coin_app'@'%' IDENTIFIED BY 'test';

# ログインする時
# mysql -h localhost -u coin_app -p
# use batch_coincheck;
