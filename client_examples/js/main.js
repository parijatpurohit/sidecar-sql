var messages = require('./generated/js/User_FindByRollAndName_pb');
var services = require('./generated/js/server_grpc_pb');

var grpc = require('grpc');

function main() {
    var client = new services.StorageServiceClient('localhost:50051', grpc.credentials.createInsecure());
    var data = new messages.User_FindByRollAndNameRequest();
    var query = new messages.User_Query();
    query.setRoll(123);
    data.setQuery(query);
    client.user_FindByRollAndName(data, function(err, response) {
        console.log('Greeting:', response.toObject());
    });
}

main();