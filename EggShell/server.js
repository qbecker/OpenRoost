var express = require('express');
var app = express();
var router = express.Router();
var path = require('path');


var rtIndex = require('./routes/Index.js');

app.use(router)


//routes
app.use(express.static(path.join(__dirname, '/public')));

app.use('/', rtIndex);


app.listen(8080);

module.exports.router = router;
