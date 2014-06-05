To set up a replica set we use the `mtools` package and issue the `mlaunch --replicaset` command. 

The first time I tried this the configuration file for the replica set was not intialized correctly and I had to go in and manually set it. Do this by running `mongo` and then create a configuration variable `config = { "_id" : "replset", "members" : [ { "_id" : 0, "host" : "localhost:27017"}, { "_id" : 1, "host" : "localhost:27018"}, { "_id" : 2, "host" : "localhost:27019" } ] }` and then call `rs.initiate(config)`. 

Running the Go code with `go run replica.go` should generate the output `Hello, my name is Bill DeRose and I am 20 years old`.
