Model 从SQL取数据写在这里

例如 UserRepository


普通情况下，这里是 User 这个 Model 的 CRUD，但也有例外情况，可能会关联多张表，
甚至是换存储，例如换成Mongo或者其他的

所以其实这里按照惯例，会定义一个 UserRepository的接口