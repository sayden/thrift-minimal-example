var thrift = require('thrift');
var Hello = require('./Hello.js');
var ttypes = require('./hello_types.js');

var connection = thrift.createConnection("localhost", 3636, {
  transport : thrift.TFramedTransport,
  protocol : thrift.TBinaryProtocol
});

connection.on('error', function(err) {
  console.error("Connection error:", err);
});

var client = thrift.createClient(Hello, connection);

var hello = new ttypes.HelloRequest({message:"world!"})

client.hello(hello, function(err, response) {
  if (err) {
    console.error("Error:", err)
  } else {
    console.log(response.message);
    connection.end();
  }
});
